# Dynamic Array (Vector)

## What I Implemented
A dynamic array that automatically resizes (grows and shrinks) to manage memory efficiently while maintaining contiguous storage.

## How It Works
Elements are stored contiguously in memory. When the array is full and I push, I allocate a new array with double capacity (next power of 2), copy all elements, and continue. When size drops to 25% of capacity after popping, I shrink to half capacity to reclaim memory.

## Complexity
**Time:** 
- Push: O(1) amortized (most pushes are O(1), occasional O(n) resize averages out)
- Pop: O(1) amortized (same reasoning)
- Insert/Delete: O(n) (must shift elements)
- Access by index: O(1) (direct memory access)

**Space:** O(n) for n elements, with some overhead (capacity ≥ size)

**Amortized Analysis:**
With 32 pushes starting from capacity 16:
- Resize at 16: copy 16 elements
- Resize at 32: copy 32 elements  
- Total: 32 pushes + 48 copies = 80 operations / 32 pushes = 2.5 ops per push = O(1) amortized

## When to Use
- When you don't know the final size and need dynamic growth
- When you want automatic memory shrinking (unlike Go slices which only grow)
- When you need frequent random access (O(1) indexing)

**Don't use when:**
- Frequent insertions/deletions in middle (use linked list instead)
- You know exact size beforehand (just allocate once)

## Code Example
```go
func (v *Vector) Push(item int) {
    if v.Capacity() == v.Size() {
        v.resize(v.nextPowerOfTwo())  // Double capacity
    }
    v.data[v.lastItemIndex()+1] = item
    v.size++
}

func (v *Vector) Pop() int {
    aux := v.data[v.lastItemIndex()]
    v.size--
    if v.Size() <= (v.Capacity() / 4) {  // Shrink at 25%
        v.resize(v.lastPowerOfTwo())
    }
    return aux
}
```

## Trade-offs
- **Pro:** O(1) random access, automatic memory management (grow AND shrink)
- **Con:** O(n) insert/delete in middle (shifting required), occasional O(n) resize
- **Alternative:** Linked list for O(1) front insert/delete, but O(n) access

**Why power of 2 capacities?**
Keeps amortized O(1) push. If we grew by fixed amount (+10), we'd resize too often → O(n²) total time for n pushes.

**Why shrink at 1/4 instead of 1/2?**
Prevents thrashing. At 50% full:
- With 1/2 threshold: pop→shrink→push→grow→pop→shrink (infinite loop!)
- With 1/4 threshold: Safe zone from 25% to 100%, need 8 consecutive pops (half elements) before shrinking

**Minimum capacity = 16:**
Avoids excessive resizing for small arrays.

## Interview Phrases
- "I'll implement a dynamic array with amortized O(1) push/pop"
- "This achieves O(1) amortized time because resize operations are infrequent - we only resize when doubling/halving, not on every operation"
- "The trade-off is occasional O(n) resize cost for predictable memory usage and automatic shrinking"
- "The 1/4 shrink threshold prevents thrashing - it creates a hysteresis between grow (100%) and shrink (25%)"

## Can I Answer These?
- [x] Why is it O(1) amortized? (48 copies + 32 pushes = avg 2.5 ops = O(1))
- [x] Why shrink at 1/4 not 1/2? (Prevents resize thrashing in push-pop sequences)
- [x] When would I NOT use this? (Frequent middle insertions - use linked list)
- [x] How is this different from Go slices? (Slices don't shrink automatically)

## Related Concepts
- [[Binary Search]] - requires sorted array, O(log n) search
- [[Hash Tables]] - alternative for fast lookup without ordering
- [[Linked Lists]] - alternative for frequent insertions/deletions
