
package gocaching

import (
	"net/http"

	hashmap "github.com/galaxyzeta/ds/map"
	"github.com/galaxyzeta/logger"
)

// GoCaching is a distributed caching server.
type GoCaching struct {
	host   string
	lru    *hashmap.LRUCache
	getter LocalGetter
	logger *logger.Logger
}

// LocalGetter represents a method for getting primitive data from local storage.
type LocalGetter func(key string) interface{}

var heartBeatRoute = "/system/beat"

// NewGoCaching returns an instance of GoCaching.
func NewGoCaching(host string, size int, getter LocalGetter) *GoCaching {
	return &GoCaching{host: host,
		lru:    hashmap.NewLRUCache(size),
		getter: getter,
		logger: logger.New("GoCache"),
	}
}

// RunCacheServer boots up the caching server.
func (g *GoCaching) RunCacheServer() {
	http.ListenAndServe(g.host, g)
}

// ServeHTTP implements http.Handler.
func (g *GoCaching) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func(rw http.ResponseWriter, _ *http.Request) {
		if k := recover(); k != nil {
			write404(rw, "Internal Server Error.")
		}
	}(w, r)

	// Handle heart-beat event
	if r.URL.Path == heartBeatRoute {
		write200(w, "Heart-beat OK")
		return
	}

	// Handle normal key-caching event.
	key := r.URL.Path[1:]
	res, err := g.getKeyLocal(key)
	g.logger.Debugf("Val received from remote, res = %v", res)
	if err != nil {
		g.logger.Fatalf("While getting key from local, an error occured.")
		write404(w, "While getting key from local, an error occured.")
		return
	}
	if res == "" {
		g.logger.Fatalf("Key Not Exist! ")
		write404(w, "Key Not Exist! ")
		return
	}
	write200(w, res)
}

// getKeyLocal is the core logic for retrieving cache.
func (g *GoCaching) getKeyLocal(key string) (string, error) {
	if ret := g.lru.Get(key); ret != nil {
		// LRU Cache hit. Return result.
		g.logger.Debugf("Key %v exists in LRU", ret)
		return ret.(string), nil
	}
	// LRU cache miss. Use getter to lookup result, and then cache it.
	res := g.getter(key).(string)
	g.logger.Debugf("Key %v not exists in LRU, trying to get from local storage, res = %v", key, res)
	g.lru.Put(key, res)
	return res, nil
}

func write404(w http.ResponseWriter, data string) {
	w.WriteHeader(404)
	w.Write([]byte(data))
}

func write200(w http.ResponseWriter, data string) {
	w.Write([]byte(data))
}
