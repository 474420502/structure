package hashmap

import (
	"fmt"
)

// Slice the KeyValue
type Slice struct {
	Key, Value interface{}
}

// HashMap map base on hash
type HashMap struct {
	hm map[interface{}]interface{}
}

// New instantiates a hash map.
func New() *HashMap {
	return &HashMap{hm: make(map[interface{}]interface{})}
}

// New instantiates a hash map with  Capacity.
func NewWithCap(cap int) *HashMap {
	return &HashMap{hm: make(map[interface{}]interface{}, cap)}
}

// Put inserts element into the map With Not Cover. if key exists return false. else return true
func (hm *HashMap) Put(key interface{}, value interface{}) bool {
	if _, ok := hm.hm[key]; !ok {
		hm.hm[key] = value
		return true
	}
	return false
}

// Set inserts element into the map With Set.
func (hm *HashMap) Set(key interface{}, value interface{}) {
	hm.hm[key] = value
}

// Get get the element by key
func (hm *HashMap) Get(key interface{}) (value interface{}, isfound bool) {
	value, isfound = hm.hm[key]
	return
}

// Remove remove the element by key
func (hm *HashMap) Remove(key interface{}) {
	delete(hm.hm, key)
}

// Empty if the hashmap is empty, return true
func (hm *HashMap) Empty() bool {
	return len(hm.hm) == 0
}

// Size return the size of hashmap
func (hm *HashMap) Size() int {
	return len(hm.hm)
}

// Keys return the all keys of hashmap. non order
func (hm *HashMap) Keys() []interface{} {
	keys := make([]interface{}, len(hm.hm))
	count := 0
	for key := range hm.hm {
		keys[count] = key
		count++
	}
	return keys
}

// Values return the all values of hashmap. non order
func (hm *HashMap) Values() []interface{} {
	values := make([]interface{}, len(hm.hm))
	count := 0
	for _, value := range hm.hm {
		values[count] = value
		count++
	}
	return values
}

// Slices return the all keyvalue of hashmap. non order
func (hm *HashMap) Slices() []Slice {
	var slices []Slice = make([]Slice, len(hm.hm))

	var i = 0
	for key, value := range hm.hm {
		s := &slices[i]
		s.Key = key
		s.Value = value
		i++
	}
	return slices
}

// Clear clear the hashmap
func (hm *HashMap) Clear() {
	hm.hm = make(map[interface{}]interface{})
}

// String print the hashmap
func (hm *HashMap) String() string {
	content := fmt.Sprintf("%v", hm.hm)
	return content
}
