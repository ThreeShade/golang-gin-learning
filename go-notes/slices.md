# Go — Slices

---

## What is a Slice?
A slice does **not store data** — it just describes a section of an underlying array.

Every slice has 3 internal components:
```
┌─────────────┐
│ pointer     │ → points to underlying array
│ length      │ → number of actual elements
│ capacity    │ → total allocated space
└─────────────┘
```

---

## Declaration
```go
// Slice literal
s := []int{1, 2, 3}

// Using make(type, length, capacity)
s := make([]int, 3)      // [0, 0, 0] len=3, cap=3
s := make([]int, 3, 6)   // [0, 0, 0] len=3, cap=6

// Nil slice
var s []int              // nil, len=0, cap=0
fmt.Println(s == nil)    // true
```

> 💡 Nil slice is safe — `len()`, `cap()`, `append()` all work on nil slices!

---

## len and cap
```go
s := []int{1, 2, 3, 4, 5}
fmt.Println(len(s))   // 5 — number of elements
fmt.Println(cap(s))   // 5 — allocated space
```

### Cap after slicing
Cap is counted from the **starting index to the end of the original array:**
```go
s := []int{1, 2, 3, 4, 5}   // cap=5
s2 := s[1:4]                   // starts at index 1

fmt.Println(s2)        // [2 3 4]
fmt.Println(len(s2))   // 3
fmt.Println(cap(s2))   // 4 (= 5 - 1, from index 1 to end)
```

> 💡 Formula: `cap = original cap - starting index`

---

## Slice Syntax `[low:high]`
- `low` is **inclusive**, `high` is **exclusive**

```go
s := []int{1, 2, 3, 4, 5}

s[1:3]  // [2, 3]  → index 1 up to (not including) 3
s[:2]   // [1, 2]  → from start up to (not including) 2
s[3:]   // [4, 5]  → from index 3 to end
```

---

## Shared Memory
Slices share memory with their underlying array — modifying one affects the other!

```go
arr := [5]int{1, 2, 3, 4, 5}
s1 := arr[0:3]   // [1, 2, 3]
s2 := arr[1:4]   // [2, 3, 4]

s1[1] = 99
fmt.Println(arr[1])  // 99 — arr affected!
fmt.Println(s2[0])   // 99 — s2 affected too!
```

### Assignment shares memory too
```go
s1 := []int{1, 2, 3}
s2 := s1         // shares same underlying array!
s2[0] = 99
fmt.Println(s1[0])  // 99 — s1 affected!
```

---

## append
Adds elements to the end of a slice and **returns a new slice:**

```go
s = append(s, 4)           // add one element
s = append(s, 4, 5, 6)    // add multiple
s = append(s, s2...)       // append another slice
```

> ⚠️ Always assign result back — `append(s, 4)` alone loses the result!

### How append works internally

**Case 1 — Enough capacity:**
```
s = [1, 2, 3]  len=3, cap=5
append(s, 4)
s = [1, 2, 3, 4]  len=4, cap=5  ← same array!
```

**Case 2 — No capacity:**
```
s = [1, 2, 3]  len=3, cap=3  ← full!
append(s, 4)
NEW array allocated!
s = [1, 2, 3, 4]  len=4, cap=6  ← doubled cap, new array!
```

### append breaks shared memory
```go
s1 := []int{1, 2, 3}   // cap=3, full!
s2 := s1
s2 = append(s2, 4)     // cap exceeded → NEW array for s2
s2[0] = 99

fmt.Println(s1)   // [1 2 3]    — s1 unaffected!
fmt.Println(s2)   // [99 2 3 4] — s2 has its own array!
```

> 💡 append adds AFTER existing elements — `make([]int, 3, 6)` creates 3 zeros first!
> ```go
> s := make([]int, 3, 6)   // [0, 0, 0]
> s = append(s, 1)          // [0, 0, 0, 1]  len=4
> ```

---

## copy — Deep Copy
Creates a **truly independent** slice — no shared memory:

```go
s1 := []int{1, 2, 3}
s2 := make([]int, len(s1))
copy(s2, s1)        // copies VALUES not memory address

s2[0] = 99
fmt.Println(s1)   // [1 2 3]   — unchanged!
fmt.Println(s2)   // [99 2 3]  — only s2 changed
```

> 💡 Difference:
> ```go
> s2 := s1       // ⚠️ shares memory — linked!
> copy(s2, s1)   // ✅ deep copy — independent!
> ```

---

## Delete an Element
```go
s := []int{1, 2, 3, 4, 5}

// Delete index 2 (value 3)
s = append(s[:2], s[3:]...)

fmt.Println(s)   // [1 2 4 5]
```

---

## Quick Reference
```go
s := []int{1, 2, 3}          // literal
s := make([]int, 3, 6)        // len=3, cap=6
len(s)                         // number of elements
cap(s)                         // allocated capacity
s[1:3]                         // [2, 3]
s = append(s, 4)               // add element
s = append(s, s2...)           // append slice
copy(dst, src)                  // deep copy
s = append(s[:i], s[i+1:]...)  // delete index i
```

---