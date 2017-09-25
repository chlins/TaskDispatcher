### TaskDispatcher

一个简单的任务分发器，采用随机抢占分发模式。

架构图：



![](sys.png)

例如： 有10个任务，每个任务下又包含多个子任务，任务会被压到栈中，此时用户定义dispatcher中worker的数量，空闲状态的worker会去pop任务执行，直到所有Task完成。（欢迎contribute 23333333333......）