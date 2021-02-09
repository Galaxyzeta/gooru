package main

import galaxyserver "galaxyzeta.com/framework/server/simple"

// Pirate version of Echo framework
func main() {
	server := galaxyserver.New()
	server.GET("/debug", func(c *galaxyserver.Context) interface{} {
		return "hello"
	})
	server.Start("localhost", 18080)
}
