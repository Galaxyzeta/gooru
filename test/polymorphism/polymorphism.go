package main

import "fmt"

type IBehaviour interface {
	DoSomething()
}

type A struct {
	data int
}

type B struct{ A }

func (a *A) DoSomething() {
	a.data = 1
	fmt.Println("Hello")
}

func main() {
	var ib IBehaviour = &B{}
	var sb = ib.(*B)
	fmt.Println(sb.data)
	fmt.Println(sb.A.data)
	ib.DoSomething()
	fmt.Println(sb.data)
	fmt.Println(sb.A.data)
}
