package deque

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

func (q *Deque[T]) Len() int {
	return q.count
}

func (q *Deque[T]) AppendLast(val T) *Deque[T] {
	if q.count == 0 {
		q.head = &node[T]{
			val: val,
		}
		q.tail = q.head
		q.count++
		return q
	}
	q.tail.next = &node[T]{
		val:  val,
		prev: q.tail,
	}
	q.tail = q.tail.next
	q.count++
	return q
}

func (q *Deque[T]) AppendFirst(val T) *Deque[T] {
	if q.count == 0 {
		q.head = &node[T]{
			val: val,
		}
		q.tail = q.head
		q.count++
		return q
	}
	q.head.prev = &node[T]{
		val:  val,
		next: q.head,
	}
	q.head = q.head.prev
	q.count++
	return q
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

// func (q *Deque[T]) InsertAt(val T, at uint) *Deque[T] {
// }
