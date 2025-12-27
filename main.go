package main

import (
	"fmt"
	"software-engineering-with-go/arrays"
	"software-engineering-with-go/queues"
)

func main() {
	//testVectors()
	// testLinkedLists()
	testQueues()
}

func testVectors() {
	fmt.Println("=== Vector Implementation Test ===")
	fmt.Println()

	// Test 1: NewVector with capacity < 16 (should default to 16)
	fmt.Println("Test 1: Creating vector with capacity 5 (should use 16)")
	v := arrays.NewVector(5)
	fmt.Printf("  size=%d, capacity=%d\n\n", v.Size(), v.Capacity())

	// Test 2: NewVector with capacity > 16 (should round to power of 2)
	fmt.Println("Test 2: Creating vector with capacity 20 (should use 32)")
	v2 := arrays.NewVector(20)
	fmt.Printf("  size=%d, capacity=%d\n\n", v2.Size(), v2.Capacity())

	// Test 3: Push elements and test resize up
	fmt.Println("Test 3: Pushing elements to test auto-resize")
	v = arrays.NewVector(0)
	for i := 1; i <= 20; i++ {
		v.Push(i * 10)
	}
	fmt.Printf("  After 20 pushes: size=%d, capacity=%d\n", v.Size(), v.Capacity())
	fmt.Printf("  Data: %v\n\n", v.GetData())

	// Test 4: at() with bounds checking
	fmt.Println("Test 4: Testing At() method")
	fmt.Printf("  At(5) = %d\n", v.At(5))
	fmt.Printf("  At(15) = %d\n\n", v.At(15))

	// Test 5: insert and prepend
	fmt.Println("Test 5: Testing Insert and Prepend")
	v.Insert(10, 999)
	fmt.Printf("  After Insert(10, 999): %v\n", v.GetData())
	v.Prepend(5)
	fmt.Printf("  After Prepend(5): %v\n\n", v.GetData())

	// Test 6: find
	fmt.Println("Test 6: Testing Find")
	fmt.Printf("  Find(999) = %d\n", v.Find(999))
	fmt.Printf("  Find(100) = %d\n", v.Find(100))
	fmt.Printf("  Find(5000) = %d (not found)\n\n", v.Find(5000))

	// Test 7: delete
	fmt.Println("Test 7: Testing Delete")
	deleted := v.Delete(11)
	fmt.Printf("  Deleted value at index 11: %d\n", deleted)
	fmt.Printf("  After Delete: size=%d, capacity=%d\n\n", v.Size(), v.Capacity())

	// Test 8: remove (removes all occurrences)
	fmt.Println("Test 8: Testing Remove")
	v.Push(100)
	v.Push(100)
	v.Push(100)
	fmt.Printf("  Before Remove(100): %v\n", v.GetData())
	v.Remove(100)
	fmt.Printf("  After Remove(100): %v\n\n", v.GetData())

	// Test 9: pop and resize down
	fmt.Println("Test 9: Testing Pop and auto resize down")
	fmt.Printf("  Before popping: size=%d, capacity=%d\n", v.Size(), v.Capacity())
	for v.Size() > 5 {
		v.Pop()
	}
	fmt.Printf("  After popping to 5 elements: size=%d, capacity=%d\n", v.Size(), v.Capacity())
	fmt.Printf("  Final data: %v\n\n", v.GetData())

	// Test 10: isEmpty
	fmt.Println("Test 10: Testing IsEmpty")
	fmt.Printf("  IsEmpty() = %v\n", v.IsEmpty())
	for v.Size() > 0 {
		v.Pop()
	}
	fmt.Printf("  After popping all: IsEmpty() = %v\n", v.IsEmpty())
	fmt.Printf("  Final capacity (should be 16): %d\n", v.Capacity())

	fmt.Println("\n=== All tests completed! ===")
}

func testLinkedLists() {
	fmt.Println("=== LinkedList Implementation Test ===")
	fmt.Println()

	// Test 1: Create list from array
	// fmt.Println("Test 1: Creating list from array [45, 34, 56]")
	// ll := lists.NewListFromArrayOfNumbers([]int{45, 34, 56})
	// fmt.Printf("  Size: %d\n", ll.Size())
	// fmt.Printf("  ValueAt(0): %d\n", ll.ValueAt(0))
	// fmt.Printf("  ValueAt(1): %d\n", ll.ValueAt(1))
	// fmt.Printf("  ValueAt(2): %d\n", ll.ValueAt(2))
	// fmt.Println()

	// Test 2: Create list manually
	// fmt.Println("Test 2: Creating list manually")
	// head := lists.NewNode(10)
	// ll2 := lists.NewList(head)
	// ll2.Head.Next = lists.NewNode(20)
	// ll2.Head.Next.Next = lists.NewNode(30)
	// fmt.Printf("  Size: %d\n", ll2.Size())
	// fmt.Printf("  ValueAt(0): %d\n", ll2.ValueAt(0))
	// fmt.Printf("  ValueAt(1): %d\n", ll2.ValueAt(1))
	// fmt.Printf("  ValueAt(2): %d\n", ll2.ValueAt(2))
	// fmt.Println()

	// Test 3: IsEmpty
	// fmt.Println("Test 3: Testing IsEmpty")
	// emptyList := lists.NewListFromArrayOfNumbers([]int{})
	// fmt.Printf("  Empty list IsEmpty(): %v\n", emptyList.IsEmpty())
	// fmt.Printf("  Non-empty list IsEmpty(): %v\n", ll.IsEmpty())

	// Test 4: PushFront
	// fmt.Println("Test 4: Testing PushFront and PopBack")
	// ll3 := lists.NewListFromArrayOfNumbers([]int{1, 2, 3})
	// ll3.PushFront(4)
	// lists.PrintList(ll3)
	// ll3.PopFront()
	// ll3.PopFront()
	// lists.PrintList(ll3)
	// ll3.PushBack(33)
	// lists.PrintList(ll3)
	// ll3.PopBack()
	// ll3.PopBack()
	// lists.PrintList(ll3)

	// Test edge case: PopBack with single element
	// fmt.Println("\nTest 5: PopBack with single element")
	// ll4 := lists.NewListFromArrayOfNumbers([]int{100})
	// fmt.Printf("Before PopBack: ")
	// lists.PrintList(ll4)
	// val := ll4.PopBack()
	// fmt.Printf("Popped value: %d\n", val)
	// fmt.Printf("After PopBack (should be empty): ")
	// lists.PrintList(ll4)
	// fmt.Printf("IsEmpty: %v\n", ll4.IsEmpty())

	// Test 6: Front and Back data
	// fmt.Println("\n Test6: Front() and Back() should return data")
	// ll5 := lists.NewListFromArrayOfNumbers([]int{33, 66})
	// fmt.Println(ll5.Front())
	// fmt.Println(ll5.Back())

	// Test 7: Insert at various positions
	// fmt.Println("\nTest 7: Insert value at index")
	// ll6 := lists.NewListFromArrayOfNumbers([]int{33, 66})
	// fmt.Printf("Initial: ")
	// lists.PrintList(ll6)

	// ll6.Insert(0, 2)   // Insert at beginning
	// fmt.Printf("After Insert(0, 2): ")
	// lists.PrintList(ll6)

	// ll6.Insert(2, 99)  // Insert in middle
	// fmt.Printf("After Insert(2, 99): ")
	// lists.PrintList(ll6)

	// ll6.Insert(4, 88)  // Insert at end
	// fmt.Printf("After Insert(4, 88): ")
	// lists.PrintList(ll6)

	// // Test edge case: Insert into empty list
	// fmt.Println("\nTest 8: Insert into empty list")
	// ll7 := lists.NewListFromArrayOfNumbers([]int{})
	// ll7.Insert(0, 100)
	// fmt.Printf("After Insert(0, 100): ")
	// lists.PrintList(ll7)

	// Test 8: Erase at index
	// ll8 := lists.NewListFromArrayOfNumbers([]int{1, 4, 8})
	// ll8.RemoveValue(1)
	// lists.PrintList(ll8)

	// Test 9: Value_n_from_end
	// ll9 := lists.NewListFromArrayOfNumbers([]int{1, 2, 3, 4, 5, 6, 7})
	// fmt.Print(ll9.Value_n_from_end(7))

	// Test 10: Reverse list O(1)
	// ll10 := lists.NewListFromArrayOfNumbers([]int{1, 2, 3, 4, 5, 6, 7})
	// ll10.ReverseSelf()
	// lists.PrintList(ll10)

	fmt.Println("\n=== All tests completed! ===")
}

func testQueues() {
	q := queues.NewQueue(nil, nil)
	q.Enqueue(34)
	q.Enqueue(24)
	q.Enqueue(35)
	q.Enqueue(55)

	fmt.Printf("Dequeued: %v\n", q.Dequeue())
	fmt.Printf("Dequeued: %v\n", q.Dequeue())
	fmt.Printf("Dequeued: %v\n", q.Dequeue())
	fmt.Printf("Dequeued: %v\n", q.Dequeue())

	q.PrintSelf()
}
