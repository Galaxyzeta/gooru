package main

import (
	"flag"

	"galaxyzeta.com/framework/cache/gocaching"
	"galaxyzeta.com/logger"
)

var localDB map[string]string = map[string]string{
	"Tom": "hello",
	"Sam": "world",
	"Tim": "haha",
}

var localGetter gocaching.LocalGetter = func(key string) interface{} {
	return localDB[key]
}

func main() {
	var host string
	var api bool
	flag.StringVar(&host, "host", "", "The host of current cache server.")
	flag.BoolVar(&api, "api", false, "Whether to create APIServer.")
	flag.Parse()
	peers := []string{
		"localhost:18080", "localhost:18081", "localhost:18082",
	}
	log := logger.New("Main")
	cache := gocaching.NewGoCaching(host, 64, localGetter)
	if api {
		log.Infof("Running api server on %s", host)
		cache.RunAPIServer(host, peers...)
	}
	log.Infof("Running cache server on %s", host)
	cache.RunCacheServer()
	log.Infof("Closed")
}
