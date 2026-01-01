# Hash Tables (Separate Chaining)

## What I Implemented
A hash table using separate chaining (linked lists at each bucket) to handle collisions. Fixed-size array of linked lists with a simple additive hash function.

## How It Works
I hash the key to get a bucket index (0 to size-1). If the bucket is empty, I create a new linked list. Otherwise, I add/update the key-value pair in the existing list at that bucket. On lookup, I hash the key and search the linked list at that bucket.

**Hash function:** Sum ASCII values of all characters, then modulo to fit in array range.

**Collision handling:** Separate chaining - all keys that hash to the same bucket go into a linked list at that bucket.

## Complexity
**Time:**
- Insert: O(1) average, O(n) worst case (all keys hash to same bucket)
- Delete: O(1) average, O(n) worst case
- Lookup: O(1) average, O(n) worst case

**Space:** O(n + m) where n = number of elements, m = table size (buckets)

**Average case assumes:**
- Good hash function (uniform distribution)
- Load factor α = n/m stays reasonable (< 0.75)

## When to Use
Use hash tables when you need:
- **Fast lookups** by key (O(1) average)
- **Key-value associations** (dictionaries, caches)
- **Set operations** (membership testing, deduplication)
- **Frequency counting** (count occurrences)

**Common use cases:**
- Databases (indexing)
- Caches (memoization)
- Symbol tables (compilers)
- Two-sum problem (complement lookup)

**Don't use when:**
- You need ordering (use BST or sorted array)
- You need range queries (use tree)
- Keys aren't hashable

## Code Example
```go
type HashTable struct {
    Keys []*LinkedList
    Size int
}

func (ht *HashTable) Hash(key string) int {
    total := 0
    for _, char := range key {
        total += int(char)
    }
    return total % ht.Size
}

func (ht *HashTable) Insert(key string, value int) {
    hashedKey := ht.Hash(key)
    
    if ht.Keys[hashedKey] == nil {
        newNode := NewNode(key, value)
        ht.Keys[hashedKey] = NewList(newNode)
    } else {
        ht.Keys[hashedKey].AddOrUpdate(key, value)
    }
}

func (ht *HashTable) ValueAt(key string) int {
    hashedKey := ht.Hash(key)
    return ht.Keys[hashedKey].ValueAt(key)  // O(list length)
}
```

## Hash Function Analysis

**My simple hash (additive):**
```go
hash = (a + b + c) % size
```

**Problem:** Anagrams collide!
- "abc" = 97+98+99 = 294
- "bca" = 98+99+97 = 294
- Same hash!

**Better approach (polynomial rolling hash):**
```go
hash = 0
prime := 31
for _, char := range key {
    hash = hash*prime + int(char)
}
return hash % size

// "abc" = (a*31² + b*31 + c) - position matters!
// "bca" = (b*31² + c*31 + a) - different hash!
```

**Why this is better:**
- Position matters (not commutative)
- Prime number (31) improves distribution
- Horner's method (efficient - no exponentiation)

**What makes a good hash function:**
1. Uniform distribution (spread keys evenly)
2. Fast to compute
3. Deterministic (same key → same hash)
4. Minimize collisions

## Trade-offs: Separate Chaining vs Linear Probing

**Separate Chaining (my implementation):**
```
Bucket 3: dog → cat → ant
Bucket 5: rat
Bucket 6: bat
```

**Linear Probing:**
```
[_][_][_][dog][cat][ant][rat][bat]
         3    4    5    6    7
```

| Aspect | Separate Chaining | Linear Probing |
|--------|-------------------|----------------|
| Collision handling | Linked lists | Probe next buckets |
| Cache performance | ❌ Worse (scattered nodes) | ✅ Better (contiguous) |
| Memory overhead | ❌ Pointers per node | ✅ Just array |
| Clustering | ✅ Local to bucket | ❌ Spreads across buckets |
| Performance at high load | ✅ Degrades gracefully | ❌ Severe degradation |
| Deletion | ✅ Simple (remove node) | ❌ Complex (tombstones) |

**When to choose which:**
- **Low load factor (<50%):** Linear probing (cache wins)
- **High load factor (>70%):** Separate chaining (clustering hurts linear probing)
- **Unpredictable data:** Separate chaining (safer, more predictable)
- **Embedded systems:** Linear probing (less memory)

**Clustering example (linear probing):**
```
Insert "dog" (hash=3): [_][_][_][D][_][_][_][_]
Insert "cat" (hash=3): [_][_][_][D][C][_][_][_]
Insert "ant" (hash=3): [_][_][_][D][C][A][_][_]
Insert "rat" (hash=5): [_][_][_][D][C][A][R][_]  <- "rat" hashes to 5, but 5 is taken!

Now cluster 3-6 is full. Next insertion to ANY of these buckets probes multiple slots.
```

With separate chaining, "rat" at bucket 5 is independent - no cluster pollution.

## Load Factor and Resizing

**Load factor α = n/m** (elements / buckets)

**Performance by load factor:**
- α = 0.5: Both methods perform well
- α = 0.75: Separate chaining still O(1), linear probing starts degrading
- α = 0.9: Separate chaining O(1-2), linear probing can be O(10+)
- α > 1.0: Separate chaining still works (lists grow), linear probing fails

**When to resize:**
Typically resize when α > 0.75:
1. Allocate new array (double size)
2. Rehash all keys (hash values change due to new modulo)
3. Insert into new table

**Amortized O(1) like dynamic array.**

## Interview Phrases
- "I'll implement a hash table with separate chaining for predictable O(1) average-case performance"
- "This achieves fast lookups by hashing the key to a bucket, then searching the linked list at that bucket"
- "The trade-off is pointer overhead and cache performance for graceful degradation at high load factors"
- "My simple additive hash has collision issues with anagrams - a better approach is polynomial rolling hash where position matters"
- "Separate chaining handles high load factors better than linear probing because clusters are isolated to individual buckets"

## Can I Answer These?
- [x] Why is it O(1) average? (Good hash distributes keys evenly, short lists)
- [x] Why O(n) worst case? (All keys hash to same bucket = one long list)
- [x] Why does my hash function cause anagram collisions? (Addition is commutative, position doesn't matter)
- [x] How to improve hash function? (Polynomial rolling hash: hash = hash*31 + char)
- [x] Separate chaining vs linear probing? (Chaining better at high load, probing better for cache)
- [x] What is load factor? (n/m - ratio of elements to buckets)
- [x] When to resize? (When α > 0.75 typically)

## Common Interview Follow-ups

**Q: How would you handle collisions without linked lists?**
A: Linear probing (check next bucket), quadratic probing (check i² buckets away), or double hashing.

**Q: What if you need to maintain insertion order?**
A: Use LinkedHashMap - hash table + doubly-linked list connecting all entries in insertion order.

**Q: Implement LRU cache using hash table.**
A: Hash table for O(1) lookup + doubly-linked list for O(1) move-to-front and eviction.

**Q: Two-sum problem: find two numbers that sum to target.**
A: Hash table! For each num, check if (target - num) exists in hash. O(n) time, O(n) space.

## Related Concepts
- [[Linked Lists]] - used for separate chaining
- [[Dynamic Array]] - similar resizing strategy
- [[Binary Search Tree]] - alternative for ordered data
- [[Two Pointers]] - alternative for two-sum on sorted array
