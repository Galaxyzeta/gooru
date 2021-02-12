package consistenthash_test

import (
	"fmt"
	"testing"

	"galaxyzeta.com/consistenthash"
)

func TestConsistentHash(t *testing.T) {
	c := consistenthash.NewConsistentHash()
	c.Add("6")
	c.Add("4")
	c.Add("2")

	fmt.Println(c.Get("2"))
	fmt.Println(c.Get("23"))
	fmt.Println(c.Get("123"))
}
