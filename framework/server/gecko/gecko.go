package gecko

import "net/http"

// Gecko is an HTTP server framework.
type Gecko struct {
	addr   string
	router *Router
}

// HandlerFunc is a processor for http server.
type HandlerFunc func(sess *Session)

// New creates a new Gecko server.
func New(addr string) *Gecko {
	return &Gecko{addr: addr, router: NewRouter()}
}

// Run the Gecko server.
func (g *Gecko) Run() {
	http.ListenAndServe(g.addr, g)
}

// ServeHTTP implements http.Handler
func (g *Gecko) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, params := g.SearchHandler(HTTPMethod(Get), r.URL.Path)
	if handler != nil {
		handler(NewSession(w, r, params))
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	} else {
		w.WriteHeader(400)
		w.Write([]byte("NOT-FOUND"))
	}

}

// GET registers a [GET] method url.
func (g *Gecko) GET(url string, handler HandlerFunc) {
	g.router.insert(Get, url, handler)
}

// POST registers a [POST] method url.
func (g *Gecko) POST(url string, handler HandlerFunc) {
	g.router.insert(Post, url, handler)
}

// PUT registers a [PUT] method url.
func (g *Gecko) PUT(url string, handler HandlerFunc) {
	g.router.insert(Put, url, handler)
}

// DELETE registers a [DELETE] method url.
func (g *Gecko) DELETE(url string, handler HandlerFunc) {
	g.router.insert(Delete, url, handler)
}
