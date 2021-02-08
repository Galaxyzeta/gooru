package hashmap

// Map is an interface defining essential elems for a map.
type Map interface {
	Get(k interface{}) interface{}
	Put(k interface{}, v interface{})
	Size() int
	Delete(k interface{})
}
