# API Migration Strategy

This document describes how legacy collection APIs in this repository should evolve without forcing an immediate breaking release.

It builds on:

- [api_standard.md](./api_standard.md)
- [api_consistency_review.md](./api_consistency_review.md)

## Current State

The repository now has standardized, non-breaking compatibility entry points for the main API families:

- writes: `InsertIfAbsent`, `Upsert`, `Delete`, `Len`
- ordered iterators: `SeekGEExact`, `SeekGTExact`, `SeekLEExact`, `SeekLTExact`

Legacy methods remain available and unchanged.

## Migration Goal

The goal is not to remove all legacy names immediately. The goal is to converge user-facing code on one semantic contract while keeping historical behavior available during a transition period.

## Compatibility Policy

### Phase 1: Standard Entry Points First

Status: in progress

Rules:

1. Keep legacy APIs unchanged.
2. Add standardized semantic aliases.
3. Prefer standardized names in docs, examples, and new code.

This phase is already implemented for the main map and ordered-tree packages.

### Phase 2: Documentation And Example Convergence

Status: in progress

Rules:

1. All new examples should prefer:
   - `InsertIfAbsent`
   - `Upsert`
   - `Delete`
   - `Len`
   - `Seek*Exact` when exact-match status matters
2. Package docs should explain historical names only as compatibility surfaces.
3. README-level summaries should describe the standardized entry points instead of reinforcing legacy semantics.

### Phase 3: Soft Deprecation

Status: planned

Rules:

1. Mark typo aliases and ambiguous names as compatibility-only in docs.
2. Prefer comments such as:
   - `Deprecated: use Valid.`
   - `Compatibility alias. Prefer InsertIfAbsent.`
3. Do not remove symbols in this phase.

Primary candidates:

- `Vaild`
- package-specific uses of `Add` or `Set` where the semantic meaning is ambiguous outside that package
- iterator `Seek*` variants that hide exact-match status when a `Seek*Exact` alternative now exists

### Phase 4: Breaking Release Cleanup

Status: planned

Rules:

1. Only remove or rename legacy APIs in a versioned breaking release.
2. Group breaking changes by contract family rather than by package.
3. Publish migration notes with mechanical replacements.

Recommended breaking groups:

#### Group A: Size/Len cleanup

- standardize on `Len() int`
- keep `Size()` only if there is a strong type-specific reason

#### Group B: Remove/Delete cleanup

- standardize on `Delete() (value, ok)`
- retain `Remove` only where the package is intentionally non-map-like

#### Group C: Seek cleanup

- standardize on `Seek*Exact` for exact-match reporting
- consider changing legacy `Seek*` to return `bool` only in a major release, if desired

#### Group D: Naming cleanup

- remove `Vaild`
- consider aligning `queue/priority` package naming with its directory

## Recommended Replacements

| Legacy usage | Preferred replacement | Notes |
| --- | --- | --- |
| `Put` when caller expects insert-only semantics | `InsertIfAbsent` | Explicit and consistent |
| `Set` when caller expects overwrite-or-create semantics | `Upsert` | Explicit and consistent |
| `Remove` when caller needs removed value + status | `Delete` | Avoids per-package return ambiguity |
| `Size` in cross-package code | `Len` | Avoids `int`/`uint`/`int64` drift |
| `SeekGE` plus manual exact-check logic | `SeekGEExact` | Single-call exact status |
| `SeekGT` plus manual exact-check logic | `SeekGTExact` | Boundary behavior remains explicit |
| `SeekLE` plus manual exact-check logic | `SeekLEExact` | Single-call exact status |
| `SeekLT` plus manual exact-check logic | `SeekLTExact` | Boundary behavior remains explicit |
| `Vaild` | `Valid` | Typo compatibility only |

## Risk Guidance

### Safe to prefer immediately

- `InsertIfAbsent`
- `Upsert`
- `Delete`
- `Len`
- `Valid`
- `Seek*Exact`

### Keep but avoid in new code

- `Put`
- `Set`
- `Remove`
- `Size`
- `Vaild`
- `Seek*` when exact-match status matters

### Requires major-version decision

- removing legacy names
- changing return types of legacy methods
- renaming the `treequeue` package

## Bottom Line

The repository no longer needs to choose between compatibility and coherence.

The practical path is:

1. keep legacy APIs stable
2. move users to standardized entry points
3. defer true removals to a deliberate breaking release