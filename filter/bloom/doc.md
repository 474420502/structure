# bloom

```go
import "github.com/474420502/structure/filter/bloom"
```

`filter/bloom` implements a compact Bloom filter with binary serialization helpers.

## Features

- construct from bit capacity with `New`
- reconstruct from binary data with `NewByDecode`
- add/check values via `Add` and `Contains`
- byte-oriented variants `AddBytes` and `ContainsBytes`
- occupancy inspection with `Cap`, `HitSize`, and `HitRatio`
- `Reset`, `Encode`, and `Decode`

## API Snapshot

- `New(bitsCap uint64) *Bloom`
- `NewByDecode(reader io.Reader) *Bloom`
- `Add(key interface{}) bool`
- `AddBytes(key []byte) bool`
- `Contains(key interface{}) bool`
- `ContainsBytes(key []byte) bool`
- `Cap() uint64`
- `HitSize() uint64`
- `HitRatio() float64`
- `Reset()`
- `Encode() *bytes.Buffer`
- `Decode(reader io.Reader)`

## Notes

- The implementation uses a single `fnv.New64()` hash stream.
- Like any Bloom filter, `Contains` may return false positives.
- The constructor comment recommends sizing bits to roughly `estimated_keys * 10`.

## Validation

Behavior is covered by [bloom_test.go](./bloom_test.go).