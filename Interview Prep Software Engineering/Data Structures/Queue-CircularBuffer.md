# Queue - Circular Buffer Implementation

## What I Implemented
A FIFO queue using a fixed-size array with circular wraparound (via modulo) to reuse space efficiently. When the buffer is full, it overwrites the oldest data.

## How It Works
I maintain a fixed-size array with `ReadIndex` and `WriteIndex`. Both indices increment forever, but I use modulo (`% Capacity`) to wrap them around the array. When enqueueing, I write to `Buffer[WriteIndex % Capacity]` and increment WriteIndex. When dequeuing, I read from `Buffer[ReadIndex % Capacity]` and increment ReadIndex. If the buffer is full when enqueueing, I also advance ReadIndex (overwriting oldest unread data).

**Why circular?** Without wraparound, indices keep growing (0,1,2,3,4...) and you can't reuse indices 0-2 after dequeuing - it's wasted space forever. Modulo wraps: `WriteIndex=25, Capacity=8 → 25%8=1`, reusing position 1.

## Complexity
**Time:**
- Enqueue: O(1) - arithmetic and write
- Dequeue: O(1) - arithmetic and read
- IsEmpty/IsFull: O(1) - simple comparisons

**Space:** O(capacity) - fixed size buffer allocated upfront

## When to Use
Use circular buffer when you need:
- **Fixed memory footprint** (embedded systems, real-time systems)
- **Predictable performance** (no dynamic allocation)
- **Overwrite semantics** (newest data matters, old data can be dropped)

**Common use cases:**
- Audio/video buffers (drop old frames if consumer is slow)
- Sensor data logging (keep recent readings)
- Real-time systems (bounded memory)
- Producer-consumer with bounded queue

**Don't use when:**
- You need unlimited capacity (use linked list queue)
- Data loss is unacceptable (this overwrites when full)
- You need random access to all elements

## Code Example
```go
type CircularBuffer struct {
    Buffer     []int
    Size       int
    Capacity   int
    WriteIndex int
    ReadIndex  int
}

func (cb *CircularBuffer) Enqueue(value int) {
    posWrite := cb.WriteIndex % cb.Capacity  // Wraparound
    
    cb.Buffer[posWrite] = value
    cb.WriteIndex++
    
    if cb.Size == cb.Capacity {
        cb.ReadIndex++  // Buffer full - overwrite oldest, advance read
    } else {
        cb.Size++
    }
}

func (cb *CircularBuffer) Dequeue() int {
    if cb.Size == 0 {
        panic("Can't dequeue from empty queue")
    }
    
    posRead := cb.ReadIndex % cb.Capacity  // Wraparound
    value := cb.Buffer[posRead]
    
    cb.ReadIndex++
    cb.Size--
    return value
}

func (cb *CircularBuffer) IsEmpty() bool {
    return cb.Size == 0
}

func (cb *CircularBuffer) IsFull() bool {
    return cb.Size == cb.Capacity
}
```

## Trade-offs
- **Pro:** Fixed memory, O(1) operations, better cache locality than linked list, no allocations during runtime
- **Con:** Fixed capacity, overwrites data when full (data loss), slightly complex logic
- **Alternative:** Linked list queue for unlimited capacity, no data loss

**Circular vs Non-Circular Array:**

Non-circular problem:
```
Enqueue 1,2,3,4,5 -> WriteIndex = 5
Dequeue 3 times   -> ReadIndex = 3

Array: [X, X, X, 4, 5, _, _, ...]
                  ↑read    ↑write

Problem: Indices 0-2 wasted forever!
WriteIndex keeps growing: 6,7,8,9...
Eventually run out of array space even though 0-2 are free.
```

Circular solution:
```
WriteIndex = 25, Capacity = 8
25 % 8 = 1  -> wraps back to index 1

Reuses space! Indices wrap: 0,1,2,3,4,5,6,7,0,1,2...
```

**Why maintain explicit Size?**
Without Size, can't distinguish empty from full:
- Empty: ReadIndex == WriteIndex
- Full: ReadIndex == WriteIndex (after wraparound!)

With Size, it's unambiguous:
- Empty: Size == 0
- Full: Size == Capacity

## Overwrite Behavior

**What happens when full:**
```
Capacity = 4, Buffer = [1, 2, 3, 4]
ReadIndex = 0, WriteIndex = 4, Size = 4

Enqueue(5):
posWrite = 4 % 4 = 0  -> overwrites position 0 (data 1)
ReadIndex++ -> now 1 (skip the overwritten data)
Buffer = [5, 2, 3, 4]
         ↑overwritten
```

**When overwriting position 1:**
```
WriteIndex = 25
25 % 8 = 1

Last wrote to position 1 when WriteIndex was:
- 1 (first time)
- 9 (wrapped once: 9%8=1)  
- 17 (wrapped twice: 17%8=1)
- 25 (now, wrapped 3 times)

Overwriting data from WriteIndex=17 (8 enqueues ago)
```

## Interview Phrases
- "I'll implement a circular buffer using modulo for wraparound to reuse array space efficiently"
- "This achieves O(1) enqueue/dequeue with fixed memory - ideal for embedded or real-time systems"
- "The trade-off is bounded capacity and overwrite semantics for predictable memory usage"
- "The modulo operation maps unbounded indices to bounded array positions: WriteIndex=25, Capacity=8 → position 1"

## Can I Answer These?
- [x] Why circular? (Reuse space - otherwise indices grow unbounded and waste early positions)
- [x] Why modulo? (Wraps indices back to valid array range [0, Capacity-1])
- [x] What happens when full? (Overwrites oldest unread data, advances ReadIndex)
- [x] Why track Size? (Distinguish empty from full - both have ReadIndex==WriteIndex)
- [x] When would I NOT use this? (Need unlimited capacity or can't tolerate data loss)

## Design Choices

**Overwrite vs Block:**
- **My implementation:** Overwrites oldest data when full
- **Alternative:** Block/error on enqueue when full

When to overwrite:
- Real-time data (latest sensor readings)
- Audio/video (drop old frames)
- Logging (recent logs matter)

When to block:
- No data loss acceptable
- Need backpressure to slow producer

**Indices grow unbounded:**
```
ReadIndex and WriteIndex keep incrementing (0,1,2,3...forever)
Rely on modulo to map to array positions
```

**Potential issue:** After 2^64 operations, indices overflow (extremely unlikely in practice)

**Solution if needed:** Reset both indices when buffer becomes empty, or use wraparound arithmetic for size calculation: `Size = (WriteIndex - ReadIndex) % Capacity` (handles overflow naturally)

## Common Interview Follow-ups

**Q: How to avoid overwriting - make it block instead?**
A: On IsFull(), return false from Enqueue (or panic). Don't advance ReadIndex. Producer must handle full buffer.

**Q: What if ReadIndex/WriteIndex overflow?**
A: Practically never happens (2^64 operations). If concerned, reset both to 0 when buffer is empty.

**Q: How to implement with unbounded capacity?**
A: Use linked list queue instead - no fixed capacity, no overwriting. See [[Queue-LinkedList]].

**Q: How is this different from a ring buffer?**
A: Same concept - "ring buffer" and "circular buffer" are synonyms.

## Related Concepts
- [[Queue-LinkedList]] - unlimited capacity alternative
- [[Arrays-DynamicArray]] - similar modulo-based wraparound for capacity management
- [[Stack-LIFO]] - LIFO vs Queue's FIFO
