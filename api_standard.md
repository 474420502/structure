# API Standard Draft

This document defines the preferred semantic contract for collection APIs in this repository.

It does not remove existing methods. It establishes the preferred names for future code and for non-breaking compatibility wrappers.

## Preferred Core Methods

### `InsertIfAbsent(key, value) bool`

- Inserts only when the key is not present.
- Returns `true` when a new entry was inserted.
- Returns `false` when the key already existed and the collection was left unchanged.

### `Upsert(key, value) bool`

- Ensures the key exists with the provided value.
- Returns `true` when an existing entry was replaced or updated.
- Returns `false` when a new entry was inserted.

### `Delete(key) (value, ok)`

- Removes the key if present.
- Returns the removed value and `true` on success.
- Returns the zero or `nil` value and `false` when the key is absent.

### `Len() int`

- Returns the number of elements as `int`.
- This is the preferred cross-package size accessor for new code.
- Existing `Size()` methods remain supported for compatibility.

### Ordered iterator `Seek*Exact(key) bool`

- Applies to ordered iterators that position to boundary matches such as `SeekGE`, `SeekGT`, `SeekLE`, and `SeekLT`.
- Performs the same positioning as the legacy `Seek*` call.
- Returns `true` when the queried key existed exactly.
- Returns `false` when the iterator was positioned to the nearest valid neighbor or became invalid.

Preferred names:

- `SeekGEExact`
- `SeekGTExact`
- `SeekLEExact`
- `SeekLTExact`

## Legacy Method Guidance

Legacy names stay supported, but their semantics vary by package and should not be used as the repository-wide contract.

- `Put`: historical alias in several packages. Often means insert-if-absent, but not always.
- `Set`: historical alias with especially inconsistent meaning across packages.
- `Remove`: historical alias with inconsistent return signatures.
- `Size`: historical alias with inconsistent integer types.
- `SeekGE` / `SeekGT` / `SeekLE` / `SeekLT`: historical iterator entry points with inconsistent return signatures.

## Migration Rule

For non-breaking convergence:

1. Keep legacy methods unchanged.
2. Add preferred semantic wrappers with the contracts above.
3. Update examples and docs to prefer the standardized names.

## First Convergence Batch

The first batch applies this standard to:

- `map/hashmap`
- `map/linkedhashmap`
- `tree/indextree`
- `map/orderedmap.go`