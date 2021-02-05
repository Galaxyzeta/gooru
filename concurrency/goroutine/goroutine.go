package goroutine

import (
	"runtime"
	"strconv"
	"strings"
	"time"
)

// GoID get current goroutine's identifier from runtime stack.
//
// @Warning: might be slow !
func GoID() int {
	buf := getStack()
	str := strings.Split(string(buf), " ")[1]
	ret, err := strconv.Atoi(str)
	if err != nil {
		panic("cannot get goid")
	}
	return ret
}

// Sleep is a wrapper for time.Sleep()
func Sleep(duration int, unit time.Duration) {
	time.Sleep(time.Duration(duration) * unit)
}

func getStack() []byte {
	buf := make([]byte, 64)
	runtime.Stack(buf, false)
	return buf
}
