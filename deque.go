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

func (q *Deque[T]) PopFirst() *T {
	if q.head == nil {
		return nil
	}
	head := q.head
	q.head = q.head.next
	head.next = nil
	if q.head != nil {
		q.head.prev = nil
	}
	q.count--
	return &head.val
}

func (q *Deque[T]) PopLast() *T {
	if q.tail == nil {
		return nil
	}
	tail := q.tail
	q.tail = q.tail.prev
	tail.prev = nil
	if q.tail != nil {
		q.tail.next = nil
	}
	q.count--
	return &tail.val
}

func (q *Deque[T]) First() *T {
	if q.head == nil {
		return nil
	}
	return &q.head.val
}

func (q *Deque[T]) Last() *T {
	if q.tail == nil {
		return nil
	}
	return &q.tail.val
}
