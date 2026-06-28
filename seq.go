package functional

import "iter"

// MapSeq returns a lazy sequence containing f applied to each item in seq.
func MapSeq[E, R any](seq iter.Seq[E], f func(E) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for item := range seq {
			if !yield(f(item)) {
				return
			}
		}
	}
}

// FilterSeq returns a lazy sequence containing items that satisfy keep.
func FilterSeq[E any](seq iter.Seq[E], keep func(E) bool) iter.Seq[E] {
	return func(yield func(E) bool) {
		for item := range seq {
			if keep(item) && !yield(item) {
				return
			}
		}
	}
}

// RejectSeq returns a lazy sequence containing items that do not satisfy drop.
func RejectSeq[E any](seq iter.Seq[E], drop func(E) bool) iter.Seq[E] {
	return FilterSeq(seq, func(item E) bool {
		return !drop(item)
	})
}

// ReduceSeq folds seq from left to right into a single value.
func ReduceSeq[E, R any](seq iter.Seq[E], initial R, f func(R, E) R) R {
	acc := initial
	for item := range seq {
		acc = f(acc, item)
	}
	return acc
}

// FlatMapSeq maps each item to a sequence and returns a lazy flattened sequence.
func FlatMapSeq[E, R any](seq iter.Seq[E], f func(E) iter.Seq[R]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for item := range seq {
			for mapped := range f(item) {
				if !yield(mapped) {
					return
				}
			}
		}
	}
}

// FindSeq returns the first item that satisfies match.
func FindSeq[E any](seq iter.Seq[E], match func(E) bool) (E, bool) {
	for item := range seq {
		if match(item) {
			return item, true
		}
	}
	var zero E
	return zero, false
}

// AnySeq reports whether any item satisfies match.
func AnySeq[E any](seq iter.Seq[E], match func(E) bool) bool {
	_, ok := FindSeq(seq, match)
	return ok
}

// AllSeq reports whether every item satisfies match.
func AllSeq[E any](seq iter.Seq[E], match func(E) bool) bool {
	for item := range seq {
		if !match(item) {
			return false
		}
	}
	return true
}

// Collect returns a slice containing each item in seq.
func Collect[E any](seq iter.Seq[E]) []E {
	var out []E
	for item := range seq {
		out = append(out, item)
	}
	return out
}

// Take returns a lazy sequence of at most n items from seq.
func Take[E any](seq iter.Seq[E], n int) iter.Seq[E] {
	return func(yield func(E) bool) {
		if n <= 0 {
			return
		}

		count := 0
		for item := range seq {
			if !yield(item) {
				return
			}
			count++
			if count >= n {
				return
			}
		}
	}
}
