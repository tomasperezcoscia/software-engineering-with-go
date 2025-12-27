package queues

import "fmt"

type CircularBuffer struct {
	Buffer     []int
	Size       int
	Capacity   int
	WriteIndex int
	ReadIndex  int
}

func NewCircularBuffer(capacity int) *CircularBuffer {
	return &CircularBuffer{
		Buffer:     make([]int, capacity),
		Size:       0,
		Capacity:   capacity,
		WriteIndex: 0,
		ReadIndex:  0,
	}
}

func (cb *CircularBuffer) Enqueue(value int) {
	posWrite := cb.WriteIndex % cb.Capacity

	cb.Buffer[posWrite] = value
	cb.WriteIndex++

	if cb.Size == cb.Capacity {
		cb.ReadIndex++
	} else {
		cb.Size++
	}

	cb.Size++
}

func (cb *CircularBuffer) Dequeue() int {
	if cb.Size == 0 {
		panic("Cant dequeue from an empty queue")
	}

	posRead := cb.ReadIndex % cb.Capacity
	value := cb.Buffer[posRead]

	cb.ReadIndex++
	cb.Size--
	return value
}

func (cb *CircularBuffer) IsEmpty() bool {
	return cb.Size == 0
}

func (cb *CircularBuffer) IsFull() bool {
	return cb.Size == cb.Capacity
}

func (cb *CircularBuffer) PrintSelf() {
	for i := range cb.Buffer {
		fmt.Printf("[%v]", cb.Buffer[i])
	}
}
