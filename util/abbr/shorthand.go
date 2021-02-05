package abbr

import "reflect"

// Predicate is a short hand for judging whether a equals b.
func Predicate(a interface{}, b interface{}) bool {
	if a == b {
		return true
	}
	return false
}

// CondExpEq is a short hand for a == b? c : d.
func CondExpEq(a interface{}, b interface{}, eqret interface{}, neqret interface{}) interface{} {
	if a == b {
		return eqret
	}
	return neqret
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
