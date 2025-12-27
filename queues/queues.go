package queues

import "fmt"

type QueueNode struct {
	Data int
	Next *QueueNode
}

type Queue struct {
	Front *QueueNode
	Back  *QueueNode
}

func NewQueueNode(data int) *QueueNode {
	return &QueueNode{
		Data: data,
		Next: nil,
	}
}

func NewQueue(front *QueueNode, back *QueueNode) *Queue {
	return &Queue{
		Front: front,
		Back:  back,
	}
}

func (q *Queue) Enqueue(value int) {
	newNode := NewQueueNode(value)
	if q.Front == nil {
		q.Front = newNode
		q.Back = newNode
		return
	}
	q.Back.Next = newNode
	q.Back = newNode
}

func (q *Queue) Dequeue() int {
	if q.Front == nil {
		panic("Cannot dequeue from an empty queue")
	}
	if q.Front == q.Back {
		value := q.Front.Data
		q.Front = nil
		q.Back = nil
		return value
	}
	value := q.Front.Data
	q.Front = q.Front.Next
	return value
}

func (q *Queue) PrintSelf() {
	current := q.Front
	for current != nil {
		fmt.Printf("|%v|", current.Data)
		current = current.Next
	}
}
