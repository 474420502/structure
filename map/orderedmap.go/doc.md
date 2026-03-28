# orderedmap

```go
import "github.com/474420502/structure/map/orderedmap.go"
```

`map/orderedmap.go` is an ordered map implementation based on `tree/indextree`.

## Features

- Ordered key-value storage with O(log n) get, put, remove operations
- Index-based access: `Index(i)` to get key/value at position, `IndexOf(key)` to get position of key
- Iterator support for forward/backward traversal and seeking
- Range operations: `RemoveIndex`, `Trim`, `Split`

## Usage

```go
om := orderedmap.New(compare.Any[int]())
om.Put(1, "one")
om.Put(2, "two")
om.Set(1, "ONE") // updates value

v, ok := om.Get(1) // v == "ONE", ok == true
idx := om.IndexOf(2) // idx == 1

om.Remove(1)

keys := om.Keys()   // [2]
values := om.Values() // ["two"]

// Iterator
iter := om.Iterator()
for iter.SeekToFirst(); iter.Valid(); iter.Next() {
    fmt.Println(iter.Key(), iter.Value())
}
```

## Performance

- Get: O(log n)
- Put: O(log n)
- Set: O(log n)
- Remove: O(log n)
- Index/IndexOf: O(log n)