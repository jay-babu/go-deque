package deque

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	q := Deque[int]{}
	if q.Len() != 0 {
		t.Error("q.Len() =", q.Len(), "expect 0")
	}
}

func TestPrepend(t *testing.T) {
	q := Deque[int]{}
	q.AppendFirst(1)
	if q.Len() != 1 {
		t.Error("q.Len() =", q.Len(), "expect 1")
	}
	var val int
	val, _ = q.First()
	if val != 1 {
		t.Error("q.First() =", val, "expect 1")
	}
	q.AppendFirst(2)
	if q.Len() != 2 {
		t.Error("q.Len() =", q.Len(), "expect 2")
	}
	val, _ = q.First()
	if val != 2 {
		t.Error("q.First() =", val, "expect 2")
	}
}

func TestPostpend(t *testing.T) {
	q := Deque[int]{}
	q.AppendLast(1)
	if q.Len() != 1 {
		t.Error("q.Len() =", q.Len(), "expect 1")
	}
	var val int
	val, _ = q.First()
	if val != 1 {
		t.Error("q.First() =", val, "expect 1")
	}
	q.AppendLast(2)
	if q.Len() != 2 {
		t.Error("q.Len() =", q.Len(), "expect 2")
	}
	val, _ = q.Last()
	if val != 2 {
		t.Error("q.Last() =", val, "expect 2")
	}
}

func TestPrePostpend(t *testing.T) {
	q := Deque[int]{}
	q.AppendLast(1)
	if q.Len() != 1 {
		t.Error("q.Len() =", q.Len(), "expect 1")
	}
	var val int
	val, _ = q.First()
	if val != 1 {
		t.Error("q.First() =", val, "expect 1")
	}
	q.AppendFirst(2)
	if q.Len() != 2 {
		t.Error("q.Len() =", q.Len(), "expect 2")
	}
	val, _ = q.First()
	if val != 2 {
		t.Error("q.First() =", val, "expect 2")
	}
	q.AppendLast(2)
	if q.Len() != 3 {
		t.Error("q.Len() =", q.Len(), "expect 3")
	}
	val, _ = q.Last()
	if val != 2 {
		t.Error("q.Last() =", val, "expect 2")
	}
}

func TestPops(t *testing.T) {
	q := Deque[int]{}
	q.AppendLast(1).AppendLast(2).AppendLast(3)
	var val int

	val, _ = q.PopFirst()
	if val != 1 {
		t.Error("q.Prepop() = ", val, "expect 1")
	}
	val, _ = q.PopFirst()
	if val != 2 {
		t.Error("q.Prepop() = ", val, "expect 2")
	}
	val, _ = q.PopFirst()
	if val != 3 {
		t.Error("q.Prepop() = ", val, "expect 3")
	}
}
