package main

import "fmt"

const MIN_NUMBER_OF_N = 16

type Vector struct {
	data     []int
	size     int
	capacity int
}

func main() {
	v := &Vector{
		data:     make([]int, 2),
		size:     0,
		capacity: 2,
	}

	v.push(10)

	v.push(20)

	v.push(30)

	v.push(10)

	v.push(20)

	v.push(30)

	v.insert(1, 1)

	v.prepend(35)

	v.insert(1, 1)

	v.delete(1)

	v.remove(10)

	v.restartTrashToCero()

	fmt.Printf(" size=%d, cap=%d, data=%v\n", v.size, v.capacity, v.data[:v.capacity])

}

func (v *Vector) Size() int {
	return v.size
}

func (v *Vector) Capacity() int {
	return v.capacity
}

func (v *Vector) isEmpty() bool {
	return v.size == 0
}

func (v *Vector) at(index int) int {
	return v.data[index]
}

func (v *Vector) isItemAt(item int, index int) bool {
	return v.data[index] == item
}

func (v *Vector) nextPowerOfTwo() int {
	return v.capacity * 2
}

func (v *Vector) lastPowerOfTwo() int {
	return v.capacity / 2
}

func (v *Vector) resize(newCapacity int) {
	aux := make([]int, newCapacity)
	//This happens when array resizes up
	if newCapacity > v.capacity {
		for i := range v.size {
			aux[i] = v.data[i]
		}
	} else {
		//This happens when array resizes down
		for i := range newCapacity {
			aux[i] = v.data[i]
		}
	}
	v.data = aux
	v.capacity = newCapacity
}

func (v *Vector) lastItemIndex() int {
	return v.Size() - 1
}

func (v *Vector) push(item int) {
	if v.Capacity() == v.Size() {
		v.resize(v.nextPowerOfTwo())
	}
	v.data[v.lastItemIndex()+1] = item
	v.size++
}

func (v *Vector) insert(index int, item int) {
	if index > v.size || index < 0 {
		panic("Index is out of range")
	}
	if v.Capacity() == v.Size() {
		v.resize(v.nextPowerOfTwo())
	}
	for i := v.size; i > index; i-- {
		v.data[i] = v.data[i-1]
	}
	v.data[index] = item
	v.size++
}

func (v *Vector) prepend(item int) {
	v.insert(0, item)
}

func (v *Vector) pop() int {
	if v.Size() == 0 {
		panic("Cannot pop from an empty vector")
	}
	aux := v.data[v.lastItemIndex()]
	v.data[v.lastItemIndex()] = 0
	v.size--
	if v.Size() <= (v.Capacity() / 4) {
		v.resize(v.lastPowerOfTwo())
	}
	return aux
}

func (v *Vector) delete(index int) int {
	if v.Size() == 0 {
		panic("Cannot delete from an empty vector")
	}
	if index >= v.size || index < 0 {
		panic("Index is out of range")
	}
	aux := v.data[index]
	for i := index; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}
	v.size--
	if v.Size() > 0 && v.Size() <= v.Capacity()/4 {
		v.resize(v.lastPowerOfTwo())
	}
	return aux
}

func (v *Vector) remove(item int) {
	writeIndex := 0

	for readIndex := 0; readIndex < v.size; readIndex++ {
		if v.data[readIndex] != item {
			v.data[writeIndex] = v.data[readIndex]
			writeIndex++
		}
	}

	v.size = writeIndex
	if v.Size() <= v.Capacity()/4 {
		v.resize(v.lastPowerOfTwo())
	}
}

func (v *Vector) find(item int) int {
	for i := 0; i < v.size; i++ {
		if v.data[i] == item {
			return i
		}
	}
	return -1
}

// For testing purposes
func (v *Vector) restartTrashToCero() {
	for i := v.size; i < v.capacity; i++ {
		v.data[i] = 0
	}
}
