# **<span style="color:#ff1744">Variables in Go — Complete Mental Model</span>**

---

# **<span style="color:#ff6f00">1. What is a Variable in Go?</span>**

---

## **<span style="color:#3a86ff">Definition</span>**

```text id="v1def"
A variable is a named storage location that holds a value of a specific type
```

---

## **<span style="color:#3a86ff">Core Idea</span>**

```text id="v1core"
Variable = name → memory → typed value
```

---

## **<span style="color:#8338ec">Example</span>**

```go id="v1ex"
var age int = 21
```

---

## **<span style="color:#8338ec">Analogy</span>**

```text id="v1ana"
Variable = labeled box
Label = name
Box = memory
Content = value
```

---

# **<span style="color:#ff1744">2. Ways to Declare & Initialize Variables</span>**

---

# **<span style="color:#8338ec">A. Full Declaration (Explicit Type)</span>**

```go id="v2a"
var age int = 21
```

✔ Type specified
✔ Value assigned

---

# **<span style="color:#8338ec">B. Type Inference</span>**

```go id="v2b"
var age = 21
```

✔ Compiler infers type (`int`)

---

# **<span style="color:#8338ec">C. Short Declaration (Most Used)</span>**

```go id="v2c"
age := 21
```

✔ Only inside functions
✔ Automatically infers type

---

# **<span style="color:#8338ec">D. Multiple Variables</span>**

```go id="v2d"
var a, b int = 10, 20
```

---

## **<span style="color:#3a86ff">Mixed Initialization</span>**

```go id="v2d2"
var a, b = 10, "hello"
```

---

# **<span style="color:#8338ec">E. Block Declaration</span>**

```go id="v2e"
var (
    x int = 10
    y string = "Go"
)
```

---

# **<span style="color:#8338ec">F. Zero Values (IMPORTANT)</span>**

```go id="v2f"
var x int
var s string
```

Default values:

```text id="v2fout"
int → 0
string → ""
bool → false
pointer → nil
```

---

# **<span style="color:#ff1744">3. Rules of Variable Declaration</span>**

---

## **<span style="color:#3a86ff">Rule 1: Must Be Used</span>**

```text id="v3r1"
Unused variables → compilation error
```

---

## **<span style="color:#3a86ff">Rule 2: := Only Inside Functions</span>**

```go id="v3r2"
func main() {
    x := 10  // ✔
}
```

```go id="v3r2_bad"
x := 10  // ❌ outside function
```

---

## **<span style="color:#3a86ff">Rule 3: No Re-declaration in Same Scope</span>**

```go id="v3r3"
x := 10
x := 20  // ❌ error
```

---

## **<span style="color:#3a86ff">Rule 4: Type Safety</span>**

```go id="v3r4"
var x int = "hello" // ❌
```

---

# **<span style="color:#ff1744">4. Scoping Rules in Go</span>**

---

# **<span style="color:#8338ec">A. Global Scope</span>**

```go id="v4a"
var globalVar = 100
```

✔ Accessible everywhere in package

---

# **<span style="color:#8338ec">B. Function Scope</span>**

```go id="v4b"
func main() {
    x := 10
}
```

✔ Accessible only inside function

---

# **<span style="color:#8338ec">C. Block Scope</span>**

```go id="v4c"
if true {
    y := 20
}
```

✔ Exists only inside block

---

# **<span style="color:#8338ec">D. Shadowing (VERY IMPORTANT)</span>**

```go id="v4d"
x := 10

if true {
    x := 20  // new variable (shadow)
}
```

---

## **<span style="color:#3a86ff">Effect</span>**

```text id="v4dout"
Outer x remains unchanged
```

---

# **<span style="color:#ff1744">5. Internal Mechanics</span>**

---

## **<span style="color:#8338ec">Memory Allocation</span>**

```text id="v5mem"
Variables stored in:
- Stack (fast)
- Heap (if escapes)
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="v5ex"
func foo() int {
    x := 10
    return x
}
```

```text id="v5exout"
x stored in stack → destroyed after function ends
```

---

## **<span style="color:#8338ec">Escape Analysis</span>**

```text id="v5escape"
If variable is needed outside function → moved to heap
```

---

# **<span style="color:#ff1744">6. Standard Practices</span>**

---

## **<span style="color:#3a86ff">Prefer Short Declaration</span>**

```go id="v6a"
x := 10
```

✔ Cleaner and idiomatic

---

## **<span style="color:#3a86ff">Use var for Zero Value or Global</span>**

```go id="v6b"
var count int
```

---

## **<span style="color:#3a86ff">Group Related Variables</span>**

```go id="v6c"
var (
    name = "Arj"
    age  = 21
)
```

---

## **<span style="color:#3a86ff">Use Meaningful Names</span>**

```text id="v6name"
userAge ✔
x ❌ (unless small scope)
```

---

# **<span style="color:#ff1744">7. Precautions (Very Important)</span>**

---

## **<span style="color:#3a86ff">1. Avoid Shadowing Bugs</span>**

```go id="v7a"
err := doSomething()

if err != nil {
    err := handleError() // ❌ shadowing
}
```

---

## **<span style="color:#3a86ff">2. Don’t Ignore Zero Values</span>**

```go id="v7b"
var x int  // default = 0
```

✔ Always be aware

---

## **<span style="color:#3a86ff">3. Avoid Unused Variables</span>**

```go id="v7c"
x := 10 // ❌ if unused
```

---

## **<span style="color:#3a86ff">4. Be Careful with := in Loops</span>**

```go id="v7d"
for i := 0; i < 5; i++ {
    // new scope each iteration
}
```

---

## **<span style="color:#3a86ff">5. Explicit Types When Needed</span>**

```go id="v7e"
var price float64 = 10
```

✔ Avoid unintended inference

---

# **<span style="color:#ff1744">8. Combined Mental Model</span>**

---

```text id="v8model"
Declare variable →
Allocate memory →
Assign type →
Store value →
Use within scope →
Destroy when scope ends
```

---

# **<span style="color:#ff1744">9. Simple Exercise</span>**

---

## **<span style="color:#3a86ff">Goal</span>**

```text id="v9goal"
Understand scope + shadowing
```

---

## **<span style="color:#3a86ff">Task</span>**

---

### **Write this code**

```go id="v9task"
package main

import "fmt"

func main() {
    x := 10

    if true {
        x := 20
        fmt.Println("Inner:", x)
    }

    fmt.Println("Outer:", x)
}
```

---

## **<span style="color:#8338ec">Expected Output</span>**

```text id="v9out"
Inner: 20
Outer: 10
```

---

# **<span style="color:#ff1744">Final Summary</span>**

---

### **Variables in Go**

```text id="v10sum1"
Strongly typed memory containers
```

---

### **Key Features**

```text id="v10sum2"
Type safety
Explicit declaration
Zero values
Strict usage rules
```

---

### **Important Concepts**

```text id="v10sum3"
Scope
Shadowing
Escape analysis
```

---

# **<span style="color:#ff1744">Key Insight</span>**

```text id="v10insight"
Go variables are simple on surface,
but governed by strict rules for safety and performance
```

---

If you want deeper system-level understanding:

```text id="v10next"
- "Stack vs heap in Go (deep dive)"
- "Escape analysis with real examples"
- "Pointers and variable memory behavior"
```
