package assert

import "fmt"

// EQ panics if actual != expect.
func EQ(actual interface{}, expect interface{}) {
	if actual == expect {
		return
	}
	panic(fmt.Sprintf("expect %v but actual is %v", expect, actual))
}

// NEQ panics if actual == expect.
func NEQ(actual interface{}, expect interface{}) {
	if actual != expect {
		return
	}
	panic(fmt.Sprintf("expect %v but actual is %v", expect, actual))
}

// ThrowsPanic asserts a panic to be thrown.
func ThrowsPanic(fn func()) {
	defer func() {
		p := recover()
		if p == nil {
			panic("expect panic to be thrown, but none.")
		}
	}()
	fn()
}
