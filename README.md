# Gooru

Gooru (Guru) 用于存放无聊程序员 Galaxyzeta 造的 golang 轮子。欢迎找出各种 bug 或提交新轮子！

Gooru is a playground. It intends to hold any Golang stuff written by boring programmer Galaxyzeta. Feel free to propose issues and make contributions !

内容简介：
- algo：算法
	- compare：
		- comparator：比较器，强行实现了任意基本类型的比较，此外还提供一些常用比较函数。
	- sort：排序算法轮子。
- concurrency：并发工具包
	- async：异步任务。
		- delayedjob：使用基于轮询的 `TimerWheel` 完成异步任务。存在性能问题，考虑改写后废弃原版。
		- intervaljob：基于 `time.Ticker` 完成周期性任务。
	- spinLock：自旋锁，直到任务完成，一直非阻塞自旋。
	- semaphore：基于 `Mutex` 和 `chan` 分别实现了信号量。
	- goroutine：获得关于 `goroutine` 的一些信息。
	- synchronizer：
		- aqs：抄了 java AQS，提供了 `IAQS` 接口，用于实现各种同步容器。
		- reentrantlock：可重入锁，golang 提供的 `mutex` 是不可重入的。
	- waiter: 对 `WaitGroup` 的二次封装，用于等待若干个函数的执行完毕。
- ds：数据结构包
	- list
		- singlelinkedlist：单向线程不安全链表。可用作栈或队列，性能比系统提供的 list 要好。
		- doublelinkedlist：双向线程不安全链表。
		- list_interface：提供了`List` `Stack` `Queue` 的接口定义。
	- map
		- hashmap：二次封装的，线程不安全的 `map`。
		- map：提供 `Map` 接口定义。
		- safe_hashmap：通过互斥锁实现的线程安全 `map`，性能比 `sync.Map` 略差。
		- lru: 提供基于 `LRU` 淘汰策略的缓存 `map`。
	- tree
		- bst：二叉搜索树，线程不安全，提供了基于 bst 的 `BSTMap`，保证插入元素的有序性。
	- deprecated：
		- hashmap：手动实现散列表，使用链地址法解决哈希冲突。性能不如 `map`。
- dp：设计模式。
- ioc：简陋的依赖注入框架，可以循环依赖。
- logger：日志打印。
- server：各种网络模型的尝试。
- test：滚轮子测试。
- tutorial：入门 golang。
- util：工具。
	- common：一些缩写和转化工具。
	- alias：一些函数的缩写。例如 `P1Consumer` 表示一个入参的无返回值函数。
	- benchmark：用于效率测试。
	- assert：断言工具，断言失败会导致 `panic`。

一些大胆想法：
- 实现一个 Stream 操纵任何切片进行变形。
- 把 python enumerate 和 zip 等函数抄过来