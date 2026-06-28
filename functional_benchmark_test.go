package functional

import (
	"slices"
	"testing"
)

var benchmarkSink []int
var benchmarkIntSink int

func benchmarkInput(size int) []int {
	out := make([]int, size)
	for i := range out {
		out[i] = i
	}
	return out
}

func BenchmarkMap(b *testing.B) {
	input := benchmarkInput(10_000)

	b.Run("functional", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			benchmarkSink = Map(input, func(n int) int {
				return n * 2
			})
		}
	})

	b.Run("plain-loop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out := make([]int, len(input))
			for j, n := range input {
				out[j] = n * 2
			}
			benchmarkSink = out
		}
	})
}

func BenchmarkFilter(b *testing.B) {
	input := benchmarkInput(10_000)

	b.Run("functional", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			benchmarkSink = Filter(input, func(n int) bool {
				return n%2 == 0
			})
		}
	})

	b.Run("plain-loop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out := make([]int, 0, len(input))
			for _, n := range input {
				if n%2 == 0 {
					out = append(out, n)
				}
			}
			benchmarkSink = out
		}
	})
}

func BenchmarkReduce(b *testing.B) {
	input := benchmarkInput(10_000)

	b.Run("functional", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			benchmarkIntSink = Reduce(input, 0, func(total, n int) int {
				return total + n
			})
		}
	})

	b.Run("plain-loop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			total := 0
			for _, n := range input {
				total += n
			}
			benchmarkIntSink = total
		}
	})
}

func BenchmarkSeqPipeline(b *testing.B) {
	input := benchmarkInput(10_000)

	b.Run("seq", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			benchmarkSink = Collect(MapSeq(FilterSeq(slices.Values(input), func(n int) bool {
				return n%2 == 0
			}), func(n int) int {
				return n * 2
			}))
		}
	})

	b.Run("plain-loop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			out := make([]int, 0, len(input))
			for _, n := range input {
				if n%2 == 0 {
					out = append(out, n*2)
				}
			}
			benchmarkSink = out
		}
	})
}
