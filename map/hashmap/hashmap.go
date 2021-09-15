package hashmap

import "fmt"

type Slice struct {
	Key, Value interface{}
}

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

// Cover inserts element into the map With Cover.
func (hm *HashMap) Cover(key interface{}, value interface{}) {
	hm.hm[key] = value
}

func (hm *HashMap) Get(key interface{}) (value interface{}, isfound bool) {
	value, isfound = hm.hm[key]
	return
}

func (hm *HashMap) Remove(key interface{}) {
	delete(hm.hm, key)
}

func (hm *HashMap) Empty() bool {
	return len(hm.hm) == 0
}

func (hm *HashMap) Size() int {
	return len(hm.hm)
}

func (hm *HashMap) Keys() []interface{} {
	keys := make([]interface{}, len(hm.hm))
	count := 0
	for key := range hm.hm {
		keys[count] = key
		count++
	}
	return keys
}

func (hm *HashMap) Values() []interface{} {
	values := make([]interface{}, len(hm.hm))
	count := 0
	for _, value := range hm.hm {
		values[count] = value
		count++
	}
	return values
}

func (hm *HashMap) Slices() []Slice {
	var slices []Slice = make([]Slice, len(hm.hm))

	var i = 0
	for key, value := range hm.hm {
		s := slices[i]
		s.Key = key
		s.Value = value
	}
	return slices
}

func (hm *HashMap) Clear() {
	hm.hm = make(map[interface{}]interface{})
}

func (hm *HashMap) String() string {
	content := fmt.Sprintf("%v", hm.hm)
	return content
}
