# Linked List - Reverse In-Place

## What I Implemented
In-place reversal of a singly-linked list using three pointers to flip the direction of all node pointers.

## How It Works
I traverse the list once with three pointers: `prev`, `current`, and `next`. For each node, I save the forward reference in `next` (before breaking it), then reverse `current`'s pointer to point backwards at `prev`. Then I advance all three pointers forward. At the end, `prev` points to what was the tail (now the new head).

## Complexity
**Time:** O(n) - single pass, visit each node exactly once  
**Space:** O(1) - only three pointer variables, regardless of list size (in-place reversal)

**Why O(1) space?**
No matter if the list has 10 nodes or 1,000,000 nodes, I only use three pointers. I'm rewiring existing nodes, not creating new ones.

## When to Use
- Need to reverse a linked list in-place without extra memory
- Common interview question to test pointer manipulation
- Part of more complex algorithms (palindrome checking, reversing sublists)

**Don't use when:**
- You can use extra space (could just build new list in reverse order)
- You need the original order preserved (this mutates the list)

## Code Example
```go
func (ll *LinkedList) ReverseSelf() {
    if ll.IsEmpty() || ll.Head.Next == nil {
        return  // Empty or single node - already "reversed"
    }
    
    var prev *Node = nil
    current := ll.Head
    next := ll.Head.Next
    
    for next != nil {
        current.Next = prev  // Reverse the pointer
        prev = current       // Move prev forward
        current = next       // Move current forward
        next = next.Next     // Move next forward
    }
    
    current.Next = prev  // Handle last node
    ll.Head = current    // Update head to new front (old tail)
}
```

## Visual Walkthrough
```
Original: 1 → 2 → 3 → nil

Step 1: prev=nil, current=1, next=2
After reversing: nil ← 1    2 → 3 → nil

Step 2: prev=1, current=2, next=3  
After reversing: nil ← 1 ← 2    3 → nil

Step 3: prev=2, current=3, next=nil
After reversing: nil ← 1 ← 2 ← 3

Result: prev points to 3 (new head)
```

## Trade-offs
- **Pro:** O(1) space, in-place (no extra memory), O(n) time optimal
- **Con:** Mutates original list, somewhat complex pointer logic
- **Alternative:** Create new list in reverse order - simpler but uses O(n) extra space

**Why three pointers?**
- **prev:** Tracks the new head (the node we just finished processing)
- **current:** The node we're currently reversing  
- **next:** **CRITICAL** - saves forward reference before we break it

**Why do we need `next`?**
After we execute `current.Next = prev`, we've broken the forward link. If we tried to use `current.Next` to advance, we'd go BACKWARDS (to prev) instead of forwards. The `next` pointer saves the original forward reference before we break it.

## Interview Phrases
- "I'll implement in-place reversal using three pointers to rewire the nodes"
- "This achieves O(n) time with a single pass and O(1) space - only three pointer variables regardless of list size"
- "The key insight is saving the forward reference in `next` before breaking the link with `current.Next = prev`"
- "The trade-off is pointer manipulation complexity for optimal space efficiency"

## Can I Answer These?
- [x] Why three pointers? (Need `next` to save forward ref before breaking the link)
- [x] Why is it O(1) space? (Only 3 pointers, not proportional to n)
- [x] What does `prev` represent at the end? (The new head - what was the tail)
- [x] Why set `ll.Head = prev`? (prev points to the new head after reversal)
- [x] What breaks with only two pointers? (Can't advance forward after reversing current.Next)

## Common Interview Follow-ups
**Q: Can you do this recursively?**
A: Yes, but it uses O(n) space for the call stack. Iterative is better for space efficiency.

**Q: How would you reverse only part of the list (nodes m to n)?**
A: Similar approach but track the node before m and after n to reconnect properly.

**Q: How to detect if reversal worked correctly?**
A: Walk the list and verify all pointers go in expected direction. Or reverse twice and compare to original.

## Related Concepts
- [[Linked Lists - Two Pointers]] - similar pointer manipulation technique
- [[Linked Lists - Cycle Detection]] - also uses multiple pointers
- [[Dynamic Array]] - compare O(1) space here vs O(n) for array reversal (if creating new array)
