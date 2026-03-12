## 3. Packages
 
### What is a Package?
A way to **group related code** together. Every Go file must belong to a package.
 
### The `main` Package
Every executable Go program must have:
```go
package main
 
func main() {
    // program starts here
}
```
 
### Importing Packages
```go
import "fmt"                 // single import
 
import (                     // grouped import (idiomatic)
    "fmt"
    "math"
    "os"
)
```
 
> 💥 Importing a package and **not using it = compile error!** (stricter than most languages)
 
### Import Tricks
```go
import (
    f "fmt"          // alias — use as f.Println()
    _ "github.com/lib/pq"  // blank identifier — import for side effects only
)
```
 
### Exported vs Unexported
Go uses **capitalisation** to control visibility — no `public`/`private` keywords!
 
| | Rule | Accessible |
|--|------|-----------|
| `SayHello()` | Capital letter | ✅ Outside package |
| `saySecret()` | Lowercase letter | ❌ Inside package only |
 
Applies to **everything** — functions, variables, types, structs!
 
### `init()` Function
A special reserved function that runs **automatically before `main()`**:
```go
func init() {
    // setup code — DB connections, config loading etc.
    // runs before main(), cannot be called manually
}
 
func main() {
    // runs after init()
}
```
 
> 💡 A file can have **multiple `init()` functions** — they all run in order!
 
---
 
// --- PACKAGES ---
package main                 // entry point package
import "fmt"                 // single import
import (                     // grouped import
    "fmt"
    f "math"                 // alias
    _ "some/pkg"             // side effects only
)
func init() {}               // runs before main(), auto-called
func Exported() {}           // capital = public
func unexported() {}         // lowercase = private
 
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