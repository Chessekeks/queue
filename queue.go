package queue

type Queue[T any] []T

func NewQueue[T any]() Queue[T] {
	return Queue[T]{}
}

func (q *Queue[T]) Enqueue(value T) {
	*q = append(*q, value)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var item T

	if q.isEmpty() {
		return item, false
	}

	qq := *q

	item = qq[0]
	*q = qq[1:]

	return item, true
}

func (q *Queue[T]) isEmpty() bool {
	return len(*q) == 0
}
