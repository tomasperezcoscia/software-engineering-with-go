# Queue - Linked List Implementation

## What I Implemented
A FIFO (First-In-First-Out) queue using a singly-linked list with both front and back pointers for O(1) operations at both ends.

## How It Works
I maintain two pointers: `Front` (where we dequeue) and `Back` (where we enqueue). When enqueueing, I append to the back. When dequeuing, I remove from the front. Both operations are O(1) because I have direct pointers to both ends - no traversal needed.

**Key learning:** I initially implemented this backwards (push at front, pop at back) which made PopBack O(n) because I had to traverse the entire list. The correct way is push at back (O(1) with back pointer), pop at front (O(1) - just move front pointer).

## Complexity
**Time:**
- Enqueue: O(1) - append to back using back pointer
- Dequeue: O(1) - remove from front using front pointer
- Peek/Front: O(1) - just read front.Data
- IsEmpty: O(1) - check if front is nil

**Space:** O(n) for n elements, plus pointer overhead per node

## When to Use
Use queue when you need FIFO behavior:
- **Task scheduling** (process tasks in order received)
- **Breadth-First Search (BFS)** (explore nodes level by level)
- **Print queue** (print jobs in order)
- **Message queues** (handle messages in arrival order)
- **Buffer for producer-consumer** (producer adds to back, consumer takes from front)

**Don't use when:**
- You need LIFO (use stack)
- You need random access (use array)
- You need priority ordering (use priority queue/heap)

## Code Example
```go
type QueueNode struct {
    Data int
    Next *QueueNode
}

type Queue struct {
    Front *QueueNode
    Back  *QueueNode
}

func (q *Queue) Enqueue(value int) {
    newNode := NewQueueNode(value)
    if q.Front == nil {
        // Empty queue - new node is both front and back
        q.Front = newNode
        q.Back = newNode
        return
    }
    // Add to back
    q.Back.Next = newNode
    q.Back = newNode
}

func (q *Queue) Dequeue() int {
    if q.Front == nil {
        panic("Cannot dequeue from empty queue")
    }
    if q.Front == q.Back {
        // Single element - clear both pointers
        value := q.Front.Data
        q.Front = nil
        q.Back = nil
        return value
    }
    // Remove from front
    value := q.Front.Data
    q.Front = q.Front.Next
    return value
}
```

## Trade-offs
- **Pro:** O(1) enqueue/dequeue, no capacity limit, simple to implement
- **Con:** Pointer overhead per node, worse cache performance than array, no random access
- **Alternative:** Circular buffer for fixed capacity, better cache locality

**Why two pointers (Front and Back)?**
- **Without Back pointer:** Enqueue would be O(n) - must traverse entire list to find end
- **With Back pointer:** Both operations are O(1)

**My learning moment:**
Initially implemented wrong way:
- Push at front (PushFront): O(1) ✓
- Pop at back (PopBack): O(n) ✗ - must traverse to find second-to-last node

Correct implementation:
- Push at back (with Back pointer): O(1) ✓
- Pop at front: O(1) ✓ - just move Front pointer

## Interview Phrases
- "I'll implement a queue using a linked list with front and back pointers for O(1) operations at both ends"
- "This achieves FIFO semantics - first in, first out - which is ideal for task scheduling and BFS"
- "The trade-off is pointer overhead and cache performance for unlimited capacity and O(1) operations"
- "The key insight is maintaining both front and back pointers - without back, enqueue becomes O(n)"

## Can I Answer These?
- [x] What is FIFO? (First In, First Out - like a line at a store)
- [x] Why do I need both Front and Back pointers? (Both enqueue and dequeue are O(1))
- [x] What was wrong with push-front, pop-back? (PopBack is O(n) on singly-linked list)
- [x] When would this be better than array? (Unlimited capacity, predictable O(1))
- [x] When would array be better? (Fixed capacity, better cache, less memory per element)

## Edge Cases
**Empty queue:**
```
Front = nil, Back = nil
Dequeue -> panic
```

**Single element:**
```
Front = Back = node
Must set both to nil when dequeuing
```

**Two elements:**
```
Front -> node1 -> node2 <- Back
After dequeue: Front = node2 = Back
```

## Common Interview Follow-ups

**Q: Implement a queue using two stacks.**
A: Stack1 for enqueue (push), Stack2 for dequeue (pop). When Stack2 is empty, pop all from Stack1 and push to Stack2. Amortized O(1).

**Q: Implement circular queue with array.**
A: Use readIndex and writeIndex with modulo to wrap around. See [[Queue - Circular Buffer]].

**Q: How to implement priority queue?**
A: Use heap (binary heap), not simple FIFO queue. Elements dequeue by priority, not insertion order.

## Related Concepts
- [[Stack-LIFO]] - LIFO counterpart to Queue's FIFO
- [[Queue-CircularBuffer]] - array-based alternative with fixed capacity
- [[LinkedList-ReverseInPlace]] - underlying data structure
- [[Breadth-First Search]] - uses queue for level-by-level traversal
