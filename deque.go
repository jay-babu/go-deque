package deque

// A Dequer holds the methods a Deque should implement.
type Dequer[T any] interface {
	Len() int
	AppendLast(...T)
	AppendFirst(...T)
	PopFirst() T
	PopLast() T
	First() T
	Last() T
	At(uint) T
	InsertAt(T, int)
}

// A Deque is the representation of a doubly-linked list.
// The zero value for DequeOld is an empty deque ready to use.
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

func (de *Deque[T]) PopFirst() T {
	first := de.l[0]
	de.l = de.l[1:]
	return first
}

func (de *Deque[T]) PopLast() T {
	last := de.l[len(de.l)-1]
	de.l = de.l[:len(de.l)-1]
	return last
}

func (de *Deque[T]) First() T {
	return de.l[0]
}

func (de *Deque[T]) Last() T {
	return de.l[len(de.l)-1]
}

func (de *Deque[T]) At(at uint) T {
	return de.l[at]
}

func (de *Deque[T]) InsertAt(val T, at int) {
	if len(de.l) == at { // nil or empty slice or after last element
		de.l = append(de.l, val)
		return
	}
	de.l = append(de.l[:at+1], de.l[at:]...) // index < len(a)
	de.l[at] = val
}
