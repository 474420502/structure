# orderedmap

```go
import "github.com/474420502/structure/map/orderedmap.go"
```

`map/orderedmap.go` is an ordered map implementation based on `tree/indextree`.

## Features

- Ordered key-value storage with O(log n) get, put, remove operations
- Index-based access: `Index(i)` to get key/value at position, `IndexOf(key)` to get position of key
- Iterator support for forward/backward traversal and seeking
- Standardized semantic aliases: `InsertIfAbsent`, `Upsert`, `Delete`, `Len`

## API Snapshot

- `New[K, V](comp compare.Compare[K]) *OrderedMap[K, V]`
- `Put(key K, value V) bool`
- `InsertIfAbsent(key K, value V) bool`
- `Set(key K, value V) bool`
- `Upsert(key K, value V) bool`
- `Get(key K) (V, bool)`
- `Remove(key K) (V, bool)`
- `Delete(key K) (V, bool)`
- `Contains(key K) bool`
- `IndexOf(key K) int64`
- `Index(index int64) (K, V)`
- `RemoveIndex(index int64) (K, V, bool)`
- `Keys() []K`
- `Values() []V`
- `Clear()`
- `Size() int64`
- `Len() int`
- `Iterator() *Iterator[K, V]`

## Usage

```go
om := orderedmap.New[int, string](compare.Any[int])
om.InsertIfAbsent(1, "one")
om.InsertIfAbsent(2, "two")
om.Upsert(1, "ONE") // updates value

v, ok := om.Get(1) // v == "ONE", ok == true
idx := om.IndexOf(2) // idx == 1

om.Delete(1)

keys := om.Keys()   // [2]
values := om.Values() // ["two"]

// Iterator
iter := om.Iterator()
for iter.SeekToFirst(); iter.Valid(); iter.Next() {
    fmt.Println(iter.Key(), iter.Value())
}
```

## Notes

- `Put` preserves the existing value when the key already exists and returns `false` in that case.
- `Set` retains the historical `indextree` semantics and is kept for compatibility.
- `InsertIfAbsent`, `Upsert`, `Delete`, and `Len` are the preferred cross-package entry points for new code.
- Iterator `Seek*` methods already return the exact-match status directly through the wrapped `indextree` iterator.

## Performance

- Get: O(log n)
- Put: O(log n)
- Set: O(log n)
- Remove: O(log n)
- Index/IndexOf: O(log n)