package arrays

const MIN_NUMBER_OF_N = 16

type Vector struct {
	data     []int
	size     int
	capacity int
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

func (v *Vector) GetData() []int {
	return v.data[:v.size]
}

func (v *Vector) IsEmpty() bool {
	return v.size == 0
}

func (v *Vector) At(index int) int {
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

func (v *Vector) Push(item int) {
	if v.Capacity() == v.Size() {
		v.resize(v.nextPowerOfTwo())
	}
	v.data[v.lastItemIndex()+1] = item
	v.size++
}

func (v *Vector) Insert(index int, item int) {
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

func (v *Vector) Prepend(item int) {
	v.Insert(0, item)
}

func (v *Vector) Pop() int {
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

func (v *Vector) Delete(index int) int {
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

func (v *Vector) Remove(item int) {
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

func (v *Vector) Find(item int) int {
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
