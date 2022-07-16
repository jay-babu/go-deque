package deque

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	q := DequeOld[int]{}
	if q.Len() != 0 {
		t.Error("q.Len() =", q.Len(), "expect 0")
	}
}

func TestPrepend(t *testing.T) {
	q := DequeOld[int]{}
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
	q := DequeOld[int]{}
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
	q := DequeOld[int]{}
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
	q := DequeOld[int]{}
	q.AppendLast(1)
	q.AppendLast(2)
	q.AppendLast(3)
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

func TestInsertAtPanic(t *testing.T) {
	// No need to check whether `recover()` is nil. Just turn off the panic.
	defer func() { _ = recover() }()
	q := DequeOld[int]{}
	q.InsertAt(5, 3)
	t.Errorf("should have panicked")
}

func TestAt(t *testing.T) {
	q := DequeOld[int]{}
	q.AppendLast(1)
	q.AppendLast(2)
	q.AppendLast(3)
	var val int
	val, _ = q.At(0)
	if val != 1 {
		t.Error("q.At(0) =", val, "expect 1")
	}
	val, _ = q.At(1)
	if val != 2 {
		t.Error("q.At(1) =", val, "expect 2")
	}
	val, _ = q.At(2)
	if val != 3 {
		t.Error("q.At(2) =", val, "expect 3")
	}
}

func TestAtPanic(t *testing.T) {
	// No need to check whether `recover()` is nil. Just turn off the panic.
	q := DequeOld[int]{}
	_, exists := q.At(5)
	if exists == true {
		t.Error("q.At(5) should return false.")
	}
}

func TestInsertAt(t *testing.T) {
	q := DequeOld[int]{}
	q.InsertAt(0, 0)
	q.InsertAt(2, 1)
	q.InsertAt(1, 1)
	q.InsertAt(3, 3)

	if q.Len() != 4 {
		t.Error("q.Len() =", q.Len(), "expect 4")
	}

	var val int
	val, _ = q.PopFirst()
	if val != 0 {
		t.Error(val, "should have been 0")
	}
	val, _ = q.PopFirst()
	if val != 1 {
		t.Error(val, "should have been 1")
	}
	val, _ = q.PopFirst()
	if val != 2 {
		t.Error(val, "should have been 2")
	}
	val, _ = q.PopFirst()
	if val != 3 {
		t.Error(val, "should have been 3")
	}
}
