package main

import (
	"fmt"

	"galaxyzeta.com/ioc"
)

// A test
type A struct {
	Name string
	Age  int
	B    *B
}

// B test
type B struct {
	Address string
}

// C test
type C struct {
	D *D
	E *E
}

// D test
type D struct {
	C *C
	E *E
}

// E test
type E struct {
	C *C
	D *D
}

func main() {
	// Test normal dependency
	ctx := ioc.New()
	ctx.Add("hello", "world")
	ctx.Add("bean", 123)
	ctx.Add("gopher", "neo city")
	ctx.Add("A", &A{})
	ctx.Add("B", &B{})
	ctx.AddDep("A", "Name", "hello")
	ctx.AddDep("A", "Age", "bean")
	ctx.AddDep("A", "B", "B")
	ctx.AddDep("B", "Address", "gopher")

	// Test cycle dependency
	ctx.Add("C", &C{})
	ctx.Add("D", &D{})
	ctx.Add("E", &E{})
	ctx.AddDep("C", "D", "D")
	ctx.AddDep("C", "E", "E")
	ctx.AddDep("D", "C", "C")
	ctx.AddDep("D", "E", "E")
	ctx.AddDep("E", "D", "D")
	ctx.AddDep("E", "C", "C")

	// Do Autowire !
	ctx.Refresh()

	fmt.Println(ctx.Get("A"))
	fmt.Println(ctx.Get("B"))
	fmt.Println(ctx.Get("C"))
	fmt.Println(ctx.Get("D"))
	fmt.Println(ctx.Get("E"))
}
