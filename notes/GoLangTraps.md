# Go 语言陷阱

1. 数组切片是一个值类型。

```go
package main

func modify(a [2]int) {
	a[0] = 100
}

func main() {
	a := [2]int{1,2}
	modify(a)
	fmt.Println(a)	// 1 2
}
```

