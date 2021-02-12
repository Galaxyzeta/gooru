package hashmap

// Map is an interface defining essential elems for a map.
type Map interface {
	Get(k interface{}) interface{}
	Put(k interface{}, v interface{})
	Size() int
	Delete(k interface{}) interface{}
	ContainsKey(k interface{}) bool
}

// Set is an interface defining a pool of unique items to be stored.
type Set interface {
	Put(v interface{})
	Size() int
	Delete(v interface{})
	Contains(v interface{}) bool
}
