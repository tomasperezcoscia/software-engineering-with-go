package lists

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func NewNode(data int) *Node {
	return &Node{
		Data: data,
		Next: nil,
	}
}

func NewList(head *Node) *LinkedList {
	return &LinkedList{
		Head: head,
	}
}

func NewListFromArrayOfNumbers(nodes []int) *LinkedList {
	if len(nodes) == 0 {
		return &LinkedList{Head: nil}
	}

	head := NewNode(nodes[0])
	current := head

	for i := 1; i < len(nodes); i++ {
		current.Next = NewNode(nodes[i])
		current = current.Next
	}

	return &LinkedList{
		Head: head,
	}
}

func PrintList(ll *LinkedList) {
	current := ll.Head
	for current != nil {
		fmt.Printf("[ %v ]->", current.Data)
		current = current.Next
	}
	fmt.Print(" nil")
	fmt.Println()
}

func (ll *LinkedList) Size() int {
	count := 0
	current := ll.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.Head == nil
}

func (ll *LinkedList) ValueAt(index int) int {
	current := ll.Head
	count := 0
	for current != nil && index != count {
		current = current.Next
		count++
	}
	return current.Data
}

func (ll *LinkedList) PushFront(value int) {
	if ll.IsEmpty() {
		ll.Head = NewNode(value)
		return
	}
	newHead := NewNode(value)
	newHead.Next = ll.Head
	ll.Head = newHead
}

func (ll *LinkedList) PushBack(value int) {
	if ll.IsEmpty() {
		ll.Head = NewNode(value)
		return
	}
	current := ll.Head
	for current.Next != nil { // Stop at the last node (when Next is nil)
		current = current.Next
	}
	current.Next = NewNode(value) // Add new node at the end
}

func (ll *LinkedList) PopFront() int {
	if ll.IsEmpty() {
		panic("Cant pop from an empty list")
	}
	value := ll.Head.Data
	ll.Head = ll.Head.Next

	return value
}

func (ll *LinkedList) PopBack() int {
	if ll.IsEmpty() {
		panic("Cant pop from an empty list")
	}

	// Special case: only one element
	if ll.Head.Next == nil {
		value := ll.Head.Data
		ll.Head = nil
		return value
	}

	current := ll.Head
	for current.Next.Next != nil {
		current = current.Next
	}
	value := current.Next.Data
	current.Next = nil

	return value
}

func (ll *LinkedList) Front() int {
	if ll.IsEmpty() {
		panic("Cant get data from an empty list")
	}
	return ll.Head.Data
}

func (ll *LinkedList) Back() int {
	if ll.IsEmpty() {
		panic("Cant get data from an empty list")
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	return current.Data
}

func (ll *LinkedList) Insert(index int, value int) {
	if index == 0 {
		ll.PushFront(value)
		return
	}
	current := ll.Head
	count := 0
	for current != nil && count < index-1 {
		current = current.Next
		count++
	}
	if current == nil {
		panic("Index out of range")
	}
	newNode := NewNode(value)
	newNode.Next = current.Next
	current.Next = newNode
}

func (ll *LinkedList) Erase(index int) {
	if index == 0 {
		ll.PopFront()
		return
	}
	current := ll.Head
	count := 0
	for current != nil && count != index-1 {
		current = current.Next
		count++
	}
	if current == nil || current.Next == nil {
		panic("Index out of range")
	}
	current.Next = current.Next.Next
}

func (ll *LinkedList) Value_n_from_end(n int) int {
	left, right := ll.Head, ll.Head
	for i := 0; n > i; i++ {
		if right == nil {
			panic("Out of range n")
		}
		right = right.Next
	}
	for right != nil {
		right = right.Next
		left = left.Next
	}
	return left.Data
}

func (ll *LinkedList) ReverseSelf() {
	if ll.IsEmpty() {
		panic("Cant reverse an empty list")
	}
	if ll.Head.Next == nil {
		return
	}
	var prev *Node = nil
	current := ll.Head
	next := ll.Head.Next
	for next != nil {
		current.Next = prev
		prev = current
		current = next
		next = next.Next
	}
	current.Next = prev
	ll.Head = current
}

func (ll *LinkedList) RemoveValue(value int) {
	if ll.IsEmpty() {
		panic("Cant erase from empty list")
	}
	if ll.Head.Data == value {
		ll.PopFront()
		return
	}
	current := ll.Head
	for current.Next != nil && current.Next.Data != value {
		current = current.Next
	}
	current.Next = current.Next.Next
}
