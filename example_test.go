package functional_test

import (
	"fmt"
	"slices"

	functional "github.com/owainlewis/functional-go"
)

func ExampleMap() {
	nums := []int{1, 2, 3}

	doubled := functional.Map(nums, func(n int) int {
		return n * 2
	})

	fmt.Println(doubled)
	// Output: [2 4 6]
}

func ExampleFilter() {
	nums := []int{1, 2, 3, 4, 5}

	evens := functional.Filter(nums, func(n int) bool {
		return n%2 == 0
	})

	fmt.Println(evens)
	// Output: [2 4]
}

func ExampleMapSeq() {
	nums := slices.Values([]int{1, 2, 3, 4, 5})

	out := functional.Collect(functional.Take(functional.MapSeq(functional.FilterSeq(nums, func(n int) bool {
		return n > 2
	}), func(n int) int {
		return n * 10
	}), 2))

	fmt.Println(out)
	// Output: [30 40]
}
