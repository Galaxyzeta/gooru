package main

import (
	"fmt"

	"galaxyzeta.com/framework/server/gecko"
)

func main() {
	g := gecko.New("localhost:8080")
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
	g.GET("/ind/:a/:b", func(sess *gecko.Session) {
		fmt.Println("[GET] /ind/:a/:b")
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
