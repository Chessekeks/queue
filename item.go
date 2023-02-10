package queue

type Item[T any] struct {
	value    T
	priority int
}
