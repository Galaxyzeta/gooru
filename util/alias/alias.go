package alias

// Consumer takes parameters and returns nothing.
type Consumer func(params ...interface{})

// P0Consumer takes no input and returns nothing
type P0Consumer func()

// P1Consumer takes one parameters and returns nothing.
type P1Consumer func(param interface{})
