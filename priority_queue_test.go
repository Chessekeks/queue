package queue

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkPriorityQueue(b *testing.B) {
	rand.Seed(time.Now().Unix())
	q := NewPriorityQueue[int]()

	for i := 0; i < b.N; i++ {
		q.Enqueue(i, rand.Intn(10))
	}

	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}
