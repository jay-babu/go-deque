package deque

import "fmt"

type Dequer[T any] interface {
	Len() int
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

type Deque[T any] struct {
	head  *node[T]
	tail  *node[T]
	count int
}

var _ Dequer[int] = (*Deque[int])(nil)

func (q *Deque[T]) Len() int {
	return q.count
}

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

func (q *Deque[T]) PopFirst() (T, bool) {
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

func (q *Deque[T]) PopLast() (T, bool) {
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

func (q *Deque[T]) First() (T, bool) {
	if q.head == nil {
		var noop T
		return noop, false
	}
	return q.head.val, true
}

func (q *Deque[T]) Last() (T, bool) {
	if q.tail == nil {
		var noop T
		return noop, false
	}
	return q.tail.val, true
}

func (q *Deque[T]) At(at uint) (T, bool) {
	if at >= uint(q.count) {
		panic(fmt.Sprintf("At Index: %d greater than current number of nodes: %d", at, q.count))
	}
	currNode := q.head
	for i := uint(0); i < at; i++ {
		currNode = currNode.next
	}
	return currNode.val, false
}

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
