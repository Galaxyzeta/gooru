# Gooru

Gooru (Guru) 用于存放无聊程序员 Galaxyzeta 造的 golang 轮子

Gooru is a scratch book. It intends to hold any Golang stuff written by boring programmer Galaxyzeta.

内容简介：
- algo：算法
	- comparator：比较器，强行实现了任意基本类型的比较
	- sort：排序算法轮子
- concurrency：并发工具包
	- delayedjob：使用基于轮询的 `TimerWheel` 完成异步任务。
	- spinLock：自旋锁，直到任务完成，一直非阻塞自旋。
	- semaphore：基于 `Mutex` 和 `chan` 分别实现了信号量。
	- goroutine：获得关于 `goroutine` 的一些信息。
	- synchronizer：
		- aqs：抄了 java AQS，提供了 `IAQS` 接口，用于实现各种同步容器。
		- reentrantlock：可重入锁，golang 提供的 `mutex` 是不可重入的。
	- waiter: 对 `WaitGroup` 的二次封装，用于等待若干个函数的执行完毕。
- ds：数据结构包
	- linkedlist：单向链表。
- dp：设计模式
- ioc：简陋的依赖注入框架，可以循环依赖。
- logger：日志打印
- server：各种网络模型的尝试
- test：滚轮子测试
- tutorial：入门 golang
- util：工具
	- abbr：一些缩写和转化工具。
	- benchmark：用于效率测试。