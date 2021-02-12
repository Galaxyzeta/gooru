# GoCaching

GoCaching 是一个分布式缓存服务，提供 API 前端，以及缓存服务后端程序。

GoCaching is a distributed caching server. It provides an API front-end service, as well as caching back-end service.

## Feature

1. 基于一致性哈希的 key 值路由。
2. 心跳机制保障后端服务的可用性。
3. 使用基于 LRU 的淘汰策略缓存 key。
4. 使用 Singleflight 机制防止一次发起多个请求。