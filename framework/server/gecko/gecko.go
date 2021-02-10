package gecko

import (
	"net/http"

	"galaxyzeta.com/concurrency/waiter"
)

// Gecko is an HTTP server framework.
type Gecko struct {
	*GroupRouter
	addr   string
	router *Router
	groups []*GroupRouter
}

// HandlerFunc processes each session from client.
type HandlerFunc func(sess *Session)

// MiddleWareFunc is an alias for HandlerFunc.
type MiddleWareFunc func(sess *Session)

// ParamMap represents a map of params.
type ParamMap map[string]string

// Session represents one single communication of
type Session struct {
	Respw           http.ResponseWriter
	Req             *http.Request
	Params          ParamMap
	Groups          []*GroupRouter
	middleWares     []MiddleWareFunc
	middleWareIndex int
}

// newSession creates a new session.
func newSession(w http.ResponseWriter, req *http.Request, params ParamMap, groups []*GroupRouter) *Session {
	ret := &Session{Respw: w, Req: req, Params: params, Groups: groups, middleWares: make([]MiddleWareFunc, 0)}
	for _, gp := range ret.Groups {
		for _, m := range gp.middleWares {
			ret.middleWares = append(ret.middleWares, m)
		}
	}
	return ret
}

// Next calls the next registered middleware.
func (s *Session) Next() {
	if s.middleWareIndex < len(s.middleWares) {
		k := s.middleWareIndex
		s.middleWareIndex++
		s.middleWares[k](s)
	}
}

// NewGecko creates a new Gecko server.
func NewGecko(addr string) *Gecko {
	return &Gecko{addr: addr, router: NewRouter(), groups: make([]*GroupRouter, 0)}
}

// Group returns a group instance with a given prefix.
func (g *Gecko) Group(prefix string) *GroupRouter {
	return &GroupRouter{gecko: g, prefix: prefix}
}

// Run the Gecko server.
func (g *Gecko) Run() {
	http.ListenAndServe(g.addr, g)
}

// ServeHTTP implements http.Handler
func (g *Gecko) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 1. find handler.
	handler, params, groups := g.SearchHandler(HTTPMethod(Get), r.URL.Path)
	// 2. execute middleware.
	if handler != nil {
		// 3. execute handler.
		wg := waiter.New() // Use Sync.Once might be more efficient.
		wg.AddAndRun(func() {
			handler(newSession(w, r, params, groups))
		})
		wg.Wait()
		// 4. do view resolution.
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	} else {
		w.WriteHeader(400)
		w.Write([]byte("NOT-FOUND"))
	}

}

// GET registers a [GET] method url.
func (g *Gecko) GET(url string, handler HandlerFunc) {
	g.router.insert(Get, url, handler, g.GroupRouter)
}

// POST registers a [POST] method url.
func (g *Gecko) POST(url string, handler HandlerFunc) {
	g.router.insert(Post, url, handler, g.GroupRouter)
}

// PUT registers a [PUT] method url.
func (g *Gecko) PUT(url string, handler HandlerFunc) {
	g.router.insert(Put, url, handler, g.GroupRouter)
}

// DELETE registers a [DELETE] method url.
func (g *Gecko) DELETE(url string, handler HandlerFunc) {
	g.router.insert(Delete, url, handler, g.GroupRouter)
}
