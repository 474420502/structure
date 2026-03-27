# astar

```go
import "github.com/474420502/structure/graph/astar"
```

`graph/astar` implements grid-based A* path finding with pluggable neighbor, cost, and weight strategies.

## Features

- build an empty grid with `New(dimX, dimY)`
- build from a text map with `NewWithTiles`
- set start/end positions with `SetTarget`
- update terrain with `SetAttr` or `SetStringTiles`
- choose 4-direction or 8-direction expansion through `Neighbor4` and `Neighbor8`
- customize path cost and heuristic weight through `CountCost` and `CountWeight`
- single-path and multi-path search through `Search` and `SearchMulti`
- render grid or path output as strings

## Core Types

- `Graph`
- `Tile`
- `Point`
- `Path` and `PathList`
- `Neighbor`, `Neighbor4`, `Neighbor8`
- `CountCost`, `CountWeight`, `SimpleCost`, `SimpleWeight`

## Useful Methods

- `SetNeighbor`, `SetCountCost`, `SetCountWeight`
- `GetTarget`, `GetPath`, `GetMultiPath`, `GetSteps`, `GetSingleSteps`
- `GetTiles`, `GetTilesWithTarget`, `GetPathTiles`, `GetSinglePathTiles`
- `Clear`, `Search`, `SearchMulti`

## Notes

- Tile attributes use byte constants such as `PLAIN`, `BLOCK`, `START`, `END`, and `PATH`.
- The implementation uses the repository's heap package for the open set.
- Tests include examples for custom cost logic and multi-path search.

## Validation

Behavior is covered by [astar_test.go](./astar_test.go) and [astar_extend_test.go](./astar_extend_test.go).