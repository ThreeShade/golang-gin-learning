# Go — Variables & Functions

---

## Variables

### 3 Ways to Declare
```go
// Way 1 — declare then assign
var number int
number = 10

// Way 2 — declare with value
var number int = 10

// Way 3 — shorthand (most common, inside functions only)
number := 10
```

### Type Inference vs Dynamic Typing
`:=` uses **type inference** — Go infers the type at compile time. It is still statically typed!
```go
number := 10       // Go infers int — forever!
number = "hello"   // 💥 compile error — can't change type!
```

### Scope of `:=`
```go
var globalVar int = 10   // ✅ package level
globalShort := 10        // ❌ := only works inside functions!
```

### Zero Values
Every variable is initialised to its zero value automatically:

| Type | Zero Value |
|------|-----------|
| `int`, `float64` | `0` |
| `string` | `""` |
| `bool` | `false` |
| `pointer`, `slice` | `nil` |

### Basic Types
```go
// Integers
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64

// Floats
float32, float64

// Others
string, bool
byte    // alias for uint8
rune    // alias for int32 (unicode character)
```

### `:=` Rules
```go
x := 10
x = 20        // ✅ reassign with =
x := 30       // ❌ already declared!

// ✅ OK if at least one variable is NEW
x, y := 30, 40   // y is new — valid!
```

### Variable Shadowing
```go
var x int = 10       // package level

func main() {
    fmt.Println(x)   // 10 — package level
    x := 50
    fmt.Println(x)   // 50 — local x shadows package x
    {
        x := 100
        fmt.Println(x)   // 100 — block x shadows local x
    }
    fmt.Println(x)   // 50 — back to local x
}
```

> ⚠️ Shadowing is valid Go but can be confusing — use carefully!

---

## Functions

### Basic Syntax
```go
func functionName(params) returnType {
    // body
}

// Example
func add(a, b int) int {
    return a + b
}
```

### Multiple Return Values
Go functions can return **more than one value!**
```go
func minMax(a, b int) (int, int) {
    return a, b
}

min, max := minMax(3, 7)

// Ignore a return value with _
min, _ := minMax(3, 7)
```

> 💡 Most common use — return value + error:
> ```go
> func divide(a, b float64) (float64, error) {
>     if b == 0 {
>         return 0, errors.New("cannot divide by zero")
>     }
>     return a / b, nil
> }
> ```

### Named Returns (Naked Return)
```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return       // automatically returns x and y
}
```

> ⚠️ Use naked returns only in short functions — hurts readability in long ones!

### Variadic Functions
Accept **any number of arguments** using `...`:
```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

sum(1, 2, 3)        // 6
sum(1, 2, 3, 4, 5)  // 15
```

> 💡 Inside the function, `nums` is a `[]int` slice!
> `fmt.Println()` is variadic too — that's why it accepts any number of args!

### First Class Functions
Functions can be assigned to variables and passed as arguments:
```go
// Function assigned to variable
double := func(n int) int {
    return n * 2
}
double(5)  // 10

// Function as parameter
func apply(nums []int, f func(int) int) []int {
    result := []int{}
    for _, n := range nums {
        result = append(result, f(n))
    }
    return result
}

apply([]int{1, 2, 3}, double)  // [2 4 6]
```

### Closures
A function that **remembers variables from its surrounding scope:**
```go
func counter() func() int {
    count := 0
    return func() int {
        count++        // remembers count between calls!
        return count
    }
}

c := counter()
c()  // 1
c()  // 2
c()  // 3

// Each closure has its OWN state!
c1 := counter()
c2 := counter()
c1()  // 1
c1()  // 2
c2()  // 1 — independent!
```

---

## Quick Reference
```go
// Variables
var x int           // zero value = 0
var x int = 10      // explicit
x := 10             // inferred (functions only)
x, y := 10, 20      // multiple declaration

// Functions
func add(a, b int) int { return a + b }           // basic
func minMax(a, b int) (int, int) { return a, b }  // multiple returns
func sum(nums ...int) int { }                      // variadic
double := func(n int) int { return n * 2 }        // first class
```

---
*Notes compiled from interactive tutoring session — keep learning! 🚀*