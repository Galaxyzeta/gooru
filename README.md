# Gooru

Gooru(Guru) 用于存放无聊程序员 Galaxyzeta 造的 golang 轮子

Gooru is a scratch book. It intends to hold any Golang stuff written by boring programmer Galaxyzeta.

内容简介：

- concurrency：并发工具包
	- delayedjob：使用基于轮询的 `TimerWheel` 完成异步任务。
	- spinLock：自旋锁，直到任务完成，一直非阻塞自旋。
	- semaphore：基于 `Mutex` 和 `chan` 分别实现了信号量。
	- goroutine：获得关于 `goroutine` 的一些信息。
	- synchronizer：
		- aqs：抽象同步器，提供了 `IAQS` 接口，用于实现各种同步容器。
		- reentrantlock：可重入锁，golang 提供的 `mutex` 是不可重入的。
- ds：数据结构包
- ioc：依赖注入框架
- logger：日志打印
- server：各种网络模型的尝试
- test：滚轮子测试
- tutorial：入门 golang