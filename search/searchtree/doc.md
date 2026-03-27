# searchtree

```go
import "github.com/474420502/structure/search/searchtree"
```

`search/searchtree` is currently a placeholder package for future search-index abstractions.

## Current State

- `Tree[VALUE any]` currently contains only an internal `values []VALUE` field.
- `Indexes[VALUE any]` is declared but empty.
- `Index[KEY, VALUE any]` is declared as an empty interface.
- There are no exported constructors or operational methods yet.

## Recommendation

Do not depend on this package as a public API yet. For a usable ordered search structure today, prefer [../treelist/doc.md](../treelist/doc.md).