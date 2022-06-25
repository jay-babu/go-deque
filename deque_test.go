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
	q.Prepend(1)
	if q.Len() != 1 {
		t.Error("q.Len() =", q.Len(), "expect 1")
	}
	if *q.First() != 1 {
		t.Error("q.First() =", *q.First(), "expect 1")
	}
	q.Prepend(2)
	if q.Len() != 2 {
		t.Error("q.Len() =", q.Len(), "expect 2")
	}
	if *q.First() != 2 {
		t.Error("q.First() =", *q.First(), "expect 2")
	}
}

func TestPostpend(t *testing.T) {
	q := Deque[int]{}
	q.Postpend(1)
	if q.Len() != 1 {
		t.Error("q.Len() =", q.Len(), "expect 1")
	}
	if *q.First() != 1 {
		t.Error("q.First() =", *q.First(), "expect 1")
	}
	q.Postpend(2)
	if q.Len() != 2 {
		t.Error("q.Len() =", q.Len(), "expect 2")
	}
	if *q.Last() != 2 {
		t.Error("q.Last() =", *q.Last(), "expect 2")
	}
}

func TestPrePostpend(t *testing.T) {
	q := Deque[int]{}
	q.Postpend(1)
	if q.Len() != 1 {
		t.Error("q.Len() =", q.Len(), "expect 1")
	}
	if *q.First() != 1 {
		t.Error("q.First() =", *q.First(), "expect 1")
	}
	q.Prepend(2)
	if q.Len() != 2 {
		t.Error("q.Len() =", q.Len(), "expect 2")
	}
	if *q.First() != 2 {
		t.Error("q.First() =", *q.First(), "expect 2")
	}
	q.Postpend(2)
	if q.Len() != 3 {
		t.Error("q.Len() =", q.Len(), "expect 3")
	}
	if *q.Last() != 2 {
		t.Error("q.Last() =", *q.Last(), "expect 2")
	}
}

func TestPops(t *testing.T) {
	q := Deque[int]{}
	q.Postpend(1)
	q.Postpend(2)
	q.Postpend(3)
	var val *int

	val = q.Prepop()
	if *val != 1 {
		t.Error("q.Prepop() = ", *val, "expect 1")
	}
	val = q.Prepop()
	if *val != 2 {
		t.Error("q.Prepop() = ", *val, "expect 2")
	}
	val = q.Postpop()
	if *val != 3 {
		t.Error("q.Prepop() = ", *val, "expect 3")
	}
}
