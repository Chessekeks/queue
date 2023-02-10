package queue

import (
	"testing"
)

func BenchmarkQueue(b *testing.B) {
	q := NewQueue[int]()

	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}
