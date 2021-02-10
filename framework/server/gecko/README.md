# Gecko

Gecko 是一个 HTTP 开发框架的造轮子实验。它是著名 web 框架 Gin 和 Echo 的简易版本。

Gecko 是指守宫，一种蜥蜴，你也可以把它当作是 Gin + Echo 的发音组合。

Gecko is an experiment of writing an HTTP Server framework. It is a simple version of Gin and Echo framework.

The name of Gecko means a lizard. It can also be considered as a combination of Gin and Echo : )

## Feature

1. 使用了简化的 `Trie` 树作为路由查找表。

2. 支持两种 `wildcard` 路由。使用 `:` 表示一个路由参数，在路径末尾使用 `*` 表示一个任意匹配。我们支持的路由匹配：
	```go
	g.GET("/ind/:a/:b/*file", func(sess *gecko.Session) {
			fmt.Println("[GET] /ind/:a/:b/*file")
			fmt.Println(sess.Params)
		})
	```
	需要注意的是，每个父路径下只能挂载一个 `wildcard` 节点，否则会出现二义性。Gecko 不允许这么做！

3. 可以使用基于分组路由的中间件，例如：
	```go
	g := gecko.New()
	gp := g.Group("/v1")
	gp.Use(middleware.Log)
	gp.GET("/index", func(sess *gecko.Session) {
		sess.Next()	// Middleware entrance.
	}) 
	```
	以后，每个访问 `/v1` 分组的请求，都会触发 `Log` 中间件。

4. 你可以自定义中间件，只要编写一个 `MiddleWareFunc` 类型的函数即可：
	```go
	// Log incoming request and its URI. This can be used as a middleware test.
	func Log(sess *gecko.Session) {
		log := logger.New("Gecko")
		log.Infof("Incoming request: [%s] %s", sess.Req.Method, sess.Req.RequestURI)
		sess.Next()
	}
	```
	以上是一个中间件，可以使用这个中间件，自动完成日志打印工作。

5. 参数匹配。使用形如 `:name` 的 `wildcard`，当路由成功匹配时，可以通过 `sess.Params` 得到解析完毕的参数。