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

func BenchmarkSet_Contains(b *testing.B) {
	b.StopTimer()
	set := testSeqSet(0, b.N, b)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = set.Contains(i)
	}
}

func BenchmarkMap_Contains(b *testing.B) {
	b.StopTimer()
	set := make(map[int]struct{}, b.N)
	for i := 0; i < b.N; i++ {
		set[i] = struct{}{}
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, _ = set[i]
	}
}

func BenchmarkSet_Equals(b *testing.B) {
	b.StopTimer()
	lhs := testSeqSet(0, size, b)
	rhs := testSeqSet(0, size, b)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = lhs.Equals(rhs)
	}
}

func BenchmarkMap_Equals(b *testing.B) {
	b.StopTimer()
	lhs := make(map[int]struct{}, size)
	rhs := make(map[int]struct{}, size)

	for i := 0; i < size; i++ {
		lhs[i] = struct{}{}
		rhs[i] = struct{}{}
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for key := range lhs {
			_, _ = rhs[key]
		}
	}
}

func BenchmarkUnion(b *testing.B) {
	b.StopTimer()
	lhs := testSeqSet(0, size, b)
	rhs := testSeqSet(size, size, b)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = Union(lhs, rhs)
	}
}
