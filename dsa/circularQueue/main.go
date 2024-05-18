package main

import "fmt"

type QueueInt struct {
	Arr   [10]int
	Front int
	Back  int
	Size  int
}

func New() *QueueInt {
	return &QueueInt{
		Arr:   [10]int{},
		Front: 0,
		Back:  0,
		Size:  10,
	}
}

func (q *QueueInt) Enqueue(x int) {
	// check if the queue is full
	if q.isFull() {
		fmt.Println("queue is full")
		return
	}

	if q.Back == q.Size && q.Front != 0 {
		q.Back = 0
	}

	q.Arr[q.Back] = x
	q.Back++

}

func (q *QueueInt) isFull() bool {
	if q.Back == (q.Front-1)%(q.Size-1) {
		return true
	}
	return false
}

func main() {
	q := New()
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}

	q.Enqueue(11)
	q.Enqueue(12)
}
