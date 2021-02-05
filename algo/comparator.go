package algo

import (
	"fmt"
	"reflect"
)

// ISortable implemented objects can be sorted.
type ISortable interface {
	Cmp(a int, b int) bool
	Len() int
	Swap(a int, b int)
}

// IntComparator is used for any []int sort.
type IntComparator struct {
	Data []int
	Asc  bool
}

func (cmp *IntComparator) Less(a int, b int) bool {
	if cmp.Asc {
		return cmp.Data[a] < cmp.Data[b]
	}
	return cmp.Data[a] > cmp.Data[b]
}
func (cmp *IntComparator) Len() int { ; return len(cmp.Data) }
func (cmp *IntComparator) Swap(a int, b int) {
	cmp.Data[a], cmp.Data[b] = cmp.Data[b], cmp.Data[a]
}

// SimpleTypeComparator is used for any comparable simple type based sort.
type SimpleTypeComparator struct {
	Data []interface{}
	Asc  bool
}

func (cmp *SimpleTypeComparator) Cmp(a int, b int) bool {
	a1 := cmp.Data[a]
	b1 := cmp.Data[b]

	kind := reflect.ValueOf(a1).Type().Kind()
	// === TEMPLATE ===
	// case reflect.Int: if (cmp.Asc) {; return a1.(int) < b1.(int) ;}; return a1.(int) > b1.(int)
	// case reflect.Int8: if (cmp.Asc) {; return a1.(int8) < b1.(int8) ;}; return a1.(int8) > b1.(int8)
	// case reflect.Int16: if (cmp.Asc) {; return a1.(int16) < b1.(int16) ;}; return a1.(int16) > b1.(int16)
	// case reflect.Int32: if (cmp.Asc) {; return a1.(int32) < b1.(int32) ;}; return a1.(int32) > b1.(int32)
	// case reflect.Int64: if (cmp.Asc) {; return a1.(int64) < b1.(int64) ;}; return a1.(int64) > b1.(int64)
	// case reflect.Uint: if (cmp.Asc) {; return a1.(uint) < b1.(uint) ;}; return a1.(uint) > b1.(uint)
	// case reflect.Uint8: if (cmp.Asc) {; return a1.(uint8) < b1.(uint8) ;}; return a1.(uint8) > b1.(uint8)
	// case reflect.Uint16: if (cmp.Asc) {; return a1.(uint16) < b1.(uint16) ;}; return a1.(uint16) > b1.(uint16)
	// case reflect.Uint32: if (cmp.Asc) {; return a1.(uint32) < b1.(uint32) ;}; return a1.(uint32) > b1.(uint32)
	// case reflect.Uint64: if (cmp.Asc) {; return a1.(uint64) < b1.(uint64) ;}; return a1.(uint64) > b1.(uint64)
	// case reflect.String: if (cmp.Asc) {; return a1.(string) < b1.(string) ;}; return a1.(string) > b1.(string)
	// case reflect.Float32: if (cmp.Asc) {; return a1.(float32) < b1.(float32) ;}; return a1.(float32) > b1.(float32)
	// case reflect.Float64: if (cmp.Asc) {; return a1.(float64) < b1.(float64) ;}; return a1.(float64) > b1.(float64)
	switch kind {
	case reflect.Int:
		if cmp.Asc {
			return a1.(int) < b1.(int)
		}
		return a1.(int) > b1.(int)
	case reflect.Int8:
		if cmp.Asc {
			return a1.(int8) < b1.(int8)
		}
		return a1.(int8) > b1.(int8)
	case reflect.Int16:
		if cmp.Asc {
			return a1.(int16) < b1.(int16)
		}
		return a1.(int16) > b1.(int16)
	case reflect.Int32:
		if cmp.Asc {
			return a1.(int32) < b1.(int32)
		}
		return a1.(int32) > b1.(int32)
	case reflect.Int64:
		if cmp.Asc {
			return a1.(int64) < b1.(int64)
		}
		return a1.(int64) > b1.(int64)
	case reflect.Uint:
		if cmp.Asc {
			return a1.(uint) < b1.(uint)
		}
		return a1.(uint) > b1.(uint)
	case reflect.Uint8:
		if cmp.Asc {
			return a1.(uint8) < b1.(uint8)
		}
		return a1.(uint8) > b1.(uint8)
	case reflect.Uint16:
		if cmp.Asc {
			return a1.(uint16) < b1.(uint16)
		}
		return a1.(uint16) > b1.(uint16)
	case reflect.Uint32:
		if cmp.Asc {
			return a1.(uint32) < b1.(uint32)
		}
		return a1.(uint32) > b1.(uint32)
	case reflect.Uint64:
		if cmp.Asc {
			return a1.(uint64) < b1.(uint64)
		}
		return a1.(uint64) > b1.(uint64)
	case reflect.String:
		if cmp.Asc {
			return a1.(string) < b1.(string)
		}
		return a1.(string) > b1.(string)
	case reflect.Float32:
		if cmp.Asc {
			return a1.(float32) < b1.(float32)
		}
		return a1.(float32) > b1.(float32)
	case reflect.Float64:
		if cmp.Asc {
			return a1.(float64) < b1.(float64)
		}
		return a1.(float64) > b1.(float64)

	default:
		panic(fmt.Sprintf("cannot compare %s !\n", kind))
	}
}
func (cmp *SimpleTypeComparator) Len() int { ; return len(cmp.Data) }
func (cmp *SimpleTypeComparator) Swap(a int, b int) {
	cmp.Data[a], cmp.Data[b] = cmp.Data[b], cmp.Data[a]
}
