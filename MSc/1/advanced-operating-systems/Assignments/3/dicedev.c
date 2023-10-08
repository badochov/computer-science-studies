#include <linux/anon_inodes.h>
#include <linux/cdev.h>
#include <linux/file.h>
#include <linux/list.h>
#include <linux/interrupt.h>
#include <linux/delay.h>
#include <linux/module.h>
#include <linux/mutex.h>
#include <linux/fs.h>
#include <linux/pci.h>
#include <linux/uaccess.h>
#include <linux/wait.h>
#include <linux/mman.h>
#include <linux/mm.h>
#include <linux/spinlock.h>
#include <linux/workqueue.h>

#include "dicedev.h"

#define DICEDEV_MAX_DEVICES 256
#define DICEDEV_ERROR_MASK                                  \
	(DICEDEV_INTR_FEED_ERROR | DICEDEV_INTR_CMD_ERROR | \
	 DICEDEV_INTR_MEM_ERROR | DICEDEV_INTR_SLOT_ERROR)

MODULE_LICENSE("GPL");

struct dicedev_buffer {
	struct dicedev_ctx *ctx;

	size_t size;
	atomic_t seed;
	size_t num_pages;
	struct {
		dma_addr_t data_dma;
		u32 *page_data;
		struct {
			dma_addr_t data_dma;
			u32 *cmds;
		} *pages;
	} page_table;
	struct {
		struct list_head lh;
		struct mutex mutex;
		size_t read;
	} results;
	u32 allowed;
};

enum {
	STATE_RUNNING,
	STATE_DONE,
	STATE_ERROR,
};

struct dicedev_ctx {
	struct dicedev_device *dev;
	struct file *f;

	struct {
		struct list_head device_queue;
		struct list_head queued;
		struct list_head tasks;
		enum { CTX_FREE, CTX_RUNNING, CTX_ERROR } state;
		atomic_t task_state;
		wait_queue_head_t wq;
		struct mutex mutex;
	} work;
};

struct dicedev_task {
	struct list_head lh;

	u32 *cmds;
	size_t size;
	struct dicedev_buffer *buf;
	u32 seed;
};

struct dicedev_task_group {
	struct list_head lh;
	struct list_head tasks;

	struct dicedev_buffer *buf;
	u32 seed;
};

struct dicedev_device {
	struct pci_dev *pdev;
	struct cdev cdev;
	int idx;
	struct device *dev;
	void __iomem *bar;

	struct {
		struct workqueue_struct *work_queue;
		struct work_struct work;
		struct list_head queued;
		struct dicedev_ctx *running;
		bool encountered_error;
		bool suspended;
		wait_queue_head_t free_wq;
		spinlock_t slock;
	} work;
};

static dev_t dicedev_devno;
static struct dicedev_device *dicedev_devices[DICEDEV_MAX_DEVICES];
static DEFINE_MUTEX(dicedev_devices_lock);
static struct class dicedev_class = {
	.name = "dicedev",
	.owner = THIS_MODULE,
};

/* Hardware handling. */

static inline void dicedev_iow(struct dicedev_device *dev, u32 reg, u32 val)
{
	iowrite32(val, dev->bar + reg);
}

static inline u32 dicedev_ior(struct dicedev_device *dev, u32 reg)
{
	u32 res = ioread32(dev->bar + reg);
	return res;
}

static void dicedev_dev_start(struct dicedev_device *dev)
{
	dicedev_iow(dev, DICEDEV_INTR, 0);
	dicedev_iow(dev, DICEDEV_INTR_ENABLE,
		    DICEDEV_INTR_FENCE_WAIT | DICEDEV_INTR_FEED_ERROR |
			    DICEDEV_INTR_MEM_ERROR | DICEDEV_INTR_SLOT_ERROR |
			    DICEDEV_INTR_CMD_ERROR);
	dicedev_iow(dev, DICEDEV_ENABLE, 1);
	dicedev_iow(dev, DICEDEV_CMD_FENCE_WAIT, 0);
}

static void dicedev_dev_stop(struct dicedev_device *dev)
{
	dicedev_iow(dev, DICEDEV_ENABLE, 0);
	dicedev_iow(dev, DICEDEV_INTR_ENABLE, 0);
}

static void dicedev_dev_restart(struct dicedev_device *dev)
{
	dicedev_dev_start(dev);
}

static bool dicedev_ctx_caused_error(struct dicedev_ctx *ctx)
{
	return atomic_read(&ctx->work.task_state) == STATE_ERROR;
}

static void dicedev_issue_cmds(struct dicedev_device *dev, u32 *cmds,
			       size_t count)
{
	size_t cmdIdx;
	u32 free = dicedev_ior(dev, CMD_MANUAL_FREE);

	for (cmdIdx = 0; cmdIdx < count; cmdIdx++, free--) {
		while (free == 0) {
			free = dicedev_ior(dev, CMD_MANUAL_FREE);
		}
		dicedev_iow(dev, CMD_MANUAL_FEED, cmds[cmdIdx]);
	}
}

static void dicedev_issue_fence(struct dicedev_device *dev)
{
	u32 cmd;

	cmd = DICEDEV_USER_CMD_FENCE_HEADER(0);

	dicedev_issue_cmds(dev, &cmd, 1);
}

static void dicedev_issue_buffer_unbind(struct dicedev_device *dev)
{
	u32 cmd = DICEDEV_USER_CMD_UNBIND_SLOT_HEADER(0);

	dicedev_issue_cmds(dev, &cmd, 1);
}

static void dicedev_issue_buffer_bind(struct dicedev_device *dev,
				      struct dicedev_buffer *buf, u32 seed)
{
	u32 cmds[4];

	cmds[0] = DICEDEV_USER_CMD_BIND_SLOT_HEADER(0, seed);
	cmds[1] = buf->allowed & 0xffffffffu;
	cmds[2] = buf->page_table.data_dma & 0xffffffffu;
	cmds[3] = (buf->page_table.data_dma >> 32) & 0xffffffffu;

	dicedev_issue_cmds(dev, cmds, 4);
}

static int dicedev_work_handle_task_group(struct dicedev_ctx *ctx,
					  struct dicedev_task_group *task_group)
{
	struct dicedev_task *task;

	dicedev_issue_buffer_bind(ctx->dev, task_group->buf, task_group->seed);

	list_for_each_entry(task, &task_group->tasks, lh) {
		dicedev_issue_cmds(ctx->dev, task->cmds, task->size);
		if (dicedev_ctx_caused_error(ctx)) {
			break;
		}
	}

	dicedev_issue_buffer_unbind(ctx->dev);

	return dicedev_ctx_caused_error(ctx);
}

static void dicedev_work_handle_ctx(struct dicedev_ctx *ctx)
{
	struct dicedev_task_group *task_group;

	list_for_each_entry(task_group, &ctx->work.tasks, lh) {
		if (dicedev_work_handle_task_group(ctx, task_group)) {
			break;
		}
	}

	dicedev_issue_fence(ctx->dev);
}

static void dicedev_work_handler(struct work_struct *work)
{
	struct dicedev_device *dev =
		container_of(work, struct dicedev_device, work.work);
	struct dicedev_ctx *ctx;
	unsigned long flags;

	spin_lock_irqsave(&dev->work.slock, flags);
	while (dev->work.running != NULL || dev->work.suspended) {
		spin_unlock_irqrestore(&dev->work.slock, flags);
		wait_event(dev->work.free_wq,
			   dev->work.running == NULL && !dev->work.suspended);
		spin_lock_irqsave(&dev->work.slock, flags);
	}

	BUG_ON(list_empty(&dev->work.queued));
	ctx = list_first_entry(&dev->work.queued, typeof(*ctx),
			       work.device_queue);
	list_del(&ctx->work.device_queue);
	dev->work.running = ctx;

	spin_unlock(&dev->work.slock);

	dicedev_work_handle_ctx(ctx);
}

static void dicedev_ctx_queue_work(struct dicedev_ctx *ctx)
{
	unsigned long flags;

	spin_lock_irqsave(&ctx->dev->work.slock, flags);
	list_add_tail(&ctx->work.device_queue, &ctx->dev->work.queued);
	spin_unlock_irqrestore(&ctx->dev->work.slock, flags);

	queue_work(ctx->dev->work.work_queue, &ctx->dev->work.work);
}

static void dicedev_isr_error(struct dicedev_device *dev)
{
	dev->work.encountered_error = true;

	dicedev_dev_restart(dev);
}

static void dicedev_isr_fence(struct dicedev_device *dev)
{
	struct dicedev_ctx *ctx = dev->work.running;

	if (dev->work.encountered_error) {
		dev->work.encountered_error = false;
		atomic_set(&ctx->work.task_state, STATE_ERROR);
	} else {
		atomic_set(&ctx->work.task_state, STATE_DONE);
	}

	dev->work.running = NULL;

	wake_up_all(&dev->work.free_wq);
	wake_up(&ctx->work.wq);
}

/* IRQ handler. */

static irqreturn_t dicedev_isr(int irq, void *opaque)
{
	struct dicedev_device *dev = opaque;
	unsigned long flags;
	u32 istatus;

	spin_lock_irqsave(&dev->work.slock, flags);

	istatus = dicedev_ior(dev, DICEDEV_INTR) &
		  dicedev_ior(dev, DICEDEV_INTR_ENABLE);

	dicedev_iow(dev, DICEDEV_INTR, istatus);

	if (istatus & DICEDEV_ERROR_MASK) {
		BUG_ON(istatus &
		       (DICEDEV_INTR_FEED_ERROR | DICEDEV_INTR_SLOT_ERROR));

		dicedev_isr_error(dev);
	}
	if (istatus & DICEDEV_INTR_FENCE_WAIT) {
		dicedev_isr_fence(dev);
	}

	spin_unlock_irqrestore(&dev->work.slock, flags);

	return IRQ_RETVAL(istatus);
}

static long dicedev_buffer_ioctl_seed(struct dicedev_buffer *buf,
				      struct dicedev_ioctl_seed __user *arg)
{
	struct dicedev_ioctl_seed seed;

	if (copy_from_user(&seed, arg, sizeof(seed))) {
		return -EFAULT;
	}

	if ((seed.seed & 0xfffff) != seed.seed) {
		return -EINVAL;
	}

	atomic_set(&buf->seed, seed.seed);
	return 0;
}

static int dicedev_cmds_validate(u32 *cmds, size_t size)
{
	size_t i;
	bool in_cmd = false;

	for (i = 0; i < size; i++) {
		if (in_cmd) {
			in_cmd = false;
		} else {
			switch (cmds[i] & 0xf) {
			case DICEDEV_USER_CMD_TYPE_NOP:
			case DICEDEV_USER_CMD_TYPE_NEW_SET:
				break;
			case DICEDEV_USER_CMD_TYPE_GET_DIE:
				in_cmd = true;
				break;
			default:
				return -EINVAL;
			}
		}
	}

	return in_cmd ? -EINVAL : 0;
}

static struct dicedev_task *dicedev_new_cmds_work(struct dicedev_buffer *buf,
						  u32 *cmds, size_t size)
{
	struct dicedev_task *task = kmalloc(sizeof(*task), GFP_KERNEL);
	if (task == NULL) {
		return NULL;
	}

	task->cmds = cmds;
	task->buf = buf;
	task->seed = atomic_read(&buf->seed);
	task->size = size;

	return task;
}

static int dicedev_ctx_add_work(struct dicedev_ctx *ctx,
				struct dicedev_task *item)
{
	int err = 0;

	mutex_lock(&ctx->work.mutex);
	if (ctx->work.state == CTX_ERROR) {
		err = -EIO;
	} else {
		list_add_tail(&item->lh, &ctx->work.queued);
	}
	mutex_unlock(&ctx->work.mutex);

	return err;
}

static ssize_t dicedev_buffer_write(struct file *file, const char __user *ubuf,
				    size_t len, loff_t *off)
{
	struct dicedev_buffer *buf = file->private_data;
	struct dicedev_task *task;
	ssize_t err = 0;
	u32 *cmds;

	if ((len % sizeof(u32)) != 0) {
		err = -EINVAL;

		goto err_init_validation;
	}

	cmds = kvmalloc(len, GFP_KERNEL);
	if (cmds == NULL) {
		err = -ENOMEM;

		goto err_malloc_buf;
	}

	if (copy_from_user(cmds, ubuf, len)) {
		goto err_copy;
	}

	task = dicedev_new_cmds_work(buf, cmds, len / sizeof(u32));
	if (task == NULL) {
		err = -ENOMEM;

		goto err_work;
	}

	if ((err = dicedev_ctx_add_work(buf->ctx, task))) {
		goto err_ctx_add;
	}

	return len;

err_ctx_add:
	kfree(task);
err_work:
err_copy:
	kfree(cmds);
err_malloc_buf:
err_init_validation:
	return err;
}

static void list_skip(struct list_head **lh, struct list_head *head, u32 skip)
{
	while (skip-- && !list_is_head(*lh, head)) {
		*lh = (*lh)->prev;
	}
}

static bool list_contains(struct list_head *lh, struct list_head *head)
{
	struct list_head *curr;

	list_for_each(curr, head) {
		if (curr == lh) {
			return true;
		}
	}
	return false;
}

static void dicedev_task_destroy(struct dicedev_task *task)
{
	list_del(&task->lh);
	kfree(task->cmds);

	kfree(task);
}

static void dicedev_task_group_destroy(struct dicedev_task_group *task_group)
{
	struct dicedev_task *task;
	struct dicedev_task *tmp;

	list_for_each_entry_safe(task, tmp, &task_group->tasks, lh) {
		dicedev_task_destroy(task);
	}

	kfree(task_group);
}

static void dicedev_ctx_tasks_destroy(struct dicedev_ctx *ctx)
{
	struct dicedev_task_group *curr;
	struct dicedev_task_group *tmp;

	list_for_each_entry_safe(curr, tmp, &ctx->work.tasks, lh) {
		dicedev_task_group_destroy(curr);
	}

	INIT_LIST_HEAD(&ctx->work.tasks);
}

static void dicedev_release_results(struct list_head *results)
{
	struct dicedev_buffer *buf;

	list_for_each_entry(buf, results, results.lh) {
		buf->results.read = 0;
		mutex_unlock(&buf->results.mutex);
	}
}

static int dicedev_ctx_split_to_task_groups(struct list_head *to_run,
					    struct list_head *tasks,
					    struct list_head *results)
{
	struct dicedev_task_group *task_group;
	struct dicedev_task *curr;
	struct dicedev_task *tmp;

	while (!list_empty(to_run)) {
		task_group = kmalloc(sizeof(*task_group), GFP_KERNEL);
		if (task_group == NULL) {
			return -EIO;
		}
		INIT_LIST_HEAD(&task_group->tasks);
		list_add_tail(&task_group->lh, tasks);

		curr = list_first_entry(to_run, typeof(*curr), lh);

		task_group->buf = curr->buf;
		task_group->seed = curr->seed;

		if (!list_contains(&task_group->buf->results.lh, results)) {
			mutex_lock(&task_group->buf->results.mutex);
			list_add(&task_group->buf->results.lh, results);
		}

		list_for_each_entry_safe(curr, tmp, to_run, lh) {
			if (task_group->buf == curr->buf &&
			    task_group->seed == curr->seed) {
				if (dicedev_cmds_validate(curr->cmds,
							  curr->size)) {
					return -EIO;
				}
				list_move_tail(&curr->lh, &task_group->tasks);
			}
		}
	}

	return 0;
}

static void dicedev_buffer_work_release(struct dicedev_buffer *buf)
{
	struct dicedev_ctx *ctx = buf->ctx;
	struct dicedev_task *task;
	struct dicedev_task *tmp;

	mutex_lock(&ctx->work.mutex);
	while (ctx->work.state == CTX_RUNNING) {
		mutex_unlock(&ctx->work.mutex);
		wait_event(ctx->work.wq, ctx->work.state != CTX_RUNNING);
		mutex_lock(&ctx->work.mutex);
	}

	list_for_each_entry_safe(task, tmp, &ctx->work.queued, lh) {
		if (task->buf == buf) {
			dicedev_task_destroy(task);
		}
	}

	mutex_unlock(&ctx->work.mutex);
	wake_up(&ctx->work.wq);
}

static int dicedev_wait_tasks(struct dicedev_ctx *ctx, struct list_head *to_run)
{
	int err;
	struct list_head results;

	if (list_empty(to_run)) {
		return 0;
	}

	INIT_LIST_HEAD(&results);

	if ((err = dicedev_ctx_split_to_task_groups(to_run, &ctx->work.tasks,
						    &results))) {
		goto split_to_task_groups;
	}

	atomic_set(&ctx->work.task_state, STATE_RUNNING);
	dicedev_ctx_queue_work(ctx);

	wait_event(ctx->work.wq,
		   atomic_read(&ctx->work.task_state) != STATE_RUNNING);

	if (atomic_read(&ctx->work.task_state) == STATE_ERROR) {
		err = -EIO;
		goto err_task_error;
	}

	dicedev_release_results(&results);
	dicedev_ctx_tasks_destroy(ctx);

	return 0;

err_task_error:
split_to_task_groups:
	dicedev_release_results(&results);
	dicedev_ctx_tasks_destroy(ctx);

	return err;
}

static int dicedev_buffer_wait(struct dicedev_buffer *buf)
{
	struct dicedev_ctx *ctx = buf->ctx;
	struct dicedev_task *task;
	struct dicedev_task *tmp;
	struct list_head to_run;
	int err;

	mutex_lock(&ctx->work.mutex);
	while (ctx->work.state == CTX_RUNNING) {
		mutex_unlock(&ctx->work.mutex);
		wait_event(ctx->work.wq, ctx->work.state != CTX_RUNNING);
		mutex_lock(&ctx->work.mutex);
	}

	if (ctx->work.state == CTX_ERROR) {
		err = -EIO;
		goto err_lock;
	}

	INIT_LIST_HEAD(&to_run);

	list_for_each_entry_safe(task, tmp, &ctx->work.queued, lh) {
		if (task->buf == buf) {
			list_move_tail(&task->lh, &to_run);
		}
	}

	ctx->work.state = CTX_RUNNING;

	mutex_unlock(&ctx->work.mutex);

	if ((err = dicedev_wait_tasks(ctx, &to_run))) {
		goto err_wait;
	}

	mutex_lock(&ctx->work.mutex);

	ctx->work.state = CTX_FREE;

	mutex_unlock(&ctx->work.mutex);
	wake_up(&ctx->work.wq);

	return 0;

err_wait:
	mutex_lock(&ctx->work.mutex);

	ctx->work.state = CTX_ERROR;
err_lock:
	mutex_unlock(&ctx->work.mutex);
	wake_up(&ctx->work.wq);

	return err;
}

static int dicedev_wait(struct dicedev_ctx *ctx, u32 skip)
{
	struct list_head *lh;
	struct list_head to_run;
	int err;

	mutex_lock(&ctx->work.mutex);
	while (ctx->work.state == CTX_RUNNING) {
		mutex_unlock(&ctx->work.mutex);
		wait_event(ctx->work.wq, ctx->work.state != CTX_RUNNING);
		mutex_lock(&ctx->work.mutex);
	}

	if (ctx->work.state == CTX_ERROR) {
		err = -EIO;
		goto err_lock;
	}

	lh = ctx->work.queued.prev;
	list_skip(&lh, &ctx->work.queued, skip);

	ctx->work.state = CTX_RUNNING;

	INIT_LIST_HEAD(&to_run);
	list_cut_position(&to_run, &ctx->work.queued, lh);

	mutex_unlock(&ctx->work.mutex);

	if ((err = dicedev_wait_tasks(ctx, &to_run))) {
		goto err_wait;
	}

	mutex_lock(&ctx->work.mutex);

	ctx->work.state = CTX_FREE;

	mutex_unlock(&ctx->work.mutex);
	wake_up(&ctx->work.wq);

	return 0;

err_wait:

	mutex_lock(&ctx->work.mutex);

	ctx->work.state = CTX_ERROR;
err_lock:
	mutex_unlock(&ctx->work.mutex);
	wake_up(&ctx->work.wq);

	return err;
}

static ssize_t dicedev_buffer_read(struct file *file, char __user *ubuf,
				   size_t len, loff_t *off)
{
	struct dicedev_buffer *buf = file->private_data;
	size_t page;
	size_t size;
	size_t in_page_offs;
	size_t read_size;
	size_t offs;
	ssize_t err = 0;

	if (len % sizeof(struct dice) != 0 || *off != 0) {
		return -EINVAL;
	}

	if ((err = dicedev_buffer_wait(buf))) {
		return err;
	}

	mutex_lock(&buf->results.mutex);

	offs = buf->results.read * sizeof(struct dice);

	len = min(len, buf->size - offs);

	for (size = 0, page = offs / DICEDEV_PAGE_SIZE; size != len;
	     size += read_size, page++) {
		in_page_offs = (offs + size) % DICEDEV_PAGE_SIZE;
		read_size = min(len - size, DICEDEV_PAGE_SIZE - in_page_offs);
		if (copy_to_user(ubuf + size,
				 ((char *)buf->page_table.pages[page].cmds) +
					 in_page_offs,
				 read_size)) {
			err = -EFAULT;
			break;
		}
	}
	buf->results.read += size / sizeof(struct dice);

	mutex_unlock(&buf->results.mutex);

	return err || size;
}

vm_fault_t dicedev_buffer_mmap_fault(struct vm_fault *vmf)
{
	struct dicedev_buffer *buf = vmf->vma->vm_file->private_data;
	struct page *page;

	if (buf->num_pages <= vmf->pgoff) {
		return VM_FAULT_SIGBUS;
	}

	page = virt_to_page(buf->page_table.pages[vmf->pgoff].cmds);

	get_page(page);
	vmf->page = page;

	return 0;
}

int dicedev_buffer_mmap(struct file *file, struct vm_area_struct *vm_area)
{
	static const struct vm_operations_struct dicedev_buffer_vm_ops = {
		.fault = dicedev_buffer_mmap_fault,
	};

	vm_area->vm_ops = &dicedev_buffer_vm_ops;

	return 0;
}

static long dicedev_ioctl_run(struct dicedev_ctx *ctx,
			      struct dicedev_ioctl_run __user *arg);

static long dicedev_buffer_ioctl(struct file *file, unsigned int cmd,
				 unsigned long arg)
{
	struct dicedev_buffer *buf = file->private_data;

	switch (cmd) {
	case DICEDEV_BUFFER_IOCTL_SEED:
		return dicedev_buffer_ioctl_seed(
			buf, (struct dicedev_ioctl_seed __user *)arg);
	case DICEDEV_IOCTL_RUN:
		return dicedev_ioctl_run(
			buf->ctx, (struct dicedev_ioctl_run __user *)arg);
	default:
		return -EINVAL;
	};
}

static int dicedev_buffer_release(struct inode *inode, struct file *file)
{
	struct dicedev_buffer *buf = file->private_data;
	struct dicedev_ctx *ctx = buf->ctx;
	size_t page;

	dicedev_buffer_work_release(buf);

	mutex_destroy(&buf->results.mutex);

	dma_free_coherent(&ctx->dev->pdev->dev, DICEDEV_PAGE_SIZE,
			  buf->page_table.page_data, buf->page_table.data_dma);
	for (page = 0; page < buf->num_pages; page++) {
		dma_free_coherent(&ctx->dev->pdev->dev, DICEDEV_PAGE_SIZE,
				  buf->page_table.pages[page - 1].cmds,
				  buf->page_table.pages[page - 1].data_dma);
	}

	kfree(buf);

	fput(ctx->f);

	return 0;
}

static const struct file_operations dicedev_buffer_file_ops = {
	.owner = THIS_MODULE,
	.write = dicedev_buffer_write,
	.read = dicedev_buffer_read,
	.compat_ioctl = dicedev_buffer_ioctl,
	.mmap_supported_flags = MAP_SHARED,
	.mmap = dicedev_buffer_mmap,
	.unlocked_ioctl = dicedev_buffer_ioctl,
	.release = dicedev_buffer_release,
};

/* Main device node handling.  */

static int dicedev_open(struct inode *inode, struct file *file)
{
	struct dicedev_device *dev =
		container_of(inode->i_cdev, struct dicedev_device, cdev);
	struct dicedev_ctx *ctx = kzalloc(sizeof *ctx, GFP_KERNEL);

	if (ctx == NULL) {
		return -ENOMEM;
	}
	ctx->dev = dev;

	mutex_init(&ctx->work.mutex);
	INIT_LIST_HEAD(&ctx->work.device_queue);
	INIT_LIST_HEAD(&ctx->work.tasks);
	INIT_LIST_HEAD(&ctx->work.queued);
	ctx->work.state = CTX_FREE;
	atomic_set(&ctx->work.task_state, STATE_RUNNING);
	init_waitqueue_head(&ctx->work.wq);

	ctx->f = file;

	file->private_data = ctx;

	return nonseekable_open(inode, file);
}

static int dicedev_release(struct inode *inode, struct file *file)
{
	struct dicedev_ctx *ctx = file->private_data;

	mutex_lock(&ctx->work.mutex);
	/* At this point all work should be proccessed. */
	BUG_ON(ctx->work.state == CTX_RUNNING);
	/* Buffers should clean work the queued. */
	BUG_ON(!list_empty(&ctx->work.queued));

	mutex_unlock(&ctx->work.mutex);

	mutex_destroy(&ctx->work.mutex);

	kfree(ctx);

	return 0;
}

static long
dicedev_ioctl_create_set(struct dicedev_ctx *ctx,
			 struct dicedev_ioctl_create_set __user *arg)
{
	struct dicedev_ioctl_create_set create_set;
	struct dicedev_buffer *buf;

	long err;
	long fd;
	size_t num_pages;
	size_t page;
	size_t total_alloc_size;
	int i;

	if (copy_from_user(&create_set, arg, sizeof(create_set))) {
		err = -EFAULT;
		goto err_copy;
	}

	// Max number of pips is 32.
	if (create_set.allowed >= (1llu << 32) || create_set.size < 0 ||
	    create_set.size > DICEDEV_MAX_BUFFER_SIZE ||
	    create_set.size % sizeof(struct dice) != 0) {
		err = -EINVAL;
		goto err_args;
	}

	num_pages = create_set.size / DICEDEV_PAGE_SIZE;
	if (num_pages * DICEDEV_PAGE_SIZE < create_set.size) {
		num_pages++;
	}

	total_alloc_size =
		sizeof(*buf) + sizeof(*buf->page_table.pages) * num_pages +
		8; // 8 is just some additional space in casse on misalignemnts.
	if (!(buf = kvzalloc(total_alloc_size, GFP_KERNEL))) {
		err = -ENOMEM;
		goto err_kmalloc;
	}

	buf->page_table.pages = (typeof(buf->page_table.pages))(buf + 1);

	buf->size = create_set.size;
	buf->num_pages = num_pages;

	if (!(buf->page_table.page_data = dma_alloc_coherent(
		      &ctx->dev->pdev->dev, DICEDEV_PAGE_SIZE,
		      &buf->page_table.data_dma, GFP_KERNEL))) {
		goto err_page_table;
	}
	memset(buf->page_table.page_data, 0, DICEDEV_PAGE_SIZE);

	for (page = 0; page < num_pages; page++) {
		if (!(buf->page_table.pages[page].cmds = dma_alloc_coherent(
			      &ctx->dev->pdev->dev, DICEDEV_PAGE_SIZE,
			      &buf->page_table.pages[page].data_dma,
			      GFP_KERNEL))) {
			goto err_pages;
		}
	}

	// Fill only present pages in page table. Other entries are zeroed in allocation.
	for (i = 0; i < num_pages; i++) {
		buf->page_table.page_data[i] =
			cpu_to_le32(((buf->page_table.pages[i].data_dma >>
				      DICEDEV_PAGE_SHIFT)
				     << 4) |
				    1);
	}

	buf->ctx = ctx;

	atomic_set(&buf->seed, DICEDEV_DEFAULT_SEED);

	buf->allowed = (u32)create_set.allowed;

	buf->results.read = 0;

	mutex_init(&buf->results.mutex);

	fd = anon_inode_getfd("dicedev.buffer", &dicedev_buffer_file_ops, buf,
			      O_RDWR);
	if (fd < 0) {
		err = fd;
		goto err_inode;
	}

	(void)get_file(ctx->f);

	return fd;

err_inode:
err_pages:
	dma_free_coherent(&ctx->dev->pdev->dev, DICEDEV_PAGE_SIZE,
			  buf->page_table.page_data, buf->page_table.data_dma);
	for (; page != 0; page--) {
		dma_free_coherent(&ctx->dev->pdev->dev, DICEDEV_PAGE_SIZE,
				  buf->page_table.pages[page - 1].cmds,
				  buf->page_table.pages[page - 1].data_dma);
	}
	kfree(buf);
err_page_table:
err_kmalloc:
err_args:
err_copy:
	return err;
}

static long dicedev_ioctl_enable_seed_increment(
	struct dicedev_ctx *ctx,
	struct dicedev_ioctl_seed_increment __user *arg)
{
	struct dicedev_ioctl_seed_increment increment;
	long err = 0;
	unsigned long flags;

	if (copy_from_user(&increment, arg, sizeof(increment))) {
		return -EFAULT;
	}

	spin_lock_irqsave(&ctx->dev->work.slock, flags);

	if (ctx->dev->work.suspended) {
		err = -EIO;
	} else {
		dicedev_iow(ctx->dev, DICEDEV_INCREMENT_SEED,
			    !!increment.enable_increment);
	}

	spin_unlock_irqrestore(&ctx->dev->work.slock, flags);

	return err;
}

static long dicedev_ioctl_wait(struct dicedev_ctx *ctx,
			       struct dicedev_ioctl_wait __user *arg)
{
	struct dicedev_ioctl_wait wait;

	if (copy_from_user(&wait, arg, sizeof(wait))) {
		return -EFAULT;
	}
	return dicedev_wait(ctx, wait.cnt);
}

static long dicedev_fd_to_buf(int fd, struct dicedev_buffer **buf_p)
{
	struct fd fd_struct = fdget(fd);
	long err = 0;

	if (fd_struct.file == NULL) {
		err = -EBADF;
		goto exit;
	}
	if (fd_struct.file->f_op != &dicedev_buffer_file_ops) {
		err = -EINVAL;
		goto exit;
	}
	*buf_p = fd_struct.file->private_data;

exit:
	fdput(fd_struct);
	return err;
}

static long dicedev_ioctl_run(struct dicedev_ctx *ctx,
			      struct dicedev_ioctl_run __user *arg)
{
	struct dicedev_ioctl_run run;
	struct dicedev_buffer *cbuf;
	struct dicedev_buffer *buf;
	struct dicedev_task *task;
	long err;
	u32 *cmds;
	size_t size;
	size_t copied;
	size_t copied_size;
	size_t page;
	size_t in_page_offset;

	if (copy_from_user(&run, arg, sizeof(run))) {
		err = -EFAULT;
		goto err;
	}

	if ((err = dicedev_fd_to_buf(run.bfd, &buf))) {
		goto err;
	}
	if ((err = dicedev_fd_to_buf(run.cfd, &cbuf))) {
		goto err;
	}

	if (buf->ctx != ctx || cbuf->ctx != ctx) {
		err = -EINVAL;
		goto err;
	}

	if (run.addr % sizeof(u32) != 0 || run.size % sizeof(u32) != 0 ||
	    run.addr > cbuf->size) {
		err = -EINVAL;
		goto err;
	}

	size = min(run.size, (u32)cbuf->size - run.addr);

	cmds = kvmalloc(size, GFP_KERNEL);
	if (cmds == NULL) {
		return -ENOMEM;
	}

	for (copied = 0, page = run.addr / DICEDEV_PAGE_SIZE,
	    in_page_offset = run.addr % DICEDEV_PAGE_SIZE;
	     copied < size; copied += copied_size, page++, in_page_offset = 0) {
		copied_size =
			min(size - copied, DICEDEV_PAGE_SIZE - in_page_offset);
		memcpy((char *)cmds + copied,
		       (char *)cbuf->page_table.pages[page].cmds +
			       in_page_offset,
		       copied_size);
	}

	task = dicedev_new_cmds_work(buf, cmds, size / sizeof(u32));
	if (task == NULL) {
		err = -ENOMEM;
		goto err_new_work;
	}

	if ((err = dicedev_ctx_add_work(ctx, task))) {
		goto err_ctx_add;
	}

	return 0;

err_ctx_add:
	kfree(task);
err_new_work:
	kfree(cmds);
err:
	return err;
}

static long dicedev_ioctl(struct file *file, unsigned int cmd,
			  unsigned long arg)
{
	struct dicedev_ctx *ctx = file->private_data;

	switch (cmd) {
	case DICEDEV_IOCTL_CREATE_SET:
		return dicedev_ioctl_create_set(
			ctx, (struct dicedev_ioctl_create_set __user *)arg);
	case DICEDEV_IOCTL_ENABLE_SEED_INCREMENT:
		return dicedev_ioctl_enable_seed_increment(
			ctx, (struct dicedev_ioctl_seed_increment __user *)arg);
	case DICEDEV_IOCTL_RUN:
		return dicedev_ioctl_run(
			ctx, (struct dicedev_ioctl_run __user *)arg);
	case DICEDEV_IOCTL_WAIT:
		return dicedev_ioctl_wait(
			ctx, (struct dicedev_ioctl_wait __user *)arg);
	default:
		return -EINVAL;
	};
}

static const struct file_operations dicedev_file_ops = {
	.owner = THIS_MODULE,
	.open = dicedev_open,
	.release = dicedev_release,
	.compat_ioctl = dicedev_ioctl,
	.unlocked_ioctl = dicedev_ioctl,
};

/* PCI driver.  */

static int dicedev_probe(struct pci_dev *pdev,
			 const struct pci_device_id *pci_id)
{
	int err, i;

	/* Allocate our structure.  */
	struct dicedev_device *dev = kzalloc(sizeof *dev, GFP_KERNEL);
	if (!dev) {
		err = -ENOMEM;
		goto out_alloc;
	}
	pci_set_drvdata(pdev, dev);
	dev->pdev = pdev;

	/* Allocate a free index.  */
	mutex_lock(&dicedev_devices_lock);
	for (i = 0; i < DICEDEV_MAX_DEVICES; i++)
		if (!dicedev_devices[i])
			break;
	if (i == DICEDEV_MAX_DEVICES) {
		err = -ENOSPC; // XXX right?
		mutex_unlock(&dicedev_devices_lock);
		goto out_slot;
	}
	dicedev_devices[i] = dev;
	dev->idx = i;
	mutex_unlock(&dicedev_devices_lock);

	/* Enable hardware resources.  */
	if ((err = pci_enable_device(pdev))) {
		goto out_enable;
	}
	if ((err = dma_set_mask_and_coherent(&pdev->dev, DMA_BIT_MASK(40)))) {
		goto out_mask;
	}
	pci_set_master(pdev);

	if ((err = pci_request_regions(pdev, "dicedev"))) {
		goto out_regions;
	}

	/* Map the BAR.  */
	if (!(dev->bar = pci_iomap(pdev, 0, 0))) {
		err = -ENOMEM;
		goto out_bar;
	}

	/* Connect the IRQ line.  */
	if ((err = request_irq(pdev->irq, dicedev_isr, IRQF_SHARED, "dicedev",
			       dev))) {
		goto out_irq;
	}

	dev->work.suspended = false;
	dev->work.encountered_error = false;
	INIT_LIST_HEAD(&dev->work.queued);
	dev->work.running = NULL;
	init_waitqueue_head(&dev->work.free_wq);
	spin_lock_init(&dev->work.slock);
	dev->work.work_queue =
		alloc_ordered_workqueue("[dice%d-cmwq]", 0, dev->idx);
	INIT_WORK(&dev->work.work, dicedev_work_handler);

	dicedev_dev_start(dev);

	/* We're live. Let's export the cdev.  */
	cdev_init(&dev->cdev, &dicedev_file_ops);
	if ((err = cdev_add(&dev->cdev, dicedev_devno + dev->idx, 1)))
		goto out_cdev;

	/* And register it in sysfs.  */
	dev->dev = device_create(&dicedev_class, &dev->pdev->dev,
				 dicedev_devno + dev->idx, dev, "dice%d",
				 dev->idx);
	if (IS_ERR(dev->dev)) {
		printk(KERN_ERR "dicedev: failed to register subdevice\n");
		/* too bad. */
		dev->dev = NULL;
	}

	return 0;

out_cdev:
	dicedev_dev_stop(dev);
	free_irq(pdev->irq, dev);
out_irq:
	pci_iounmap(pdev, dev->bar);
out_bar:
	pci_release_regions(pdev);
out_regions:
out_mask:
	pci_disable_device(pdev);
out_enable:
	mutex_lock(&dicedev_devices_lock);
	dicedev_devices[dev->idx] = 0;
	mutex_unlock(&dicedev_devices_lock);
out_slot:
	kfree(dev);
out_alloc:
	return err;
}

static void dicedev_remove(struct pci_dev *pdev)
{
	struct dicedev_device *dev = pci_get_drvdata(pdev);

	if (dev->dev) {
		device_destroy(&dicedev_class, dicedev_devno + dev->idx);
	}
	cdev_del(&dev->cdev);
	dicedev_dev_stop(dev);

	free_irq(pdev->irq, dev);
	pci_iounmap(pdev, dev->bar);
	pci_release_regions(pdev);
	pci_disable_device(pdev);

	destroy_workqueue(dev->work.work_queue);

	mutex_lock(&dicedev_devices_lock);
	dicedev_devices[dev->idx] = 0;
	mutex_unlock(&dicedev_devices_lock);

	kfree(dev);
}

static int dicedev_suspend(struct pci_dev *pdev, pm_message_t state)
{
	unsigned long flags;
	struct dicedev_device *dev = pci_get_drvdata(pdev);

	spin_lock_irqsave(&dev->work.slock, flags);
	while (dev->work.running != NULL) {
		spin_unlock_irqrestore(&dev->work.slock, flags);
		wait_event(dev->work.free_wq, dev->work.running == NULL);
		spin_lock_irqsave(&dev->work.slock, flags);
	}
	dev->work.suspended = true;
	spin_unlock_irqrestore(&dev->work.slock, flags);

	dicedev_dev_stop(dev);

	return 0;
}

static int dicedev_resume(struct pci_dev *pdev)
{
	struct dicedev_device *dev = pci_get_drvdata(pdev);
	unsigned long flags;

	dicedev_dev_start(dev);

	spin_lock_irqsave(&dev->work.slock, flags);
	dev->work.suspended = true;
	spin_unlock_irqrestore(&dev->work.slock, flags);

	wake_up_interruptible_sync(&dev->work.free_wq);

	return 0;
}

static struct pci_device_id dicedev_pciids[] = {
	{ PCI_DEVICE(DICEDEV_VENDOR_ID, DICEDEV_DEVICE_ID) },
	{ 0 }
};

static struct pci_driver dicedev_pci_driver = {
	.name = "dicedev",
	.id_table = dicedev_pciids,
	.probe = dicedev_probe,
	.remove = dicedev_remove,
	.suspend = dicedev_suspend,
	.resume = dicedev_resume,
};

/* Init & exit. */

static int dicedev_init(void)
{
	int err;
	if ((err = alloc_chrdev_region(&dicedev_devno, 0, DICEDEV_MAX_DEVICES,
				       "dicedev")))
		goto err_chrdev;
	if ((err = class_register(&dicedev_class)))
		goto err_class;
	if ((err = pci_register_driver(&dicedev_pci_driver)))
		goto err_pci;
	return 0;

err_pci:
	class_unregister(&dicedev_class);
err_class:
	unregister_chrdev_region(dicedev_devno, DICEDEV_MAX_DEVICES);
err_chrdev:
	return err;
}

static void dicedev_exit(void)
{
	pci_unregister_driver(&dicedev_pci_driver);
	class_unregister(&dicedev_class);
	unregister_chrdev_region(dicedev_devno, DICEDEV_MAX_DEVICES);
}

module_init(dicedev_init);
module_exit(dicedev_exit);
