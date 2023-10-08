#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/ioctl.h>
#include <linux/device.h>
#include <linux/uaccess.h>
#include <linux/fs.h>
#include <linux/cdev.h>

MODULE_LICENSE("GPL");

#define HELLO_IOCTL_SET_REPEATS _IO('H', 0x00)
#define HELLO_MAX_REPEATS 0x100

static const char hello_default[] = "Hello, world!\n";
static char *hello_reply;

static int hello_max_len = sizeof(hello_default);
static int hello_len;
module_param_named(bufsize, hello_max_len, int, 0);
static long hello_repeats = 1;
static dev_t hello_major;
static struct cdev hello_cdev;
static struct cdev hello_once_cdev;
static struct class hello_class = {
	.name = "hello",
	.owner = THIS_MODULE,
};
static struct device *hello_device;
static struct device *hello_once_device;

static ssize_t hello_once_read(struct file *file, char __user *buf, size_t count, loff_t *filepos)
{
	loff_t pos = *filepos;
	if (pos >= hello_len || pos < 0)
		return 0;
	if (count > hello_len - pos)
		count = hello_len - pos;
	if (copy_to_user(buf, hello_reply + pos, count))
		return -EFAULT;
	*filepos = pos + count;
	return count;
}

static ssize_t hello_read(struct file *file, char __user *buf, size_t count, loff_t *filepos)
{
	size_t file_len = hello_len * hello_repeats;
	loff_t pos = *filepos;
	loff_t end;
	if (pos >= file_len || pos < 0)
		return 0;
	if (count > file_len - pos)
		count = file_len - pos;
	end = pos + count;
	while (pos < end)
		if (put_user(hello_reply[pos++ % hello_len], buf++))
			return -EFAULT;
	*filepos = pos;
	return count;
}

static ssize_t hello_write(struct file *file, char const __user *buf, size_t count, loff_t *filepos)
{
	loff_t pos = *filepos;
	printk(KERN_ALERT "WRITE %zu\n", count);
	if (pos + count > hello_max_len || pos < 0)
		return 0;
	if (copy_from_user(hello_reply + pos, buf, count))
	{
		return -EFAULT;
	}
	hello_len = pos + count;
	hello_reply[hello_len] = '\0';
	*filepos += count;
	return count;
}

static long hello_ioctl(struct file *file, unsigned int cmd, unsigned long arg)
{
	if (cmd != HELLO_IOCTL_SET_REPEATS)
		return -ENOTTY;
	if (arg > HELLO_MAX_REPEATS)
		return -EINVAL;
	hello_repeats = arg;
	return 0;
}

static int hello_open(struct inode *ino, struct file *filep)
{
	return 0;
}

static int hello_release(struct inode *ino, struct file *filep)
{
	return 0;
}

static struct file_operations hello_once_fops = {
	.owner = THIS_MODULE,
	.read = hello_once_read,
	.open = hello_open,
	.release = hello_release,
};

static struct file_operations hello_fops = {
	.owner = THIS_MODULE,
	.read = hello_read,
	.open = hello_open,
	.unlocked_ioctl = hello_ioctl,
	.write = hello_write,
	.compat_ioctl = hello_ioctl,
	.release = hello_release,
};

static int hello_init(void)
{
	int err;

	if ((hello_reply = kmalloc(hello_max_len, GFP_KERNEL)) == NULL)
	{
		err = -ENOMEM;
		goto err_kmalloc;
	}

	memset(hello_reply, 0, hello_max_len);
	strncpy(hello_reply, hello_default, hello_max_len - 1);

	hello_len = strlen(hello_reply);
	printk(KERN_ALERT "HELLO_LEN %d\n", hello_len);
	printk(KERN_ALERT "HELLO_MAX_LEN %d\n", hello_max_len);

	if ((err = alloc_chrdev_region(&hello_major, 0, 2, "hello")))
		goto err_alloc;
	cdev_init(&hello_cdev, &hello_fops);
	if ((err = cdev_add(&hello_cdev, hello_major, 1)))
		goto err_cdev;
	cdev_init(&hello_once_cdev, &hello_once_fops);
	if ((err = cdev_add(&hello_once_cdev, hello_major + 1, 1)))
		goto err_cdev_2;
	if ((err = class_register(&hello_class)))
		goto err_class;
	hello_device = device_create(&hello_class, 0, hello_major, 0, "hello");
	if (IS_ERR(hello_device))
	{
		err = PTR_ERR(hello_device);
		goto err_device;
	}
	hello_once_device = device_create(&hello_class, 0, hello_major + 1, 0, "hello_once");
	if (IS_ERR(hello_once_device))
	{
		err = PTR_ERR(hello_once_device);
		goto err_device_2;
	}

	return 0;

err_device_2:
	device_destroy(&hello_class, hello_major);
err_device:
	class_unregister(&hello_class);
err_class:
	cdev_del(&hello_once_cdev);
err_cdev_2:
	cdev_del(&hello_cdev);
err_cdev:
	unregister_chrdev_region(hello_major, 2);
err_alloc:
	kfree(hello_reply);
	hello_reply = NULL;
err_kmalloc:
	return err;
}

static void hello_cleanup(void)
{
	device_destroy(&hello_class, hello_major + 1);
	device_destroy(&hello_class, hello_major);
	class_unregister(&hello_class);
	cdev_del(&hello_once_cdev);
	cdev_del(&hello_cdev);
	unregister_chrdev_region(hello_major, 2);
	kfree(hello_reply);
}

module_init(hello_init);
module_exit(hello_cleanup);
