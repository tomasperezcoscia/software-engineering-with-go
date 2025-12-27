package hash_tables

type hashTable struct {
	Keys []int
	Size int
}

func NewHashTable(size int) *hashTable {
	return &hashTable{
		Keys: make([]int, size),
		Size: size,
	}
}

func NumericValueOf(c rune) int {
	return int(c)
}

func (ht *hashTable) Hash(key string) int {
	total := 0
	for _, char := range key {
		total += NumericValueOf(char)
	}
	return total % ht.Size
}
