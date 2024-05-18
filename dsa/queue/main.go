package main

import "fmt"

type QueueInt struct {
	Arr   [3]int
	front int
	back  int
}

func NewQueueInt() *QueueInt {
	return &QueueInt{
		Arr:   [3]int{},
		front: 0,
		back:  0,
	}
}

func (q *QueueInt) Enqueue(x int) {
	if q.isEmpty() {
		q.front = 0
		q.back = 0
	}
	if q.isFull() {
		fmt.Println("queue is full")
		return
	}
	q.Arr[q.back] = x
	q.back++
}

func (q *QueueInt) Dequeue() {
	if q.isEmpty() {
		fmt.Println("queue is empty - can't dequeue")
		return
	}

	q.front++
}

func (q *QueueInt) isEmpty() bool {
	if q.front == q.back {
		return true
	}

	return false
}

func (q *QueueInt) isFull() bool {
	if q.back == len(q.Arr) {
		return true
	}
	return false
}

func (q *QueueInt) Front() int {
	return q.Arr[q.front]
}

func main() {
	q := NewQueueInt()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Dequeue()
	q.Enqueue(30)

	fmt.Println(q.Front())

}
