# functional-go

Small, boring functional helpers for modern Go.

`functional-go` gives you the helpers Go developers reach for most often:
`Map`, `Filter`, `Reduce`, `Find`, `Any`, `All`, `GroupBy`, `Unique`, and lazy
`iter.Seq` pipelines.

It is intentionally not a lodash clone. It is a tiny package with no
dependencies, no reflection, no code generation, and no chaining DSL.

## Why use this

Use this when a small helper makes the code clearer than a hand-written loop.

Do not use this when a loop is simpler. Go loops are good.

## Compared with lo

[`samber/lo`](https://github.com/samber/lo) is a large lodash-style toolkit.
Use it when you want a broad utility package.

Use `functional-go` when you want a smaller dependency with a stricter shape:

- Core helpers only.
- Standard-library-style names.
- Lazy `iter.Seq` helpers.
- No reflection.
- No transitive dependencies.
- Benchmarks against plain loops.

## Install

```sh
go get github.com/owainlewis/functional-go
```

## Slice helpers

```go
package main

import (
	"fmt"

	functional "github.com/owainlewis/functional-go"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}

	out := functional.Map(functional.Filter(nums, func(n int) bool {
		return n > 2
	}), func(n int) int {
		return n * 10
	})

	fmt.Println(out)
}
```

Output:

```text
[30 40 50]
```

## Iterator helpers

Go 1.23 added standard iterator support with `iter.Seq`.

Use the `Seq` helpers when you want a lazy pipeline.

```go
nums := slices.Values([]int{1, 2, 3, 4, 5})

out := functional.Collect(functional.Take(functional.MapSeq(functional.FilterSeq(nums, func(n int) bool {
	return n > 2
}), func(n int) int {
	return n * 10
}), 2))

fmt.Println(out)
// [30 40]
```

## API

Slice helpers:

- `Map`
- `Filter`
- `Reject`
- `Reduce`
- `FlatMap`
- `Partition`
- `Find`
- `Any`
- `All`
- `Contains`
- `GroupBy`
- `IndexBy`
- `Unique`
- `UniqueBy`

Map helpers:

- `Keys`
- `Values`
- `Entries`

Iterator helpers:

- `MapSeq`
- `FilterSeq`
- `RejectSeq`
- `ReduceSeq`
- `FlatMapSeq`
- `FindSeq`
- `AnySeq`
- `AllSeq`
- `Collect`
- `Take`

## Design

Opinion [high]: This library should stay small.
Flip fact: If the Go standard library adds `Map` and `Filter`, this package
should probably become unnecessary.

Rules:

- Keep names plain.
- Prefer standard library shapes.
- Keep helpers generic and allocation-conscious.
- Never use reflection.
- Never add dependencies.
- Add benchmarks for helpers that allocate or process slices.

## Performance

The helpers are thin wrappers around loops. Benchmarks compare them to plain
loops so users can make an honest call.

Run:

```sh
make bench
```

Expected tradeoff:

- Slice helpers should be close to equivalent hand-written loops.
- Iterator helpers trade some overhead for lazy pipelines and early stop.
- If a loop is clearer or faster in hot code, use the loop.

Current local benchmark shape on an M1 Max:

- `Map`: same allocation profile as a plain loop.
- `Filter`: same allocation profile as a plain loop.
- `Reduce`: zero allocations, same as a plain loop.
- `Seq` pipeline: more overhead than a fused loop, but lazy and composable.

## Development

```sh
make fmt
make test
make bench
```

## License

MIT.
