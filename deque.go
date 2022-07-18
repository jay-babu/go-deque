package deque

// A Dequer holds the methods a Deque should implement.
type Dequer[T any] interface {
	Len() int
	AppendLast(...T)
	AppendFirst(...T)
	PopFirst() (T, bool)
	PopLast() (T, bool)
	First() (T, bool)
	Last() (T, bool)
	At(uint) (T, bool)
	InsertAt(T, int)
}

// A Deque is the representation of a doubly-linked list.
// The zero value for Deque is an empty deque ready to use.
type Deque[T any] struct {
	l []T
}

var _ Dequer[int] = (*Deque[int])(nil)

func (de *Deque[T]) Len() int {
	return len(de.l)
}

func (de *Deque[T]) AppendLast(vals ...T) {
	for _, val := range vals {
		de.appendLast(val)
	}
}

func (de *Deque[T]) appendLast(val T) {
	de.l = append(de.l, val)
}

func (de *Deque[T]) AppendFirst(vals ...T) {
	for _, val := range vals {
		de.appendFirst(val)
	}
}

func (de *Deque[T]) appendFirst(val T) {
	de.l = append([]T{val}, de.l...)
}

func (de *Deque[T]) PopFirst() (T, bool) {
	if de.Len() <= 0 {
		return de.t()
	}
	first := de.l[0]
	de.l = de.l[1:]
	return first, true
}

func (de *Deque[T]) PopLast() (T, bool) {
	if de.Len() <= 0 {
		return de.t()
	}
	last := de.l[len(de.l)-1]
	de.l = de.l[:len(de.l)-1]
	return last, true
}

func (de *Deque[T]) First() (T, bool) {
	if de.Len() <= 0 {
		return de.t()
	}
	return de.l[0], true
}

func (de *Deque[T]) Last() (T, bool) {
	if de.Len() <= 0 {
		return de.t()
	}
	return de.l[len(de.l)-1], true
}

func (de *Deque[T]) At(at uint) (T, bool) {
	if de.Len() <= 0 {
		return de.t()
	}
	if int(at) > de.Len()-1 {
		return de.t()
	}
	return de.l[at], true
}

func (de *Deque[T]) InsertAt(val T, at int) {
	if len(de.l) == at { // nil or empty slice or after last element
		de.l = append(de.l, val)
		return
	}
	de.l = append(de.l[:at+1], de.l[at:]...) // index < len(a)
	de.l[at] = val
}

func (de *Deque[T]) t() (T, bool) {
	var t T
	return t, false
}
