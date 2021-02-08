package alias

// Consumer takes parameters and returns nothing.
type Consumer func(params ...interface{})

// P0Consumer takes no input and returns nothing.
type P0Consumer func()

// P1Consumer takes one parameters and returns nothing.
type P1Consumer func(param interface{})

// P2Consumer takes 2 parameters and returns nothing.
type P2Consumer func(param1 interface{}, param2 interface{})

// Any aliase for any types.
type Any interface{}

// EQ3 is a alias for for a == b? c : d.
func EQ3(a interface{}, b interface{}, eqret interface{}, neqret interface{}) interface{} {
	if a == b {
		return eqret
	}
	return neqret
}
