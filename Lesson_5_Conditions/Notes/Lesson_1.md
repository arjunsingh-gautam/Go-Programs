# **<span style="color:#ff1744">Conditional Statements in Go — Complete Mental Model</span>**

---

# **<span style="color:#ff6f00">1. What are Conditional Statements?</span>**

---

## **<span style="color:#3a86ff">Definition</span>**

```text id="c1def"
Conditional statements control program flow
based on boolean conditions
```

---

## **<span style="color:#3a86ff">Purpose</span>**

```text id="c1purpose"
Used for:
- Decision making
- Branching logic
- Validation
- Error handling
```

---

## **<span style="color:#8338ec">Core Idea</span>**

```text id="c1core"
Condition →
TRUE  → execute block A
FALSE → execute block B
```

---

# **<span style="color:#ff1744">2. Types of Conditional Statements in Go</span>**

---

## **<span style="color:#3a86ff">Main Conditional Constructs</span>**

```text id="c2types"
1. if
2. if-else
3. else-if ladder
4. nested if
5. switch
6. type switch
```

---

# **<span style="color:#ff1744">3. if Statement</span>**

---

# **<span style="color:#8338ec">Syntax</span>**

```go id="c3syntax"
if condition {
    // code
}
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="c3ex"
age := 20

if age >= 18 {
    println("Adult")
}
```

---

## **<span style="color:#3a86ff">Execution Mechanics</span>**

```text id="c3mech"
1. Evaluate condition
2. If TRUE → execute block
3. If FALSE → skip block
```

---

## **<span style="color:#8338ec">Dry Run</span>**

```text id="c3dry"
age = 20
↓
20 >= 18 → TRUE
↓
Print "Adult"
```

---

## **<span style="color:#3a86ff">When to Use</span>**

```text id="c3when"
✔ Single condition checks
✔ Validation logic
✔ Guard conditions
```

---

# **<span style="color:#ff1744">4. if-else Statement</span>**

---

# **<span style="color:#8338ec">Syntax</span>**

```go id="c4syntax"
if condition {
    // true block
} else {
    // false block
}
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="c4ex"
num := 5

if num%2 == 0 {
    println("Even")
} else {
    println("Odd")
}
```

---

## **<span style="color:#3a86ff">Execution Mechanics</span>**

```text id="c4mech"
Condition evaluated
↓
TRUE  → if block
FALSE → else block
```

---

## **<span style="color:#8338ec">Dry Run</span>**

```text id="c4dry"
num = 5
↓
5 % 2 == 0 → FALSE
↓
Execute else
↓
Print "Odd"
```

---

## **<span style="color:#3a86ff">When to Use</span>**

```text id="c4when"
✔ Binary decisions
✔ Either-or logic
```

---

# **<span style="color:#ff1744">5. else-if Ladder</span>**

---

# **<span style="color:#8338ec">Syntax</span>**

```go id="c5syntax"
if condition1 {

} else if condition2 {

} else {

}
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="c5ex"
marks := 85

if marks >= 90 {
    println("A")
} else if marks >= 75 {
    println("B")
} else {
    println("C")
}
```

---

## **<span style="color:#3a86ff">Execution Mechanics</span>**

```text id="c5mech"
Conditions checked TOP → DOWN
First TRUE block executes
Remaining skipped
```

---

## **<span style="color:#8338ec">Dry Run</span>**

```text id="c5dry"
marks = 85
↓
85 >= 90 → FALSE
↓
85 >= 75 → TRUE
↓
Print "B"
↓
Stop checking further
```

---

## **<span style="color:#3a86ff">When to Use</span>**

```text id="c5when"
✔ Multi-condition branching
✔ Ordered condition priority
```

---

# **<span style="color:#ff1744">6. Nested if Statements</span>**

---

# **<span style="color:#8338ec">Syntax</span>**

```go id="c6syntax"
if condition1 {
    if condition2 {
    }
}
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="c6ex"
age := 20
citizen := true

if age >= 18 {
    if citizen {
        println("Eligible")
    }
}
```

---

## **<span style="color:#3a86ff">Mechanics</span>**

```text id="c6mech"
Outer condition must pass
before inner condition executes
```

---

## **<span style="color:#3a86ff">When to Use</span>**

```text id="c6when"
✔ Dependent conditions
✔ Hierarchical validation
```

---

## **<span style="color:#8338ec">Precaution</span>**

```text id="c6prec"
Avoid deep nesting
→ hurts readability
```

---

# **<span style="color:#ff1744">7. if with Initialization (VERY IMPORTANT)</span>**

---

# **<span style="color:#8338ec">Special Go Feature</span>**

```go id="c7syntax"
if x := compute(); x > 10 {
    println(x)
}
```

---

## **<span style="color:#3a86ff">Mechanics</span>**

```text id="c7mech"
1. Initialize variable
2. Evaluate condition
3. Variable exists ONLY in if scope
```

---

## **<span style="color:#8338ec">Dry Run</span>**

```text id="c7dry"
x initialized
↓
Condition checked
↓
if block executes
↓
x destroyed after block
```

---

## **<span style="color:#3a86ff">Common Use Case</span>**

```go id="c7use"
if err := doSomething(); err != nil {
    println(err)
}
```

---

## **<span style="color:#8338ec">Key Insight</span>**

```text id="c7insight"
Keeps temporary variables scoped tightly
```

---

# **<span style="color:#ff1744">8. switch Statement</span>**

---

# **<span style="color:#8338ec">Syntax</span>**

```go id="c8syntax"
switch expression {
case value1:
    // code

case value2:
    // code

default:
    // code
}
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="c8ex"
day := 2

switch day {
case 1:
    println("Mon")

case 2:
    println("Tue")

default:
    println("Unknown")
}
```

---

## **<span style="color:#3a86ff">Mechanics</span>**

```text id="c8mech"
1. Evaluate expression once
2. Match cases
3. Execute matching block
4. Automatically break
```

---

## **<span style="color:#8338ec">Important Difference from C/C++</span>**

```text id="c8diff"
Go switch automatically breaks
(no fallthrough by default)
```

---

# **<span style="color:#ff1744">9. switch Without Expression</span>**

---

# **<span style="color:#8338ec">Syntax</span>**

```go id="c9syntax"
switch {
case condition1:
case condition2:
}
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="c9ex"
age := 20

switch {
case age >= 18:
    println("Adult")

case age >= 13:
    println("Teen")
}
```

---

## **<span style="color:#8338ec">Mental Model</span>**

```text id="c9model"
Acts like cleaner if-else ladder
```

---

# **<span style="color:#ff1744">10. Type Switch</span>**

---

# **<span style="color:#8338ec">Used with Interfaces</span>**

```go id="c10syntax"
switch v := i.(type) {
case int:
case string:
}
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="c10ex"
var x interface{} = "hello"

switch v := x.(type) {

case string:
    println("String:", v)

case int:
    println("Int:", v)
}
```

---

## **<span style="color:#3a86ff">Mechanics</span>**

```text id="c10mech"
Checks ACTUAL underlying type at runtime
```

---

# **<span style="color:#ff1744">11. Internal Mechanics of Conditional Statements</span>**

---

# **<span style="color:#8338ec">Execution Model</span>**

```text id="c11model"
Evaluate condition
↓
Convert to boolean
↓
Jump to matching block
↓
Skip remaining branches
```

---

## **<span style="color:#3a86ff">Compiler Optimization</span>**

```text id="c11opt"
Simple conditions may become jump instructions
```

---

# **<span style="color:#ff1744">12. Best Practices</span>**

---

## **<span style="color:#3a86ff">Do This</span>**

```text id="c12do"
✔ Keep conditions simple
✔ Use switch for many branches
✔ Use early returns
✔ Keep scopes minimal
✔ Use descriptive boolean names
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="c12ex"
if isValid {
}
```

✔ Better than:

```go id="c12bad"
if flag {
}
```

---

# **<span style="color:#ff1744">13. Precautions & Common Mistakes</span>**

---

# **<span style="color:#8338ec">1. Deep Nesting</span>**

```go id="c13a"
if a {
    if b {
        if c {
```

❌ Hard to read

---

## **<span style="color:#3a86ff">Better</span>**

```go id="c13a_fix"
if !a {
    return
}
```

---

# **<span style="color:#8338ec">2. Shadowing in if Initialization</span>**

```go id="c13b"
err := doA()

if err := doB(); err != nil {
}
```

❌ Different err variable

---

# **<span style="color:#8338ec">3. Complex Conditions</span>**

```go id="c13c"
if a && b || c && !d {
}
```

❌ Hard to reason about

---

# **<span style="color:#8338ec">4. Missing default in switch</span>**

```text id="c13d"
Unhandled cases possible
```

---

# **<span style="color:#ff1744">14. Combined Mental Model</span>**

---

```text id="c14model"
Conditionals =
decision-making gates
that control execution flow
```

---

### **Flow**

```text id="c14flow"
Evaluate →
Choose branch →
Execute →
Continue program
```

---

# **<span style="color:#ff1744">15. Simple Exercise</span>**

---

## **<span style="color:#3a86ff">Task</span>**

Write:

```text id="c15task"
1. if-else for even/odd
2. switch for weekdays
3. if initialization with error
4. nested condition for login access
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="c15ex"
num := 7

if num%2 == 0 {
    println("Even")
} else {
    println("Odd")
}
```

---

# **<span style="color:#ff1744">Final Summary</span>**

---

### **if**

```text id="c16sum1"
Simple condition checking
```

---

### **switch**

```text id="c16sum2"
Cleaner multi-branch logic
```

---

### **Special Go Feature**

```text id="c16sum3"
if with initialization
```

---

### **Core Philosophy**

```text id="c16sum4"
Simple, readable, explicit control flow
```

---

# **<span style="color:#ff1744">Key Insight</span>**

```text id="c16insight"
Go conditionals are intentionally minimal,
designed for clarity and predictable execution
```
