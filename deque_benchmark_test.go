package deque

import "testing"

var (
	queue  Deque[int]
	result int
)

func BenchmarkAppendFirst(b *testing.B) {
	q := Deque[int]{}
	for n := 0; n < b.N; n++ {
		q.AppendFirst(1)
	}
	queue = q
}

func BenchmarkAppendPopFirst(b *testing.B) {
	q := Deque[int]{}
	for n := 0; n < b.N; n++ {
		q.AppendFirst(1)
	}
	queue = q
	for n := 0; n < b.N; n++ {
		result, _ = q.PopFirst()
	}
}

func BenchmarkInsertAt(b *testing.B) {
	q := Deque[int]{}
	q.AppendFirst(1)
	q.AppendFirst(1)
	for n := 0; n < b.N; n++ {
		q.InsertAt(1, 1)
	}
	queue = q
}
