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
	)
}
