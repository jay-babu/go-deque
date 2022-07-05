package deque

import "testing"

var list []int

func BenchmarkAppendFirst_List(b *testing.B) {
	var l []int
	for n := 0; n < b.N; n++ {
		l = append(l, 1)
	}
	list = l
}

func BenchmarkAppendPopFirst_List(b *testing.B) {
	var l []int
	for n := 0; n < b.N; n++ {
		l = append(l, 1)
	}
	list = l
	for n := 0; n < b.N; n++ {
		l = l[1:]
	}
	list = l
}

func BenchmarkInsertAt_List(b *testing.B) {
	var l []int
	l = append(l, 1)
	l = append(l, 1)
	for n := 0; n < b.N; n++ {
		l = append(l[:n+1], l[n:]...)
		l[n] = 1
	}
	list = l
}
