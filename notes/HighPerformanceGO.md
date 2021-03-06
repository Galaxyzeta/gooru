# 高性能 GO 阅读记录

## 使用内存对齐提升性能

原理：缓存行共享，程序的局部性原理。

内存对齐的对齐倍数是指内存对齐时，其在内存中的开始位置必须是对齐倍数 n 的整数倍。计算：

|类型|对齐倍数|
|----|----|
|int8|1|
|int16|2|
|int32|4|
|int64|8|

例子：

```go
type struct A {
	int8
	int16
	int32
}
type struct B {
	int8
	int32
	int16
}
```

前者内存分布（8B）

🟦⬜🟦🟦🟦🟦🟦🟦

<p>int8 &nbsp int16 &nbsp int32</p>

后者内存分布（12B）

🟦⬜⬜⬜🟦🟦🟦🟦 | 🟦🟦⬜⬜

<p>int8 &nbsp int32 &nbsp int32</p>

## 反射对性能的影响

使用反射进行 new 操作，性能差 1.5 倍。

使用 `Field` 遍历注入，性能比直接赋值差 10 倍。

使用 `FieldByName` 按名注入，性能比直接赋值差 100 倍。

总结：能不反射就不要用反射。尽量避免在运行时执行反射。

## 使用空结构体节约内存

场景：

1. 作为控制作用的 `channel`，使用 `struct{}` 传值。
2. 没有成员，只有方法的结构体，使用 `type XXX struct{}` 定义。
3. 实现集合，使用 `map[interface{}]struct{}` 。

## 字符串拼接和性能比较

1. 使用 + 号，或者使用 `fmt.Sprintf()`。性能最差，因为字符串是不可变的，这样操作每次都会生成一个新的对象。
2. 使用 `strings.Builder` 拼接最好，其基于 `[]byte` 进行操作。 

## 逃逸分析

> 编译器决定内存分配位置的方式，就称之为逃逸分析(escape analysis)。逃逸分析由编译器完成，作用于编译阶段。

1. 传递指针会导致局部变量从栈分配逃逸到堆分配。
2. 动态类型 `interface{}` 会在堆上分配。
3. 每个内核线程能占据的最大栈空间是有限制的。内存不足会导致逃逸。
4. 闭包函数中引用的局部变量会逃逸到堆空间。

针对逃逸分析的优化：传递指针 VS 传递值：

1. 对于小型 `struct` 传递值，否则因为指针逃逸，会加大 GC 压力。
2. 对于复杂类型传递指针。

## 死码消除

Dead Code Elimination 是指编译器会针对不必要的语句进行内联（inline）或者消除。例如常量的消除，短函数内联等等。

针对 DCE，可以做的：

1. 能用常量表示的不要用变量。
2. 使用常量开关设计出 `Debug` 版本或 `Release` 版本，占空间少。例如：

```go
package main
const debug = true
func main() {
	if debug == true {
		fmt.Println("Debug mode")
	}
}
```

这里的 if 会被 DCE，因此编译后二进制文件的大小和 Release 版本是一样的。