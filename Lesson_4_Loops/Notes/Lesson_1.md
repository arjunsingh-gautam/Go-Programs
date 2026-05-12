# **<span style="color:#ff1744">Iterative Statements in Go — Complete Mental Model</span>**

---

# **<span style="color:#ff6f00">1. What are Iterative Statements?</span>**

---

## **<span style="color:#3a86ff">Definition</span>**

```text id="it1def"
Iterative statements repeatedly execute a block of code
until a condition changes
```

---

## **<span style="color:#3a86ff">Purpose</span>**

```text id="it1purpose"
Used for:
- Repetition
- Traversing collections
- Polling/waiting
- Continuous processing
```

---

## **<span style="color:#8338ec">Core Idea</span>**

```text id="it1core"
Loop =
Initialize →
Check condition →
Execute →
Update →
Repeat
```

---

# **<span style="color:#ff1744">2. Iteration in Go — Important Insight</span>**

---

## **<span style="color:#3a86ff">Go Has ONLY One Loop Keyword</span>**

```text id="it2only"
Go uses ONLY:
for
```

---

## **<span style="color:#8338ec">Very Important</span>**

```text id="it2poly"
Go's for loop is POLYMORPHIC
→ behaves as:
- traditional for
- while
- infinite loop
- foreach/range loop
```

---

# **<span style="color:#ff1744">3. Traditional for Loop</span>**

---

# **<span style="color:#8338ec">Syntax</span>**

```go id="it3syntax"
for initialization; condition; update {
    // code
}
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="it3ex"
for i := 0; i < 5; i++ {
    println(i)
}
```

---

## **<span style="color:#3a86ff">Execution Flow</span>**

```text id="it3flow"
1. Initialization
2. Check condition
3. Execute body
4. Update
5. Repeat
```

---

## **<span style="color:#8338ec">Step-by-Step</span>**

```text id="it3steps"
i=0 → print
i=1 → print
i=2 → print
...
until condition fails
```

---

# **<span style="color:#ff1744">4. for as While Loop (Polymorphic Behavior)</span>**

---

# **<span style="color:#8338ec">Syntax</span>**

```go id="it4syntax"
for condition {
    // code
}
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="it4ex"
x := 0

for x < 5 {
    println(x)
    x++
}
```

---

## **<span style="color:#3a86ff">Why This Works</span>**

```text id="it4why"
Initialization and update are OPTIONAL
```

---

## **<span style="color:#8338ec">Important Insight</span>**

```text id="it4insight"
Go removed separate while keyword
because for is flexible enough
```

---

# **<span style="color:#ff1744">5. Infinite Loop in Go</span>**

---

# **<span style="color:#8338ec">Syntax</span>**

```go id="it5syntax"
for {
    // infinite loop
}
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="it5ex"
for {
    println("Running forever")
}
```

---

## **<span style="color:#3a86ff">Mechanics</span>**

```text id="it5mech"
No condition
→ always true
→ never exits automatically
```

---

## **<span style="color:#8338ec">Use Cases</span>**

```text id="it5use"
✔ Servers
✔ Event loops
✔ Workers
✔ Continuous monitoring
```

---

# **<span style="color:#ff1744">6. range Loop (foreach-like)</span>**

---

# **<span style="color:#8338ec">Syntax</span>**

```go id="it6syntax"
for index, value := range collection {
}
```

---

## **<span style="color:#3a86ff">Example with Slice</span>**

```go id="it6ex"
nums := []int{10, 20, 30}

for i, v := range nums {
    println(i, v)
}
```

---

## **<span style="color:#3a86ff">Output</span>**

```text id="it6out"
0 10
1 20
2 30
```

---

## **<span style="color:#8338ec">Ignoring Values</span>**

```go id="it6ignore"
for _, v := range nums {
    println(v)
}
```

---

# **<span style="color:#ff1744">7. break Keyword</span>**

---

# **<span style="color:#8338ec">Definition</span>**

```text id="it7def"
break immediately terminates the loop
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="it7ex"
for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }

    println(i)
}
```

---

## **<span style="color:#3a86ff">Output</span>**

```text id="it7out"
0
1
2
3
4
```

---

## **<span style="color:#8338ec">Execution Flow</span>**

```text id="it7flow"
Condition met →
break executes →
loop exits immediately
```

---

# **<span style="color:#ff1744">8. continue Keyword</span>**

---

# **<span style="color:#8338ec">Definition</span>**

```text id="it8def"
continue skips current iteration
and jumps to next iteration
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="it8ex"
for i := 0; i < 5; i++ {

    if i == 2 {
        continue
    }

    println(i)
}
```

---

## **<span style="color:#3a86ff">Output</span>**

```text id="it8out"
0
1
3
4
```

---

## **<span style="color:#8338ec">Mechanics</span>**

```text id="it8mech"
continue →
skip remaining body →
run next iteration
```

---

# **<span style="color:#ff1744">9. Internal Mechanics of Loops</span>**

---

# **<span style="color:#8338ec">Loop Execution Model</span>**

```text id="it9model"
Initialize
↓
Check condition
↓
Execute body
↓
Update state
↓
Repeat
```

---

## **<span style="color:#3a86ff">Stack Behavior</span>**

```text id="it9stack"
Loop variables usually live on stack
```

---

## **<span style="color:#3a86ff">Condition Evaluation</span>**

```text id="it9cond"
Condition checked BEFORE every iteration
```

---

# **<span style="color:#ff1744">10. Precautions While Using Loops</span>**

---

# **<span style="color:#8338ec">1. Infinite Loops Accidentally</span>**

```go id="it10a"
for x < 10 {
    // forgot x++
}
```

❌ Never exits

---

# **<span style="color:#8338ec">2. Shadowing Variables</span>**

```go id="it10b"
x := 0

for x < 5 {
    x := x + 1 // ❌ shadowing bug
}
```

---

# **<span style="color:#8338ec">3. Modifying Collection During range</span>**

```text id="it10c"
Can create unexpected behavior
```

---

# **<span style="color:#8338ec">4. Expensive Work Inside Loops</span>**

```text id="it10d"
Avoid unnecessary allocations
inside loops
```

---

# **<span style="color:#ff1744">11. Best Practices</span>**

---

## **<span style="color:#3a86ff">Do This</span>**

```text id="it11do"
✔ Keep loops simple
✔ Use range for collections
✔ Use break for early exit
✔ Use continue carefully
✔ Avoid deeply nested loops
```

---

## **<span style="color:#3a86ff">Avoid This</span>**

```text id="it11avoid"
✘ Complex loop conditions
✘ Infinite loops without exit strategy
✘ Heavy logic inside loops
✘ Unclear variable updates
```

---

# **<span style="color:#ff1744">12. Combined Mental Model</span>**

---

```text id="it12model"
for loop in Go =
universal iteration construct
```

---

### **Acts As**

```text id="it12acts"
Traditional for
While
Infinite loop
Foreach/range
```

---

# **<span style="color:#ff1744">13. Simple Exercise</span>**

---

## **<span style="color:#3a86ff">Task</span>**

Write:

```text id="it13task"
1. Traditional for loop
2. While-style loop
3. Infinite loop with break
4. range loop over slice
```

---

## **<span style="color:#3a86ff">Example Infinite Loop</span>**

```go id="it13ex"
x := 0

for {
    println(x)

    if x == 3 {
        break
    }

    x++
}
```

---

# **<span style="color:#ff1744">Final Summary</span>**

---

### **Go Iteration**

```text id="it14sum1"
Go uses ONLY for loop
```

---

### **Polymorphic Nature**

```text id="it14sum2"
for behaves as:
- for
- while
- infinite loop
- foreach
```

---

### **Important Keywords**

```text id="it14sum3"
break → exit loop
continue → skip iteration
```

---

# **<span style="color:#ff1744">Key Insight</span>**

```text id="it14insight"
Go simplifies iteration by making one powerful construct
instead of many separate loop keywords
```

---

If you want next deep dive:

```text id="it14next"
- "range loop internals and copy semantics"
- "Loop variable capture bug in goroutines"
- "Compiler optimization of loops in Go"
```
