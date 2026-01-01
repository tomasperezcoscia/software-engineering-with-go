package hash_tables

import "fmt"

type Node struct {
	Key   string
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

type HashTable struct {
	Keys []*LinkedList
	Size int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		Keys: make([]*LinkedList, size),
		Size: size,
	}
}

func NewNode(key string, value int) *Node {
	return &Node{
		Key:   key,
		Value: value,
		Next:  nil,
	}
}

func NewList(head *Node) *LinkedList {
	return &LinkedList{
		Head: head,
	}
}

func (ll *LinkedList) AddOrUpdate(key string, value int) {
	current := ll.Head
	for current.Next != nil && current.Key != key {
		current = current.Next
	}
	if current.Key == key {
		current.Value = value
		return
	}
	newNode := NewNode(key, value)
	current.Next = newNode
}

func (ll *LinkedList) Delete(key string) {
	current := ll.Head
	if current == nil {
		panic("Cant delete from empty list")
	}
	if current.Key == key {
		ll.Head = ll.Head.Next
	}
	for current.Next != nil {
		if current.Next.Key == key {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
	/*if current == nil {
		panic("cant delete from nil linked list")
	}
	if current.Key == key {
		ll.Head = ll.Head.Next
		return
	}
	for current.Next != nil && current.Next.Key != key {
		current = current.Next
	}
	if current.Next.Key == key {
		current.Next = current.Next.Next
		return
	}*/
}

func (ll *LinkedList) ValueAt(key string) int {
	current := ll.Head
	for current != nil {
		if current.Key == key {
			return current.Value
		}
		current = current.Next
	}
	return -1
}

func NumericValueOf(c rune) int {
	return int(c)
}

func (ht *HashTable) Hash(key string) int {
	total := 0
	for _, char := range key {
		total += NumericValueOf(char)
	}
	return total % ht.Size
}

func (ht *HashTable) Insert(key string, value int) {
	if ht.Size == 0 {
		panic("Cant insert in a 0 size dictionary")
	}

	hashedKey := ht.Hash(key)

	if ht.Keys[hashedKey] == nil {
		newNode := NewNode(key, value)
		ht.Keys[hashedKey] = NewList(newNode)
	} else {
		ht.Keys[hashedKey].AddOrUpdate(key, value)
	}
}

func (ht *HashTable) Delete(key string) {
	if ht.Size == 0 {
		panic("Cant delete from a 0 size dictionary")
	}

	hashedKey := ht.Hash(key)
	ht.Keys[hashedKey].Delete(key)
}

func (ht *HashTable) ValueAt(key string) int {
	hashedKey := ht.Hash(key)
	return ht.Keys[hashedKey].ValueAt(key)
}

func (ht *HashTable) PrintSelf() {
	fmt.Println("--- Hash Table Content ---")
	for i := 0; i < ht.Size; i++ {
		fmt.Printf("Bucket [%d]: ", i)

		if ht.Keys[i] != nil {
			current := ht.Keys[i].Head
			for current != nil {
				fmt.Printf("{%s: %d} -> ", current.Key, current.Value)
				current = current.Next
			}
			fmt.Print("nil")
		} else {
			fmt.Print("nil")
		}
		fmt.Println()
	}
	fmt.Println("--------------------------")
}
