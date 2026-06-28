package functional

// Pair is a key and value pair.
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// Map returns a new slice containing f applied to each item in s.
func Map[E, R any](s []E, f func(E) R) []R {
	out := make([]R, len(s))
	for i, item := range s {
		out[i] = f(item)
	}
	return out
}

// Filter returns a new slice containing each item that satisfies keep.
func Filter[S ~[]E, E any](s S, keep func(E) bool) S {
	out := make(S, 0, len(s))
	for _, item := range s {
		if keep(item) {
			out = append(out, item)
		}
	}
	return out
}

// Reject returns a new slice containing each item that does not satisfy drop.
func Reject[S ~[]E, E any](s S, drop func(E) bool) S {
	return Filter(s, func(item E) bool {
		return !drop(item)
	})
}

// Reduce folds s from left to right into a single value.
func Reduce[E, R any](s []E, initial R, f func(R, E) R) R {
	acc := initial
	for _, item := range s {
		acc = f(acc, item)
	}
	return acc
}

// FlatMap maps each item to a slice and returns the concatenated result.
func FlatMap[E, R any](s []E, f func(E) []R) []R {
	out := make([]R, 0, len(s))
	for _, item := range s {
		out = append(out, f(item)...)
	}
	return out
}

// Partition splits s into two slices based on keep.
func Partition[S ~[]E, E any](s S, keep func(E) bool) (matched S, rest S) {
	matched = make(S, 0, len(s))
	rest = make(S, 0, len(s))
	for _, item := range s {
		if keep(item) {
			matched = append(matched, item)
		} else {
			rest = append(rest, item)
		}
	}
	return matched, rest
}

// Find returns the first item that satisfies match.
func Find[E any](s []E, match func(E) bool) (E, bool) {
	for _, item := range s {
		if match(item) {
			return item, true
		}
	}
	var zero E
	return zero, false
}

// Any reports whether any item satisfies match.
func Any[E any](s []E, match func(E) bool) bool {
	_, ok := Find(s, match)
	return ok
}

// All reports whether every item satisfies match.
func All[E any](s []E, match func(E) bool) bool {
	for _, item := range s {
		if !match(item) {
			return false
		}
	}
	return true
}

// Contains reports whether s contains target.
func Contains[E comparable](s []E, target E) bool {
	for _, item := range s {
		if item == target {
			return true
		}
	}
	return false
}

// GroupBy groups items by the key returned from key.
func GroupBy[E any, K comparable](s []E, key func(E) K) map[K][]E {
	out := make(map[K][]E)
	for _, item := range s {
		k := key(item)
		out[k] = append(out[k], item)
	}
	return out
}

// IndexBy indexes items by the key returned from key.
//
// If key returns the same value for more than one item, the later item wins.
func IndexBy[E any, K comparable](s []E, key func(E) K) map[K]E {
	out := make(map[K]E, len(s))
	for _, item := range s {
		out[key(item)] = item
	}
	return out
}

// Unique returns a new slice with duplicate comparable values removed.
func Unique[S ~[]E, E comparable](s S) S {
	seen := make(map[E]struct{}, len(s))
	out := make(S, 0, len(s))
	for _, item := range s {
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		out = append(out, item)
	}
	return out
}

// UniqueBy returns a new slice with duplicate keys removed.
//
// The first item for each key is kept.
func UniqueBy[S ~[]E, E any, K comparable](s S, key func(E) K) S {
	seen := make(map[K]struct{}, len(s))
	out := make(S, 0, len(s))
	for _, item := range s {
		k := key(item)
		if _, ok := seen[k]; ok {
			continue
		}
		seen[k] = struct{}{}
		out = append(out, item)
	}
	return out
}

// Keys returns the keys from m.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	out := make([]K, 0, len(m))
	for key := range m {
		out = append(out, key)
	}
	return out
}

// Values returns the values from m.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	out := make([]V, 0, len(m))
	for _, value := range m {
		out = append(out, value)
	}
	return out
}

// Entries returns the key and value pairs from m.
func Entries[M ~map[K]V, K comparable, V any](m M) []Pair[K, V] {
	out := make([]Pair[K, V], 0, len(m))
	for key, value := range m {
		out = append(out, Pair[K, V]{Key: key, Value: value})
	}
	return out
}
