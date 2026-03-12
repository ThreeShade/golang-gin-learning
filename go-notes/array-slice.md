##  Arrays and Slices

### Arrays 
- **Fixed size** — size is part of the type
- `[3]int` and `[5]int` are **different types**
- Rarely used in practice

```go
arr := [3]int{10, 20, 30}
fmt.Println(len(arr)) // 3
```

> ⚠️ `append` does **NOT** work on arrays — compile error!

---

### Slices
- **Dynamic size** — can grow and shrink
- Used **90% of the time** in real Go code
- Backed by an underlying array

#### Two ways to declare a slice:
```go
s := []int{1, 2, 3}         // slice literal
s := make([]int, 3)          // using make — creates [0, 0, 0]
s := make([]int, 3, 5)       // make(type, length, capacity)
```

#### Nil slice
```go
var s []int
fmt.Println(s == nil)  // true
fmt.Println(len(s))    // 0  — safe to use!
```

> 💡 A nil slice is safe — `len()`, `cap()`, and `append()` all work on nil slices!

---

### `make([]int, length, capacity)`
| Argument | Meaning |
|----------|---------|
| `length` | Number of elements (zero-valued) |
| `capacity` | Allocated memory before reallocation |

```go
s := make([]int, 3, 5)
fmt.Println(len(s)) // 3
fmt.Println(cap(s)) // 5
fmt.Println(s)      // [0 0 0]
```

---

### Slice Syntax `[low:high]`
- `low` is **inclusive**, `high` is **exclusive**

```go
s := []int{1, 2, 3, 4, 5}

s[1:3]  // [2, 3]  → index 1 up to (not including) 3
s[:2]   // [1, 2]  → from start up to (not including) 2
s[3:]   // [4, 5]  → from index 3 to end
```

---

### Slices Share Memory with Arrays
A slice is a **window** into an underlying array — not a copy!

```go
arr := [3]int{10, 20, 30}
s := arr[0:2]    // s points to the same memory as arr

s[0] = 99
fmt.Println(arr[0]) // 99 — arr is also changed!
fmt.Println(s[0])   // 99
```

> ⚠️ Modifying a slice element **also modifies the original array!**

---

### Slice Copy Shares Memory Too

```go
s1 := []int{1, 2, 3}
s2 := s1         // copies the slice header, NOT the data
s2[0] = 99

fmt.Println(s1[0]) // 99 — s1 is also affected!
fmt.Println(s2[0]) // 99
```

> ⚠️ `s2 := s1` does NOT deep copy — both share the same underlying array!

#### But `append` breaks the shared memory link!

```go
s1 := []int{1, 2, 3}
s2 := s1
s2 = append(s2, 4)  // cap exceeded → Go allocates NEW array for s2
s2[0] = 99

fmt.Println(s1[0])  // 1  — s1 unaffected! s2 now has its own array
fmt.Println(s2[0])  // 99
```

> 💡 When `append` exceeds capacity, Go allocates a **new underlying array** — the two slices no longer share memory!

---

### `len` vs `cap`
| | Meaning |
|--|---------|
| `len(s)` | Number of **actual elements** in the slice |
| `cap(s)` | Total **allocated memory** (before Go needs to reallocate) |

```go
s := []int{1, 2, 3}
fmt.Println(len(s), cap(s)) // 3 3

s = append(s, 4)
s = append(s, 5)
fmt.Println(len(s), cap(s)) // 5 6
```

> 💡 When a slice runs out of capacity, Go **doubles the capacity** automatically to avoid constant memory reallocation. This is an implementation detail and may vary.

---

## Quick Reference Cheatsheet

```go
// --- POINTERS ---
x := 42
p := &x          // get address
*p = 100         // dereference and set
var np *int      // nil pointer
if np != nil { } // always check before dereferencing

// --- ARRAYS ---
arr := [3]int{1, 2, 3}   // fixed size

// --- SLICES ---
s := []int{1, 2, 3}         // slice literal
s := make([]int, 3)          // [0, 0, 0]
s := make([]int, 3, 5)       // length=3, capacity=5
s = append(s, 4)             // add element
s[1:3]                       // [2, 3] — low inclusive, high exclusive
s[:2]                        // [1, 2]
s[3:]                        // from index 3 to end
len(s)                       // number of elements
cap(s)                       // allocated capacity
// ⚠️ s2 := s1 shares memory! use copy() for a true deep copy
copy(dst, src)               // deep copy a slice
```

---