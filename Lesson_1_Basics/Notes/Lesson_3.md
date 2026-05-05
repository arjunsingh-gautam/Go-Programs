# **<span style="color:#ff1744">Go Basic Boilerplate — Complete Mental Model</span>**

---

# **<span style="color:#ff6f00">1. Minimal Go Program (Boilerplate)</span>**

---

## **<span style="color:#3a86ff">Example</span>**

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go")
}
```

---

## **<span style="color:#8338ec">What This Contains</span>**

```text
1. package main  → defines program type
2. import        → brings external functionality
3. main()        → entry point of execution
```

---

# **<span style="color:#ff1744">2. package main — What & Why</span>**

---

## **<span style="color:#3a86ff">Definition</span>**

```text
package main tells Go:
"This is an executable program"
```

---

## **<span style="color:#3a86ff">Why It Is Needed</span>**

```text
Only package main can produce a runnable binary
```

---

## **<span style="color:#3a86ff">Execution Rule</span>**

```text
Go looks for:
package main + func main()
```

If both exist:

```text
→ Program runs
```

Otherwise:

```text
→ Compilation error
```

---

## **<span style="color:#8338ec">Analogy</span>**

```text
package main = "This is the app"
main() = "Start execution here"
```

---

# **<span style="color:#ff1744">3. Can You Run Go Without package main?</span>**

---

## **<span style="color:#3a86ff">No (for executables)</span>**

```text
Without package main:
❌ Cannot run using `go run`
❌ Cannot build executable
```

---

## **<span style="color:#3a86ff">But You Can Write Libraries</span>**

```go
package mathutils
```

✔ Used as:

```go
import "github.com/arj/mathutils"
```

---

## **<span style="color:#8338ec">Key Idea</span>**

```text
main package = runnable program
other packages = reusable libraries
```

---

# **<span style="color:#ff1744">4. import — Why It Exists</span>**

---

## **<span style="color:#3a86ff">Definition</span>**

```text
import brings external packages into your file
```

---

## **<span style="color:#3a86ff">Example</span>**

```go
import "fmt"
```

→ gives access to:

```go
fmt.Println()
```

---

## **<span style="color:#3a86ff">Without import</span>**

```text
❌ You cannot use functions from other packages
```

---

## **<span style="color:#8338ec">Important Insight</span>**

```text
Go does NOT allow implicit imports
Everything must be explicit
```

---

# **<span style="color:#ff1744">5. Can You Run Go Without import?</span>**

---

## **<span style="color:#3a86ff">Yes — if you don’t need external functionality</span>**

```go
package main

func main() {
}
```

✔ This runs

---

## **<span style="color:#3a86ff">But</span>**

```text
No print, no IO, no libraries
→ very limited
```

---

# **<span style="color:#ff1744">6. main() Function — Entry Point</span>**

---

## **<span style="color:#3a86ff">Definition</span>**

```text
func main() is where execution starts
```

---

## **<span style="color:#3a86ff">Rules</span>**

```text
- Must be in package main
- Must have NO parameters
- Must return NOTHING
```

---

## **<span style="color:#3a86ff">Execution Flow</span>**

```text
1. Program starts
2. Go runtime initializes
3. Calls main()
4. Executes line by line
```

---

## **<span style="color:#8338ec">Example</span>**

```go
func main() {
    println("Step 1")
    println("Step 2")
}
```

Execution:

```text
Step 1 → Step 2
```

---

# **<span style="color:#ff1744">7. Why Capital Letters Matter (VERY IMPORTANT)</span>**

---

## **<span style="color:#3a86ff">Core Rule</span>**

```text
Capitalized = EXPORTED (public)
Lowercase   = PRIVATE (package-only)
```

---

## **<span style="color:#3a86ff">Example</span>**

### **Package code**

```go
package utils

func Add() {}     // Exported
func subtract() {} // Private
```

---

### **Usage in another file**

```go
import "utils"

utils.Add()       // ✔ Works
utils.subtract()  // ❌ Error
```

---

## **<span style="color:#8338ec">Why This Exists</span>**

```text
Go does NOT use keywords like public/private

Instead:
→ Naming convention controls visibility
```

---

## **<span style="color:#3a86ff">Applies To</span>**

```text
✔ Functions
✔ Structs
✔ Variables
✔ Fields inside structs
✔ Interfaces
```

---

## **<span style="color:#8338ec">Example (Struct)</span>**

```go
type User struct {
    Name string // exported
    age  int    // private
}
```

---

# **<span style="color:#ff1744">8. Internal Mechanics (How Everything Works Together)</span>**

---

## **<span style="color:#8338ec">Compilation Flow</span>**

```text
1. Go reads files
2. Groups them by package
3. Finds package main
4. Looks for func main()
5. Resolves imports
6. Builds dependency graph
7. Compiles to binary
8. Executes main()
```

---

## **<span style="color:#3a86ff">Import Resolution</span>**

```text
import "fmt"
↓
Go finds it in standard library
↓
Links it during compilation
```

---

## **<span style="color:#3a86ff">Visibility Resolution</span>**

```text
Access symbol:
↓
Check first letter:
    Uppercase → allowed
    Lowercase → restricted
```

---

# **<span style="color:#ff1744">9. Common Beginner Mistakes</span>**

---

## **<span style="color:#3a86ff">Mistake 1</span>**

```text
Using lowercase function across packages
→ causes undefined error
```

---

## **<span style="color:#3a86ff">Mistake 2</span>**

```text
Missing package main
→ cannot run program
```

---

## **<span style="color:#3a86ff">Mistake 3</span>**

```text
Unused imports
→ Go compiler throws error
```

---

## **<span style="color:#8338ec">Why?</span>**

```text
Go enforces clean, minimal code
```

---

# **<span style="color:#ff1744">10. Simple Exercise (Do This Yourself)</span>**

---

## **<span style="color:#3a86ff">Goal</span>**

```text
Understand package, import, and export rules
```

---

## **<span style="color:#3a86ff">Steps</span>**

---

### **1. Create project**

```bash
mkdir go-practice
cd go-practice
go mod init go-practice
```

---

### **2. Create utils package**

```text
go-practice/
 ├── main.go
 └── utils/
      └── math.go
```

---

### **3. Write utils/math.go**

```go
package utils

func Add(a int, b int) int {
    return a + b
}

func subtract(a int, b int) int {
    return a - b
}
```

---

### **4. Write main.go**

```go
package main

import (
    "fmt"
    "go-practice/utils"
)

func main() {
    fmt.Println(utils.Add(2, 3))
    // fmt.Println(utils.subtract(5, 2)) // try uncommenting
}
```

---

### **5. Run**

```bash
go run main.go
```

---

## **<span style="color:#8338ec">Observe</span>**

```text
✔ Add works
❌ subtract fails (not exported)
```

---

# **<span style="color:#ff1744">11. Key Insight (CRITICAL)</span>**

```text
Go is simple by design:

package → structure
main() → entry point
import → explicit dependencies
Capital letters → visibility control
```

---

# **<span style="color:#ff1744">Final Summary</span>**

---

### **package main**

```text
Marks program as executable
```

---

### **main()**

```text
Starting point of execution
```

---

### **import**

```text
Brings external code into scope
```

---

### **Capitalization**

```text
Controls visibility (public/private)
```

---

### **Core Philosophy**

```text
Go prefers simplicity, explicitness, and strict rules
```

---

If you want next deep topics:

```text
- "How Go packages work internally"
- "Build system and compilation in Go"
- "Interfaces and struct mechanics in depth"
```
