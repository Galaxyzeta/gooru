package main

import (
	"fmt"

	"galaxyzeta.com/framework/server/gecko"
	"galaxyzeta.com/framework/server/gecko/middleware"
)

func main() {
	middleWare()
}

func defaultRouter() {
	g := gecko.NewGecko("localhost:8080")
	g.GET("/index", func(_ *gecko.Session) {
		fmt.Println("[GET] /index")
	})
	g.POST("/index", func(_ *gecko.Session) {
		fmt.Println("[POST] /index")
	})
	g.GET("/ind", func(_ *gecko.Session) {
		fmt.Println("[GET] /ind")
	})
	g.GET("/admin", func(_ *gecko.Session) {
		fmt.Println("[GET] /admin")
	})
	g.GET("/index/pos", func(_ *gecko.Session) {
		fmt.Println("[GET] /index/pos")
	})
	g.GET("/ind/:a/:b/*file", func(sess *gecko.Session) {
		fmt.Println("[GET] /ind/:a/:b/*file")
		fmt.Println(sess.Params)
	})
	g.GET("/index/:id/asd", func(sess *gecko.Session) {
		fmt.Println("[GET] /index/:id/asd")
		fmt.Println(sess.Params)
	})
	g.GET("/favicon", func(_ *gecko.Session) {
		fmt.Println("[GET] /favicon")
	})
	g.GET("/file/*filename", func(sess *gecko.Session) {
		fmt.Println("[GET] /file/*filename")
		fmt.Println(sess.Params)
	})
	g.Run()
}

func middleWare() {
	g := gecko.NewGecko("localhost:8080")
	g1 := g.Group("/v1")
	g1.Use(middleware.Log)
	g1.GET("/index", func(sess *gecko.Session) {
		fmt.Println("[GET] /v1/index")
		sess.Next()
	})
	g1.GET("/hello/:name/:password", func(sess *gecko.Session) {
		fmt.Println(sess.Params)
	})
	g.Run()
}
