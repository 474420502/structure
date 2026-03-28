# API Consistency Review

This document summarizes public API inconsistencies across the main containers in this repository.

## Scope

The review focuses on these API families because they shape cross-package expectations:

- `Set`
- `Put` / `Add`
- `Remove`
- `Size`
- iterator `Seek*`

Representative packages reviewed:

- `tree/indextree`
- `tree/avl`
- `tree/avls`
- `tree/treelist`
- `set/treeset`
- `queue/priority`
- `map/hashmap`
- `map/linkedhashmap`
- `map/orderedmap.go`

## Executive Summary

The current API surface is usable within each package, but it is not coherent across packages.

The main design issue is not naming alone. The deeper problem is that the same method names often carry different semantic contracts:

- `Set` sometimes means "upsert", sometimes "update existing only".
- `Put` sometimes means "insert if absent", sometimes "always insert", sometimes "insert duplicate keys allowed".
- `bool` return values do not have one shared meaning.
- `Remove` result types vary between `void`, `value`, and `(value, ok)`.
- `Size` switches between `int`, `uint`, and `int64`.
- iterator `Seek*` methods sometimes return exact-match status and sometimes return nothing.

That means package-to-package substitution is difficult, wrappers have to adapt semantics manually, and users have to re-learn each package instead of relying on one repository-wide convention.

## Matrix

| Package | `Set` | `Put` / `Add` | `Remove` | `Size` | `Seek*` |
| --- | --- | --- | --- | --- | --- |
| `tree/indextree` | upsert, returns `bool` with mixed meaning | insert-if-absent, returns inserted status | returns removed value or `nil` | `int64` | returns `bool` exact-match status |
| `tree/avl` | upsert, returns inserted status | insert-if-absent, returns inserted status | `(value, ok)` | `uint` | returns `bool` exact-match status |
| `tree/avls` | upsert, returns inserted status | always inserts, duplicates allowed | `(value, ok)` | `uint` | returns `bool` exact-match status |
| `tree/treelist` | not part of this review focus | not part of this review focus | not part of this review focus | not part of this review focus | returns `bool` exact-match status |
| `set/treeset` | upsert, returns inserted status | uses `Add` instead of `Put` | `(value, ok)` | `uint` | returns nothing |
| `queue/priority` | no `Set` | insert-if-absent among same-key-first lookup, returns inserted status | `(value, ok)` | `int` | returns nothing |
| `map/hashmap` | unconditional overwrite, no return | insert-if-absent, returns inserted status | no return | `int` | not applicable |
| `map/linkedhashmap` | update-existing-only, returns updated status | insert-if-absent, returns inserted status | `(value, ok)` | `uint` | not applicable |
| `map/orderedmap.go` | mirrors `indextree` | mirrors `indextree` | `(value, ok)` wrapper over `indextree` | `int64` | returns `bool` exact-match status |

## Verified Evidence

### 1. `Set` does not have one repository-wide meaning

Cases:

- `tree/avl` and `set/treeset`: upsert, return whether a new key was inserted.
- `tree/indextree`: upsert, but its `bool` means "existing key was replaced" for existing keys and "new key inserted" only when the tree was empty. For non-empty trees, insertion returns `false`.
- `map/hashmap`: unconditional overwrite and no return value.
- `map/linkedhashmap`: update-existing-only, returns `true` only when the key already exists.

Impact:

- The same code pattern cannot be reused safely across containers.
- `orderedmap` inherits `indextree` semantics, so the inconsistency propagates upward.

Evidence:

- `tree/indextree/tree.go`
- `tree/avl/tree.go`
- `tree/avls/tree.go`
- `set/treeset/tree.go`
- `map/hashmap/hashmap.go`
- `map/linkedhashmap/linked_hashmap.go`
- `map/orderedmap.go/map.go`

### 2. `Put` is also inconsistent, and `Add` introduces a third alias

Cases:

- `tree/indextree`, `tree/avl`, `map/hashmap`, `map/linkedhashmap`: `Put` is insert-if-absent.
- `tree/avls`: `Put` always inserts and allows duplicate keys.
- `set/treeset`: equivalent operation is called `Add`, not `Put`.
- `queue/priority`: `Put` inserts, but the structure is semantically closer to a multimap/priority tree than to a map.

Impact:

- A caller cannot infer duplicate-key behavior from the method name alone.
- Generic examples and wrappers are forced to special-case package behavior.

### 3. `Remove` result contracts differ too much

Cases:

- `tree/indextree`: returns removed value or `nil`.
- `map/hashmap`: no result.
- `tree/avl`, `tree/avls`, `set/treeset`, `queue/priority`, `map/linkedhashmap`, `map/orderedmap.go`: `(value, ok)`.

Impact:

- Error handling and missing-key handling differ even when the logical operation is identical.
- Wrapper packages need conversion code solely because base packages disagree.

### 4. `Size` uses three integer types without a clear rule

Observed:

- `int64`: `tree/indextree`, `map/orderedmap.go`
- `uint`: `tree/avl`, `tree/avls`, `set/treeset`, `map/linkedhashmap`
- `int`: `queue/priority`, `map/hashmap`

Impact:

- Callers pay conversion cost for cross-package composition.
- Negative-index support in some structures pushes APIs toward signed integers anyway.
- The mixed use of `uint` is especially awkward in Go public APIs because indices, lengths, and loops are usually driven by `int`.

### 5. iterator `Seek*` contracts are split into two styles

Style A:

- `tree/indextree`
- `tree/avl`
- `tree/avls`
- `tree/treelist`
- `map/orderedmap.go`

These return `bool`, where `true` means an exact match was found, while the iterator may still land on the nearest valid neighbor when the return value is `false`.

Style B:

- `set/treeset`
- `queue/priority`

These do not return anything. Callers must inspect iterator validity and key/value after the call.

Impact:

- Users cannot transfer iterator usage patterns across ordered structures.
- It weakens discoverability because the exact-match status is sometimes explicit and sometimes hidden.

### 6. package naming and directory naming diverge in `queue/priority`

Observed:

- Files under `queue/priority` use package name `treequeue`.

Impact:

- It breaks the intuitive directory-to-package expectation present in the rest of the repository.
- This increases import aliasing overhead in examples and user code.

### 7. Some APIs expose historical behavior rather than a normalized contract

Examples:

- `tree/indextree.RemoveRange` uses `panic` for invalid argument order instead of returning an error.
- `map/linkedhashmap.Set` has a narrower contract than nearly every other `Set` in the repository.
- `set/treeset` still carries the historical `Vaild` spelling alongside `Valid`.

These do not all require immediate breaking changes, but they do show that the public surface evolved per-package rather than from a shared contract.

## Risk Ranking

### High priority, low breakage

1. Keep normalizing naming aliases where possible.
2. Document exact semantics of `Set`, `Put`, `Add`, `Remove`, and `Seek*` in package docs.
3. Standardize iterator docs around `Valid()` as the primary name.

### Medium priority, medium breakage

1. Introduce repository-wide semantic guidance:
   - `Put`: insert if absent
   - `Set`: upsert
   - `Add`: only for true multiset / duplicate-key containers
2. Add compatibility wrappers for packages that do not match the guidance.
3. Add explicit helper methods when semantics must stay different, for example `Update`, `Upsert`, `InsertIfAbsent`.

### High priority, breaking changes

1. Unify `Remove` toward `(value, ok)`.
2. Unify `Size` toward one signed type, preferably `int` for idiomatic Go APIs or `int64` only where indexing semantics require it consistently.
3. Unify ordered iterator `Seek*` to always return exact-match status.
4. Rename `queue/priority` package from `treequeue` to a name aligned with the directory.

## Recommended Normalization Order

1. Documentation pass: explicitly state per-package semantics.
2. Compatibility pass: add alias methods or wrapper helpers without breaking callers.
3. New API policy: freeze shared meaning for `Set`, `Put`, `Remove`, `Size`, and `Seek*`.
4. Breaking-change release: collapse legacy variants after a transition window.

## Bottom Line

The repository does not mainly suffer from isolated naming mistakes. It has a shared-contract problem.

If this codebase is expected to feel like one container library rather than many unrelated experiments, the next meaningful step is not another small rename. It is to define one semantic standard for core methods and then migrate packages toward it in stages.