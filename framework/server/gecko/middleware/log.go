package middleware

import (
	"github.com/galaxyzeta/framework/server/gecko"
	"github.com/galaxyzeta/logger"
)

// Log incoming request and its URI. This can be used as a middleware test.
func Log(sess *gecko.Session) {
	log := logger.New("Gecko")
	log.Infof("Incoming request: [%s] %s", sess.Req.Method, sess.Req.RequestURI)
	sess.Next()
}
