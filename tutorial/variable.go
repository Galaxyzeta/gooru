package tutorial

import (
	"fmt"
	"strconv"
)

// Variable :
// test basic variables
func Variable() {
	// 1. def through var
	var a string = "hello world"
	fmt.Println(a)
	// 2. def by type indication
	b := 23.456
	fmt.Println(b)
	// 3. def a lot of variables
	var (
		c int
		d float32
	)
	fmt.Println(c, d)
	// 4. var swap
	a1, b1 := "golang", "java"
	a1, b1 = b1, a1
	// 5. ptr
	var ptr = &b1
	fmt.Printf("%p %s", ptr, *ptr)
	// 6. new
	var by = new(byte)
	fmt.Println(by)
	// 7. slice
	var list [5]int
	fmt.Println(list)
}

// ByteAndRune test
func ByteAndRune() {
	var test1 uint32 = '操'
	var test2 rune = '淦'
	var test3 byte = 'a'
	var test4 uint8 = 'b'
	fmt.Printf("%c %c %c %c", test1, test2, test3, test4)
}

// SliceAndArray test
func SliceAndArray() {
	// == Definition ==
	var arr1 [3]int = [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr1)
	fmt.Println(arr2)

	// == About Slice [len] and [cap]

	// Slice is array without max limit
	// Slice = java ArrayList
	// Slice = python3 list
	var slice1 []int = arr1[2:3]
	var slice2 []int = []int{1, 2, 3}
	var slice3 []int = arr1[0:1:3]
	fmt.Printf("var = %d len = %d cap = %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("var = %d len = %d cap = %d\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("var = %d len = %d cap = %d\n", slice3, len(slice3), cap(slice3))

	// == Slice Problems ==
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6:8]
	fmt.Printf("myslice为 %d, 其长度为: %d, cap = %d\n", myslice, len(myslice), cap(myslice))

	myslice = myslice[:cap(myslice)]
	fmt.Printf("myslice为 %d, 其长度为: %d, cap = %d\n", myslice, len(myslice), cap(myslice))
	fmt.Printf("myslice的第四个元素为: %d\n", myslice[cap(myslice)-1])

	fmt.Printf("Address of myslice = %p, Address of numbers = %p\n", &myslice, &numbers4[4])

	myslice = append(myslice, 100)
	fmt.Printf("myslice为 %d, 其长度为: %d, cap = %d\n", myslice, len(myslice), cap(myslice))
	fmt.Printf("numbers4 = %d", numbers4)
}

// Dictionary test
func Dictionary() {

	// == Definition
	var dict1 map[string]int = map[string]int{"fuck": 0, "shit": 1}
	dict2 := map[string]int{"dick": 2, "testicle": 3}
	dict3 := make(map[string]int)

	fmt.Println(dict1, dict2, dict3)

	// == Exist
	v, ok := dict1["fuck"]
	fmt.Printf("Exist = %t, Value = %d\n", ok, v)

	// == Delete
	delete(dict2, "dick")
	delete(dict2, "ass") // nothing happens

	// == Dict Traverse
	concat := ""
	for k, v := range dict1 {
		concat += k
		concat += strconv.Itoa(v)
	}

	fmt.Println(concat)
}

// ReverseArray test
func ReverseArray(arr []int) {

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}

}

// PanicTest
func PanicTest() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	panic("crash")

}
