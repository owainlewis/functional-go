package functional

import (
	"iter"
	"reflect"
	"slices"
	"testing"
)

type namedInts []int

func TestMap(t *testing.T) {
	got := Map([]int{1, 2, 3}, func(n int) string {
		return string(rune('a' + n - 1))
	})

	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestFilterPreservesNamedSliceType(t *testing.T) {
	got := Filter(namedInts{1, 2, 3, 4}, func(n int) bool {
		return n%2 == 0
	})

	if _, ok := any(got).(namedInts); !ok {
		t.Fatalf("got %T, want namedInts", got)
	}

	want := namedInts{2, 4}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestReject(t *testing.T) {
	got := Reject([]int{1, 2, 3, 4}, func(n int) bool {
		return n%2 == 0
	})

	want := []int{1, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestReduce(t *testing.T) {
	got := Reduce([]int{1, 2, 3, 4}, 0, func(total, n int) int {
		return total + n
	})

	if got != 10 {
		t.Fatalf("got %d, want 10", got)
	}
}

func TestFlatMap(t *testing.T) {
	got := FlatMap([]int{1, 2, 3}, func(n int) []int {
		return []int{n, n * 10}
	})

	want := []int{1, 10, 2, 20, 3, 30}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestPartition(t *testing.T) {
	evens, odds := Partition([]int{1, 2, 3, 4}, func(n int) bool {
		return n%2 == 0
	})

	if !reflect.DeepEqual(evens, []int{2, 4}) {
		t.Fatalf("evens got %v", evens)
	}
	if !reflect.DeepEqual(odds, []int{1, 3}) {
		t.Fatalf("odds got %v", odds)
	}
}

func TestFindAnyAllContains(t *testing.T) {
	nums := []int{1, 2, 3, 4}

	got, ok := Find(nums, func(n int) bool { return n > 2 })
	if !ok || got != 3 {
		t.Fatalf("Find got %d, %v, want 3, true", got, ok)
	}

	if !Any(nums, func(n int) bool { return n == 4 }) {
		t.Fatal("Any got false, want true")
	}

	if !All(nums, func(n int) bool { return n > 0 }) {
		t.Fatal("All got false, want true")
	}

	if !Contains(nums, 2) {
		t.Fatal("Contains got false, want true")
	}
}

func TestGroupByIndexByUnique(t *testing.T) {
	type user struct {
		ID   int
		Team string
	}

	users := []user{
		{ID: 1, Team: "eng"},
		{ID: 2, Team: "ops"},
		{ID: 3, Team: "eng"},
	}

	grouped := GroupBy(users, func(u user) string {
		return u.Team
	})
	if len(grouped["eng"]) != 2 || len(grouped["ops"]) != 1 {
		t.Fatalf("unexpected groups: %v", grouped)
	}

	indexed := IndexBy(users, func(u user) int {
		return u.ID
	})
	if indexed[2].Team != "ops" {
		t.Fatalf("unexpected index: %v", indexed)
	}

	if got := Unique([]int{1, 1, 2, 3, 3}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Fatalf("Unique got %v", got)
	}

	uniqueUsers := UniqueBy(users, func(u user) string {
		return u.Team
	})
	if !reflect.DeepEqual(uniqueUsers, users[:2]) {
		t.Fatalf("UniqueBy got %v, want %v", uniqueUsers, users[:2])
	}
}

func TestMapHelpers(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2}

	keys := Keys(input)
	slices.Sort(keys)
	if !reflect.DeepEqual(keys, []string{"a", "b"}) {
		t.Fatalf("Keys got %v", keys)
	}

	values := Values(input)
	slices.Sort(values)
	if !reflect.DeepEqual(values, []int{1, 2}) {
		t.Fatalf("Values got %v", values)
	}

	entries := Entries(input)
	if len(entries) != 2 {
		t.Fatalf("Entries got %d entries, want 2", len(entries))
	}
}

func TestSeqHelpersAreLazy(t *testing.T) {
	pulled := 0
	source := func(yield func(int) bool) {
		for i := 0; i < 100; i++ {
			pulled++
			if !yield(i) {
				return
			}
		}
	}

	got := Collect(Take(MapSeq(FilterSeq(source, func(n int) bool {
		return n%2 == 0
	}), func(n int) int {
		return n * 10
	}), 3))

	want := []int{0, 20, 40}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	if pulled != 5 {
		t.Fatalf("sequence pulled %d items, want 5", pulled)
	}
}

func TestSeqPredicates(t *testing.T) {
	nums := slices.Values([]int{1, 2, 3, 4})

	got, ok := FindSeq(nums, func(n int) bool {
		return n > 2
	})
	if !ok || got != 3 {
		t.Fatalf("FindSeq got %d, %v, want 3, true", got, ok)
	}

	if !AnySeq(slices.Values([]int{1, 2, 3}), func(n int) bool { return n == 2 }) {
		t.Fatal("AnySeq got false, want true")
	}

	if !AllSeq(slices.Values([]int{1, 2, 3}), func(n int) bool { return n > 0 }) {
		t.Fatal("AllSeq got false, want true")
	}
}

func TestFlatMapSeq(t *testing.T) {
	seq := FlatMapSeq(slices.Values([]int{1, 2}), func(n int) iter.Seq[int] {
		return slices.Values([]int{n, n * 10})
	})

	got := Collect(seq)
	want := []int{1, 10, 2, 20}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
