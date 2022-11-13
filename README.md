# type-param-common

Type parameter primitives and commons.

This repository contains type-parameter version of containers
and set, iterator, queue, deque, stack implementations.

## Suffixed with `-param` to avoid name overlapping

Wrappers of Go std lib are suffixed with `-param`

## Package descriptions

### Wrappers of `container` of standard

- heap-param
  - STATUS: done. doc comments needs improvement.
- list-param
  - STATUS: done. most of behavior consistent to `container/list`.
- ring-param
  - STATUS: Half done. Needs good testing.

Heap-param wraps `container/heap` of standard library.
see `./heap.go`, `./filterable_heap.go` and corresponding test files for example usages.

list-param and ring-param wraps `container/list` and `container/ring` respectively.

*Element[T] and *Ring[T] have addtitional methods, `Get`, `Set` and `Unwrap` to replace direct mutation/observation of `elementOrRing.Value`. Some of their methods are changed their returned value from single `ret T` to `(ret T, ok bool)`. This second boolean value indicates internal Value was not nil, which means returned `ret T` is zero value if false.

### sync-param

Wrappers of `sync.Map` and `sync.Pool`.

### slice

Deque[T], Queue[T], Stack[T] and helper functions. It eases pain of `write-deque-everywhere`.

- Deque
- Queue
- Stack
- Helper functions
  - Useful for unsorted slice.

### iterator

Exerimental iterator impl for go.

- STATUS: half done. requires addition of missing tests. add benchmark for performance comparision.

`iterator` package's main struct is `Iterator[T]`.

```go
type Iterator[T any] struct {
	SeIterator[T]
}

type Nexter[T any] interface {
	Next() (next T, ok bool)
}

// Singly ended iterator.
type SeIterator[T any] interface {
	Nexter[T]
}
```

Usage example:

```go
package main

import (
	"fmt"
	"strings"

	"github.com/ngicks/type-param-common/iterator"
)

func main() {
	fmt.Println(
		iterator.Fold[iterator.EnumerateEnt[string]](
			iterator.Enumerate[string](
				iterator.
					FromSlice([]string{"foo", "bar", "baz"}).
					Map(func(s string) string { return s + s }).
					Exclude(func(s string) bool { return strings.Contains(s, "az") }).
					MustReverse(),
			),
			func(accumulator map[string]int, next iterator.EnumerateEnt[string]) map[string]int {
				accumulator[next.Next] = next.Count
				return accumulator
			},
			map[string]int{},
		),
	)  // outputs: map[barbar:0 foofoo:1]
}
```

It requires Singly ended iterator as embeded field. Iterator[T] is thin helper type. Chance to have breaking change (e.g. new field) is low.

Iterator[T] can be created directly from

- slice []T
- listparam.List[T]
- channel <-chan T

```go
var iter iterator.Iterator[string]
iter = iterator.FromSlice(strSlice)
// or
// for size-fixed iterator
iter = iterator.FromFixedList(list)
// for growing iterator
iter = iterator.FromList(list)
// or
iter = iterator.FromChannel(make(chan string))
```

Input SeIterator[T] can optionally implements interfaces shown below. Those will be used in iterator methods of corresponding name.

```go
type SizeHinter interface {
	SizeHint() int
}

type Reverser[T any] interface {
	Reverse() (rev SeIterator[T], ok bool)
}

type Unwrapper[T any] interface {
	Unwrap() SeIterator[T]
}


// Doubly ended iterator.
type DeIterator[T any] interface {
	Nexter[T]
	NextBacker[T]
}
```
