package set

import "testing"

const size = 100

func testSeqSet(start, count int, b *testing.B) Set[int] {
	b.Helper()

	items := make([]int, count)
	for i := 0; i < count; i++ {
		items[i] = start + i
	}
	return New(items...)
}

func Benchmark_Contains(b *testing.B) {
	b.Run("Set", func(b *testing.B) {
		set := testSeqSet(0, b.N, b)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = set.Contains(i)
		}
	})

	b.Run("Map", func(b *testing.B) {
		set := make(map[int]struct{}, b.N)
		for i := 0; i < b.N; i++ {
			set[i] = struct{}{}
		}
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_, _ = set[i]
		}
	})
}

func Benchmark_Equals(b *testing.B) {
	b.Run("Set", func(b *testing.B) {
		lhs := testSeqSet(0, size, b)
		rhs := testSeqSet(0, size, b)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = lhs.Equals(rhs)
		}
	})

	b.Run("Map", func(b *testing.B) {
		lhs := make(map[int]struct{}, size)
		rhs := make(map[int]struct{}, size)

		for i := 0; i < size; i++ {
			lhs[i] = struct{}{}
			rhs[i] = struct{}{}
		}
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			for key := range lhs {
				_, _ = rhs[key]
			}
		}
	})
}

func BenchmarkUnion(b *testing.B) {
	lhs := testSeqSet(0, size, b)
	rhs := testSeqSet(size, size, b)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Union(lhs, rhs)
	}
}
