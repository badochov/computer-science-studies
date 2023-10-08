# Dicedev driver

Solution puts write and run tasks on a per context queue. 

Then on ioctl wait the tasks are split into groups of task that have the same slot and buffer and put to the per device ordered work queue.
Work queue executes task from context one context at a time to be sure which context caused the error. Either all task from the context are valid or all of them result in -EIO. Fence is added following all tasks from given context. Afte fence isr is received driver queues next context's tasks. Only slot 0 is used by the drivers.


The resoning for this architertuce is that it provides as predictable interface as it can with this devices as the user always know in which place in the buffer the results are and can easily avoid results overwriting each other.