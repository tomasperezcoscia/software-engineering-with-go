package main

import "fmt"

const MIN_NUMBER_OF_N = 16

type Vector struct {
	data     []int
	size     int
	capacity int
}

func main() {
	fmt.Println("=== Vector Implementation Test ===")
	fmt.Println()

	// Test 1: NewVector with capacity < 16 (should default to 16)
	fmt.Println("Test 1: Creating vector with capacity 5 (should use 16)")
	v := NewVector(5)
	fmt.Printf("  size=%d, capacity=%d\n\n", v.Size(), v.Capacity())

	// Test 2: NewVector with capacity > 16 (should round to power of 2)
	fmt.Println("Test 2: Creating vector with capacity 20 (should use 32)")
	v2 := NewVector(20)
	fmt.Printf("  size=%d, capacity=%d\n\n", v2.Size(), v2.Capacity())

	// Test 3: Push elements and test resize up
	fmt.Println("Test 3: Pushing elements to test auto-resize")
	v = NewVector(0)
	for i := 1; i <= 20; i++ {
		v.push(i * 10)
	}
	fmt.Printf("  After 20 pushes: size=%d, capacity=%d\n", v.Size(), v.Capacity())
	fmt.Printf("  Data: %v\n\n", v.data[:v.size])

	// Test 4: at() with bounds checking
	fmt.Println("Test 4: Testing at() method")
	fmt.Printf("  at(5) = %d\n", v.at(5))
	fmt.Printf("  at(15) = %d\n\n", v.at(15))

	// Test 5: insert and prepend
	fmt.Println("Test 5: Testing insert and prepend")
	v.insert(10, 999)
	fmt.Printf("  After insert(10, 999): %v\n", v.data[:v.size])
	v.prepend(5)
	fmt.Printf("  After prepend(5): %v\n\n", v.data[:v.size])

	// Test 6: find
	fmt.Println("Test 6: Testing find")
	fmt.Printf("  find(999) = %d\n", v.find(999))
	fmt.Printf("  find(100) = %d\n", v.find(100))
	fmt.Printf("  find(5000) = %d (not found)\n\n", v.find(5000))

	// Test 7: delete
	fmt.Println("Test 7: Testing delete")
	deleted := v.delete(11)
	fmt.Printf("  Deleted value at index 11: %d\n", deleted)
	fmt.Printf("  After delete: size=%d, capacity=%d\n\n", v.Size(), v.Capacity())

	// Test 8: remove (removes all occurrences)
	fmt.Println("Test 8: Testing remove")
	v.push(100)
	v.push(100)
	v.push(100)
	fmt.Printf("  Before remove(100): %v\n", v.data[:v.size])
	v.remove(100)
	fmt.Printf("  After remove(100): %v\n\n", v.data[:v.size])

	// Test 9: pop and resize down
	fmt.Println("Test 9: Testing pop and auto resize down")
	fmt.Printf("  Before popping: size=%d, capacity=%d\n", v.Size(), v.Capacity())
	for v.Size() > 5 {
		v.pop()
	}
	fmt.Printf("  After popping to 5 elements: size=%d, capacity=%d\n", v.Size(), v.Capacity())
	fmt.Printf("  Final data: %v\n\n", v.data[:v.size])

	// Test 10: isEmpty
	fmt.Println("Test 10: Testing isEmpty")
	fmt.Printf("  isEmpty() = %v\n", v.isEmpty())
	for v.Size() > 0 {
		v.pop()
	}
	fmt.Printf("  After popping all: isEmpty() = %v\n", v.isEmpty())
	fmt.Printf("  Final capacity (should be 16): %d\n", v.Capacity())

	fmt.Println("\n=== All tests completed! ===")
}

// NewVector creates a new vector with the given initial capacity
// If initialCapacity < 16, defaults to 16
// Otherwise rounds up to next power of 2
func NewVector(initialCapacity int) *Vector {
	capacity := calculateInitialCapacity(initialCapacity)

	return &Vector{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

// calculateInitialCapacity returns the proper power-of-2 capacity
func calculateInitialCapacity(n int) int {
	if n < MIN_NUMBER_OF_N {
		return MIN_NUMBER_OF_N
	}

	power := MIN_NUMBER_OF_N
	for power < n {
		power *= 2
	}
	return power
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
	if index < 0 || index >= v.size {
		panic("Index out of bounds")
	}
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
	// Ensure we never go below minimum capacity
	if newCapacity < MIN_NUMBER_OF_N {
		newCapacity = MIN_NUMBER_OF_N
	}

	aux := make([]int, newCapacity)
	// This happens when array resizes up
	if newCapacity > v.capacity {
		for i := range v.size {
			aux[i] = v.data[i]
		}
	} else {
		// This happens when array resizes down
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
