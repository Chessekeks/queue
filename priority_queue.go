package queue

type PriorityQueue[T any] []Item[T]

func NewPriorityQueue[T any]() PriorityQueue[T] {
	return PriorityQueue[T]{}
}

func (q *PriorityQueue[T]) Enqueue(value T, priority int) {
	item := Item[T]{
		value:    value,
		priority: priority,
	}

	qq := *q

	for i, qItem := range qq {
		if qItem.priority >= item.priority {
			continue
		}

		tail := append([]Item[T]{item}, qq[i:]...)
		qq = append(qq[:i], tail...)

		*q = qq
		return
	}

	*q = append(*q, item)
	return
}

func (q *PriorityQueue[T]) Dequeue() (T, bool) {
	var value T

	if q.isEmpty() {
		return value, false
	}

	qq := *q

	item := qq[0]
	*q = qq[1:]

	return item.value, true
}

func (q *PriorityQueue[T]) isEmpty() bool {
	return len(*q) == 0
}
