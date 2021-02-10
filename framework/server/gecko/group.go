package gecko

// GroupRouter is a logical combination of routers. Middlewares can be registered on a group to work on all sub routes.
type GroupRouter struct {
	gecko       *Gecko
	prefix      string
	middleWares []MiddleWareFunc
}

// Use attaches a middleware function to the certain group, which will be executed before each session's handler function.
func (g *GroupRouter) Use(middlewareFunc ...MiddleWareFunc) {
	g.middleWares = append(g.middleWares, middlewareFunc...)
}

func (g *GroupRouter) insert(method HTTPMethod, url string, handler HandlerFunc) {
	g.gecko.router.insert(method, url, handler, g)
}

// GET registers a [GET] method url.
func (g *GroupRouter) GET(url string, handler HandlerFunc) {
	g.gecko.router.insert(Get, url, handler, g)
}

// POST registers a [POST] method url.
func (g *GroupRouter) POST(url string, handler HandlerFunc) {
	g.gecko.router.insert(Post, url, handler, g)
}

// PUT registers a [PUT] method url.
func (g *GroupRouter) PUT(url string, handler HandlerFunc) {
	g.gecko.router.insert(Put, url, handler, g)
}

// DELETE registers a [DELETE] method url.
func (g *GroupRouter) DELETE(url string, handler HandlerFunc) {
	g.gecko.router.insert(Delete, url, handler, g)
}
