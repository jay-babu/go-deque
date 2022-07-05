package deque

import "fmt"

// A Dequer holds the methods a Deque should implement.
type Dequer[T any] interface {
	Len() uint
	AppendLast(...T)
	AppendFirst(...T)
	PopFirst() (T, bool)
	PopLast() (T, bool)
	First() (T, bool)
	Last() (T, bool)
	At(uint) (T, bool)
	InsertAt(T, uint)
}

type node[T any] struct {
	val  T
	next *node[T]
	prev *node[T]
}

// A Deque is the representation of a doubly-linked list.
// The zero value for Deque is an empty deque ready to use.
type Deque[T any] struct {
	head  *node[T]
	tail  *node[T]
	count uint
}

var _ Dequer[int] = (*Deque[int])(nil)

// Returns the number of nodes in the Deque
func (q *Deque[T]) Len() uint {
	return q.count
}

// Appends vals of type T to the Deque.
// Input vals is read from left to right, such that the last input, will be the last val in the deque.
func (q *Deque[T]) AppendLast(vals ...T) {
	for _, val := range vals {
		q.appendLast(val)
	}
}

func (q *Deque[T]) appendLast(val T) {
	if q.count == 0 {
		q.head = &node[T]{
			val: val,
		}
		q.tail = q.head
		q.count++
		return
	}
	q.tail.next = &node[T]{
		val:  val,
		prev: q.tail,
	}
	q.tail = q.tail.next
	q.count++
}

// Appends vals of type T to the Deque.
// Input vals is read from left to right, such that the last input, will be the first val in the deque.
func (q *Deque[T]) AppendFirst(vals ...T) {
	for _, val := range vals {
		q.appendFirst(val)
	}
}

func (q *Deque[T]) appendFirst(val T) {
	if q.count == 0 {
		q.head = &node[T]{
			val: val,
		}
		q.tail = q.head
		q.count++
		return
	}
	q.head.prev = &node[T]{
		val:  val,
		next: q.head,
	}
	q.head = q.head.prev
	q.count++
}

// Removes and returns the first val in the Deque.
// Exists returns false if Deque is empty along with the zero value for val.
func (q *Deque[T]) PopFirst() (val T, exists bool) {
	if q.head == nil {
		var noop T
		return noop, false
	}
	head := q.head
	q.head = q.head.next
	head.next = nil
	if q.head != nil {
		q.head.prev = nil
	}
	q.count--
	return head.val, true
}

// Removes and returns the last val in the Deque.
// Exists returns false if Deque is empty along with the zero value for val.
func (q *Deque[T]) PopLast() (val T, exists bool) {
	if q.tail == nil {
		var noop T
		return noop, false
	}
	tail := q.tail
	q.tail = q.tail.prev
	tail.prev = nil
	if q.tail != nil {
		q.tail.next = nil
	}
	q.count--
	return tail.val, true
}

// Returns the first val in the Deque.
// Exists returns false if Deque is empty along with the zero value for val.
func (q *Deque[T]) First() (val T, exists bool) {
	if q.head == nil {
		var noop T
		return noop, false
	}
	return q.head.val, true
}

// Returns the last val in the Deque.
// Exists returns false if Deque is empty along with the zero value for val.
func (q *Deque[T]) Last() (val T, exists bool) {
	if q.tail == nil {
		var noop T
		return noop, false
	}
	return q.tail.val, true
}

// Returns the val at the index specified in the Deque.
// Exists returns false if Deque is empty along with the zero value for val.
func (q *Deque[T]) At(at uint) (val T, exists bool) {
	if at >= uint(q.count) {
		var noop T
		return noop, false
	}
	currNode := q.head
	for i := uint(0); i < at; i++ {
		currNode = currNode.next
	}
	return currNode.val, false
}

// Inserts the val at the index specified in the Deque.
// Panics if the index at is greater than number of available elements.
func (q *Deque[T]) InsertAt(val T, at uint) {
	if at > uint(q.count) {
		panic(fmt.Sprintf("Inserting at: %d greater than current number of nodes: %d", at, q.count))
	} else if at == 0 {
		q.AppendFirst(val)
		return
	} else if at == uint(q.count) {
		q.AppendLast(val)
		return
	}

	prevNode := q.head.prev
	currNode := q.head
	for i := uint(0); i < at; i++ {
		prevNode = currNode
		currNode = currNode.next
	}
	newNode := &node[T]{
		val:  val,
		prev: prevNode,
		next: currNode,
	}

	prevNode.next = newNode
	currNode.prev = newNode
	q.count++
}
