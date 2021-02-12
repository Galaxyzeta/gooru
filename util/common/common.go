package common

import (
	"reflect"
	"strings"
)

// Predicate is a short hand for judging whether a equals b.
func Predicate(a interface{}, b interface{}) bool {
	if a == b {
		return true
	}
	return false
}

// MaxInt returns the bigger one between a and b.
func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Stoi convert any slice into []interface{}
func Stoi(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

// HashCode calculates an item's hashcode.
func HashCode(item interface{}) uint32 {
	cvt, ok := item.(string)
	if ok {
		return StringHash(cvt)
	}
	panic("cannot hash!")
}

// StringHash calcs hash for a string
func StringHash(str string) uint32 {
	var ret uint32
	for _, v := range str {
		// ret = ret*133 + (uint32)(v)
		ret += uint32(v)
	}
	return ret
}

// Rehash XORs an uint's lower bit and higher bit
func Rehash(code uint32) uint32 {
	return (code ^ (code >> 16))
}

// IsNumeric returns whether a kind is numeric or not.
func IsNumeric(kind reflect.Kind) bool {
	if kind >= reflect.Int && kind <= reflect.Uint64 || kind >= reflect.Float32 && kind <= reflect.Float64 {
		return true
	}
	return false
}

// BuildString builds a string via StringBuilder. More efficient than using + .
func BuildString(strs ...string) string {
	sb := strings.Builder{}
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}
