# **<span style="color:#ff1744">Functions in Go — Complete Mental Model</span>**

---

# **<span style="color:#ff6f00">1. What is a Function in Go?</span>**

---

## **<span style="color:#3a86ff">Definition</span>**

```text id="f1def"
A function is a reusable block of code that performs a specific task
```

---

## **<span style="color:#3a86ff">Basic Syntax</span>**

```go id="f1code"
func functionName(parameters) returnType {
    // logic
}
```

---

## **<span style="color:#8338ec">Example</span>**

```go id="f1ex"
func add(a int, b int) int {
    return a + b
}
```

---

## **<span style="color:#8338ec">Analogy</span>**

```text id="f1ana"
Function = machine
Input → processing → output
```

---

# **<span style="color:#ff1744">2. Ways to Declare, Define, and Call Functions</span>**

---

# **<span style="color:#8338ec">A. Basic Function</span>**

```go id="f2a"
func greet() {
    println("Hello")
}
```

Call:

```go id="f2a_call"
greet()
```

---

# **<span style="color:#8338ec">B. Function with Parameters</span>**

```go id="f2b"
func greet(name string) {
    println("Hello", name)
}
```

---

# **<span style="color:#8338ec">C. Multiple Parameters (Same Type Shortcut)</span>**

```go id="f2c"
func add(a, b int) int {
    return a + b
}
```

---

# **<span style="color:#8338ec">D. Multiple Return Values (VERY IMPORTANT)</span>**

```go id="f2d"
func divide(a, b int) (int, int) {
    return a / b, a % b
}
```

Usage:

```go id="f2d_call"
q, r := divide(10, 3)
```

---

# **<span style="color:#8338ec">E. Named Return Values</span>**

```go id="f2e"
func calc(a, b int) (sum int, diff int) {
    sum = a + b
    diff = a - b
    return
}
```

---

# **<span style="color:#8338ec">F. Variadic Functions</span>**

```go id="f2f"
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

---

# **<span style="color:#8338ec">G. Anonymous Functions</span>**

```go id="f2g"
func() {
    println("Anonymous")
}()
```

---

# **<span style="color:#8338ec">H. Function as Value (First-Class Functions)</span>**

```go id="f2h"
f := func(a int) int {
    return a * 2
}

println(f(5))
```

---

# **<span style="color:#8338ec">I. Function Returning Function</span>**

```go id="f2i"
func multiplier(x int) func(int) int {
    return func(y int) int {
        return x * y
    }
}
```

---

# **<span style="color:#ff1744">3. How Go Functions Work Internally</span>**

---

# **<span style="color:#8338ec">A. Call Stack Mechanism</span>**

```text id="f3stack"
Functions execute using a call stack (LIFO)
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="f3ex"
func A() { B() }
func B() { C() }
func C() { println("Done") }
```

---

## **<span style="color:#3a86ff">Execution Flow</span>**

```text id="f3flow"
Call A
↓
Call B
↓
Call C
↓
Return C
↓
Return B
↓
Return A
```

---

# **<span style="color:#8338ec">B. Stack Frame (VERY IMPORTANT)</span>**

Each function call creates:

```text id="f3frame"
Stack Frame containing:
- Parameters
- Local variables
- Return address
```

---

## **<span style="color:#3a86ff">Visualization</span>**

```text id="f3vis"
| C frame |
| B frame |
| A frame |
```

---

# **<span style="color:#8338ec">C. Memory Behavior</span>**

```text id="f3mem"
Small data → stack
Large/escaped data → heap
```

---

## **<span style="color:#3a86ff">Escape Analysis</span>**

```text id="f3escape"
Go decides:
Should variable stay on stack or move to heap?
```

---

# **<span style="color:#ff1744">4. Special Features of Go Functions (vs Python)</span>**

---

## **<span style="color:#3a86ff">1. Multiple Return Values</span>**

```text id="f4multi"
Go supports returning multiple values natively
Python uses tuples
```

---

## **<span style="color:#3a86ff">2. Explicit Types (Static Typing)</span>**

```text id="f4types"
Go → compile-time type checking
Python → runtime typing
```

---

## **<span style="color:#3a86ff">3. No Default Arguments</span>**

```text id="f4default"
Go does NOT support default parameters
```

---

## **<span style="color:#3a86ff">4. Error Handling Pattern</span>**

```go id="f4err"
func read() (data string, err error)
```

```text id="f4errtxt"
Return error as second value
```

---

## **<span style="color:#3a86ff">5. No Function Overloading</span>**

```text id="f4over"
Same function name with different params ❌ not allowed
```

---

## **<span style="color:#3a86ff">6. Closures</span>**

```go id="f4closure"
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

---

# **<span style="color:#ff1744">5. Things to Take Care While Writing Functions</span>**

---

## **<span style="color:#3a86ff">1. Keep Functions Small</span>**

```text id="f5small"
One function = one responsibility
```

---

## **<span style="color:#3a86ff">2. Avoid Too Many Parameters</span>**

```text id="f5params"
Use struct if parameters grow
```

---

## **<span style="color:#3a86ff">3. Handle Errors Explicitly</span>**

```go id="f5err"
if err != nil {
    return err
}
```

---

## **<span style="color:#3a86ff">4. Use Meaningful Names</span>**

```text id="f5name"
addNumbers() ✔
fn1() ❌
```

---

## **<span style="color:#3a86ff">5. Avoid Global State</span>**

```text id="f5global"
Pass dependencies explicitly
```

---

## **<span style="color:#3a86ff">6. Prefer Pure Functions</span>**

```text id="f5pure"
Same input → same output
```

---

# **<span style="color:#ff1744">6. Best Practices</span>**

---

## **<span style="color:#3a86ff">Do This</span>**

```text id="f6do"
✔ Return (value, error)
✔ Use early returns
✔ Keep logic readable
✔ Write testable functions
✔ Use closures carefully
```

---

## **<span style="color:#3a86ff">Avoid This</span>**

```text id="f6avoid"
✘ Deep nested logic
✘ Ignoring errors
✘ Huge functions
✘ Side effects everywhere
```

---

# **<span style="color:#ff1744">7. Mental Model</span>**

---

```text id="f7model"
Function call →
Create stack frame →
Execute →
Return →
Destroy frame
```

---

# **<span style="color:#ff1744">8. Simple Exercise</span>**

---

## **<span style="color:#3a86ff">Goal</span>**

```text id="f8goal"
Understand multi-return + closures
```

---

## **<span style="color:#3a86ff">Task</span>**

---

### **1. Write this function**

```go id="f8task1"
func stats(a, b int) (int, int) {
    return a + b, a * b
}
```

---

### **2. Call it**

```go id="f8task2"
sum, prod := stats(3, 4)
println(sum, prod)
```

---

### **3. Write closure**

```go id="f8task3"
func counter() func() int {
    c := 0
    return func() int {
        c++
        return c
    }
}
```

---

### **4. Test**

```go id="f8task4"
c := counter()
println(c()) // 1
println(c()) // 2
```

---

# **<span style="color:#ff1744">Final Summary</span>**

---

### **Functions in Go**

```text id="f9sum1"
Reusable, typed, structured blocks
```

---

### **Key Features**

```text id="f9sum2"
Multiple returns
Closures
Static typing
Explicit error handling
```

---

### **Internal Working**

```text id="f9sum3"
Call stack + stack frames + escape analysis
```

---

# **<span style="color:#ff1744">Key Insight</span>**

```text id="f9insight"
Go functions are simple on the surface,
but backed by powerful execution mechanics and strict design rules
```

---

If you want deeper system-level mastery next:

```text id="f9next"
- "Escape analysis deep dive"
- "Stack vs heap in Go"
- "Closures memory behavior internally"
- "Inlining optimization in Go compiler"
```
