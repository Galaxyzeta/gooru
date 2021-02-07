package hashmap

// Map is an interface defining essential elems for a map.
type Map interface {
	Get(k interface{})
	Put(k interface{}, v interface{})
	Len()
	Delete(k interface{})
}
