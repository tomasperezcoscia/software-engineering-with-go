# Stack (LIFO)

## What I Implemented
A Last-In-First-Out (LIFO) data structure. I didn't implement it as a separate class because it's essentially just using push/pop operations on one end of an existing structure (Dynamic Array or Linked List).

## How It Works
Elements can only be added (push) and removed (pop) from the top. The last element pushed is the first one to be popped - like a stack of plates. I can implement this using:
- **Dynamic Array:** push/pop at the end
- **Linked List:** pushFront/popFront at the head

I chose Dynamic Array for better performance.

## Complexity
**Time:**
- Push: O(1) amortized (using Dynamic Array)
- Pop: O(1) amortized
- Peek/Top: O(1)
- Search: O(n) (not a stack operation, defeats the purpose)

**Space:** O(n) for n elements

## When to Use
Use stacks when you need LIFO behavior:
- **Undo/Redo functionality** (each action pushed, undo pops)
- **Call stacks** (function calls and returns)
- **Backtracking algorithms** (maze traversal, DFS)
- **Browser history** (back button)
- **Expression evaluation** (parsing parentheses, postfix notation)
- **Reversing** (push all, then pop all = reversed order)

**Don't use when:**
- You need random access (use array)
- You need FIFO (use queue)
- You need to search efficiently (use hash table)

## Implementation Choice: Array vs Linked List

**Why I chose Dynamic Array over Linked List:**

| Aspect | Dynamic Array | Linked List |
|--------|---------------|-------------|
| Push/Pop time | O(1) amortized | O(1) |
| Cache performance | ✅ Excellent (contiguous memory) | ❌ Poor (scattered nodes) |
| Memory overhead | ✅ Small (just array) | ❌ Pointer per node |
| Real-world speed | ✅ Faster | ❌ Slower |

**Verdict:** Dynamic Array is better for stacks in practice due to cache locality and less memory overhead.

## Code Example (Using Dynamic Array)
```go
type Stack struct {
    data []int
}

func (s *Stack) Push(value int) {
    s.data = append(s.data, value)  // Add to end
}

func (s *Stack) Pop() int {
    if s.IsEmpty() {
        panic("Empty stack")
    }
    value := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]  // Remove from end
    return value
}

func (s *Stack) Top() int {
    if s.IsEmpty() {
        panic("Empty stack")
    }
    return s.data[len(s.data)-1]  // Peek without removing
}

func (s *Stack) IsEmpty() bool {
    return len(s.data) == 0
}
```

## Trade-offs
- **Pro:** O(1) push/pop, simple concept, efficient with array implementation
- **Con:** No random access, search requires popping everything (destroys stack)
- **Alternative:** Queue for FIFO, Deque for both ends

## Advanced: Stack with O(1) getMin()

**Problem:** Implement a stack that supports push, pop, AND getMin in O(1) time.

**Challenge:**
```
push(5)  -> min = 5
push(3)  -> min = 3
push(7)  -> min = 3
pop()    -> min = 3
pop()    -> min = 5  (how to know this without scanning?)
```

**Solution: Parallel Min Stack**

Maintain two stacks:
```
Main Stack:     [5, 3, 7]
Min Stack:      [5, 3, 3]
                     ↑
                 getMin() = 3
```

**How it works:**
- On push: if `value <= currentMin`, push to both stacks. Otherwise only main stack.
- On pop: if `popped == minStack.top()`, pop from both. Otherwise only main stack.
- getMin: return `minStack.top()` → O(1)

**Space optimization example:**
```
Main Stack:     [3, 7, 3, 4]
Min Stack:      [3, 3]        (only store when new min found)
                     ↑
                getMin() = 3
```

**Complexity:**
- Push/Pop/GetMin: O(1)
- Space: O(n) worst case (decreasing sequence), much better in practice

## Interview Phrases
- "I'll implement a stack using a dynamic array for O(1) push/pop with better cache performance"
- "Stack provides LIFO semantics - last in, first out - which is ideal for backtracking and undo operations"
- "The trade-off is we lose random access, but gain O(1) add/remove from one end"
- "For getMin in O(1), I'll maintain a parallel min stack that tracks the minimum at each level"

## Can I Answer These?
- [x] What is LIFO? (Last In, First Out - like a stack of plates)
- [x] Why array over linked list? (Cache locality, less memory overhead)
- [x] Why can't I use this for FIFO? (Wrong order - need queue)
- [x] How does getMin in O(1) work? (Parallel min stack tracks min at each level)
- [x] What's the space complexity of getMin solution? (O(n) worst case, optimized by only storing when new min)

## Common Interview Follow-ups

**Q: Implement a queue using two stacks.**
A: Use one stack for enqueue (push), one for dequeue (pop). When dequeue stack is empty, pop all from enqueue stack and push to dequeue stack. Amortized O(1).

**Q: Check if parentheses are balanced: "((()))"**
A: Push opening parens, pop on closing. If stack is empty at end and never tried to pop empty = balanced.

**Q: Evaluate postfix expression: "3 4 + 5 *"**
A: Push numbers. On operator, pop two operands, apply operator, push result. Final stack top is answer.

## Related Concepts
- [[Dynamic Array]] - my implementation choice for Stack
- [[Queue]] - FIFO counterpart to Stack's LIFO
- [[Linked Lists]] - alternative implementation
- [[Depth-First Search]] - uses stack (implicit via recursion or explicit)
