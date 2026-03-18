# Go — Arrays
 
---

## What is an Array?
A **fixed-size** collection of elements of the same type. Size is part of the type!

```go
arr := [3]int{10, 20, 30}   // array of 3 ints
arr := [5]int{1, 2, 3, 4, 5} // different type from [3]int!
```

> 💡 `[3]int` and `[4]int` are **completely different types** in Go — you can't even compare them!

---

## Declaration
```go
// With values
arr := [3]int{1, 2, 3}

// With zero values
var arr [3]int        // [0, 0, 0]

// Let Go count the size
arr := [...]int{1, 2, 3}  // size inferred as 3
```

---

## Key Properties

### Fixed Size
```go
arr := [3]int{1, 2, 3}
arr = append(arr, 4)   // 💥 compile error! append only works on slices
```

### Arrays are Copied on Assignment
```go
a := [3]int{1, 2, 3}
b := a          // full COPY — b is independent!
b[0] = 99

fmt.Println(a[0])  // 1 — unchanged!
fmt.Println(b[0])  // 99
```

> 💡 Unlike slices, arrays do **NOT** share memory on assignment!

### Arrays are Passed by Value to Functions
```go
func modify(a [3]int) {
    a[0] = 99    // modifies the COPY, not original!
}

func main() {
    a := [3]int{1, 2, 3}
    modify(a)
    fmt.Println(a[0])   // 1 — original unchanged!
}
```

> ⚠️ To modify original, pass a pointer:
> ```go
> func modify(a *[3]int) {
>     a[0] = 99    // modifies the ORIGINAL!
> }
> modify(&a)
> ```

---

## Slicing an Array
You can slice an array — returns a **slice** backed by the original array:
```go
a := [5]int{1, 2, 3, 4, 5}
s := a[1:3]     // slice [2, 3]

s[0] = 99
fmt.Println(a[1])  // 99 — modifying slice affects original array!
```

> ⚠️ Shared memory rule applies — slice and array share the same memory!

---

## Comparing Arrays
Arrays can be compared with `==` only if they are the **same type:**
```go
a := [3]int{1, 2, 3}
b := [3]int{1, 2, 3}
c := [4]int{1, 2, 3, 4}

fmt.Println(a == b)   // true  — same type, same values
fmt.Println(a == c)   // 💥 compile error — different types!
```

> ⚠️ Slices **cannot** be compared with `==` (except against `nil`)!

---

## Array vs Slice — Key Differences

| | Array | Slice |
|--|-------|-------|
| Size | Fixed | Dynamic |
| Assignment | Full copy | Shares memory |
| `append` | ❌ not allowed | ✅ allowed |
| Comparison `==` | ✅ same type only | ❌ not allowed |
| Pass to function | Copied | Reference to underlying array |
| Usage | Rare | 90% of the time |

---

## Quick Reference
```go
arr := [3]int{1, 2, 3}    // declaration
arr := [...]int{1, 2, 3}  // inferred size
len(arr)                   // length
arr[0]                     // access element
arr[1:3]                   // slice from array → [2, 3]
a == b                     // compare (same type only)
modify(&arr)               // pass by pointer to modify original
```

---