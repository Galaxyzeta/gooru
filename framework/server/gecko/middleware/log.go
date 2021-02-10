package middleware

import (
	"galaxyzeta.com/framework/server/gecko"
	"galaxyzeta.com/logger"
)

// Log incoming request and its URI. This can be used as a middleware test.
func Log(sess *gecko.Session) {
	log := logger.New("Gecko")
	log.Infof("Incoming request: [%s] %s", sess.Req.Method, sess.Req.RequestURI)
	sess.Next()
}
