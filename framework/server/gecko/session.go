package gecko

import "net/http"

// Session represents one single communication of
type Session struct {
	Respw  http.ResponseWriter
	Req    *http.Request
	Params map[string]string
}

// NewSession creates a new session.
func NewSession(w http.ResponseWriter, req *http.Request, params map[string]string) *Session {
	return &Session{Respw: w, Req: req, Params: params}
}
