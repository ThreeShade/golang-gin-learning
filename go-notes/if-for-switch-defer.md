# Go — For, If/Else, Switch, Defer

---

## For Loop


### Basic For Loop
```go
// No parentheses () in Go!
for i := 1; i < 6; i++ {
    fmt.Println(i)
}
```

### For as While Loop
```go
x := 0
for x < 10 {
    x++
}
```

### Infinite Loop
```go
for {
    // runs forever — like while(true)
}
```

---

## If / Else

### Basic Syntax
```go
// No parentheses () around condition!
if x > 0 {
    fmt.Println("positive")
} else if x < 0 {
    fmt.Println("negative")
} else {
    fmt.Println("zero")
}
```

### If Initialiser Statement
Declare a variable directly inside the `if` — scoped to the block only!
```go
if x := 10; x > 5 {
    fmt.Println(x)   // ✅ accessible here
}
fmt.Println(x)       // 💥 compile error — x is out of scope!
```

> 💡 Common real world use — error checking:
> ```go
> if err := doSomething(); err != nil {
>     fmt.Println("error:", err)
> }
> ```

---

## Switch

### Basic Switch
```go
switch x {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
default:
    fmt.Println("other")
}
```

> 💡 Go cases **break automatically** — no need to write `break`!

### Fallthrough
Use `fallthrough` to explicitly execute the next case:
```go
switch x {
case 2:
    fmt.Println("two")
    fallthrough   // continues to next case!
case 3:
    fmt.Println("three")
}
```

### Expressionless Switch
No variable after `switch` — evaluates each case as a boolean, executes **first true:**
```go
switch {
case x > 100:
    fmt.Println("big")
case x > 50:
    fmt.Println("medium")
case x > 10:
    fmt.Println("small")
default:
    fmt.Println("tiny")
}
```

> 💡 Cleaner alternative to long `if/else if` chains!

---

## Defer

### Basic Defer
Pushes a function call onto a stack — executes when the **surrounding function returns:**
```go
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
// Output:
// hello
// world
```

> 💡 Common use — guaranteed resource cleanup:
> ```go
> file, _ := os.Open("file.txt")
> defer file.Close()   // always runs when function returns!
> ```

### Multiple Defers — LIFO
Multiple defers execute in **Last In First Out** order:
```go
func main() {
    defer fmt.Println("one")
    defer fmt.Println("two")
    defer fmt.Println("three")
    fmt.Println("start")
}
// Output:
// start
// three
// two
// one
```

### Defer in Loops
```go
for i := 0; i < 3; i++ {
    defer fmt.Println(i)
}
// Output:
// 2
// 1
// 0
```

> ⚠️ Be careful with defer inside large loops — all defers are held in memory until the function returns!

---

## Quick Reference
```go
// For loop
for i := 0; i < 10; i++ { }   // basic
for x < 10 { }                 // while style
for { }                         // infinite

// If/Else
if x > 0 { } else if { } else { }
if x := 10; x > 5 { }          // initialiser statement

// Switch
switch x { case 1: case 2: default: }   // auto break
switch { case x > 10: }                  // expressionless

// Defer
defer fmt.Println("last")       // runs when function returns
// Multiple defers = LIFO order
```

---
*Notes compiled from interactive tutoring session — keep learning! 🚀*