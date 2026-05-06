# **<span style="color:#ff1744">Type Conversions in Go — Complete Mental Model</span>**

---

# **<span style="color:#ff6f00">1. What is Type Conversion in Go?</span>**

---

## **<span style="color:#3a86ff">Definition</span>**

```text id="tc1def"
Type conversion = converting a value from one data type to another explicitly
```

---

## **<span style="color:#3a86ff">Syntax</span>**

```go id="tc1syntax"
T(value)
```

---

## **<span style="color:#8338ec">Example</span>**

```go id="tc1ex"
var a int = 10
var b float64 = float64(a)
```

---

## **<span style="color:#8338ec">Key Insight</span>**

```text id="tc1insight"
Go does NOT support implicit type conversion
→ Everything must be explicit
```

---

# **<span style="color:#ff1744">2. Why Go Requires Explicit Conversion</span>**

---

## **<span style="color:#3a86ff">Reason</span>**

```text id="tc2reason"
To avoid hidden bugs and unintended data loss
```

---

## **<span style="color:#8338ec">Example Problem</span>**

```text id="tc2prob"
int → float → precision issues
int64 → int32 → overflow risk
```

---

# **<span style="color:#ff1744">3. Types of Type Conversion in Go</span>**

---

# **<span style="color:#8338ec">A. Numeric Conversions</span>**

---

## **<span style="color:#3a86ff">Example</span>**

```go id="tc3a"
var a int = 10
var b float64 = float64(a)
```

---

## **<span style="color:#3a86ff">Between Integers</span>**

```go id="tc3a2"
var x int64 = 100
var y int32 = int32(x)
```

---

## **<span style="color:#8338ec">Important Behavior</span>**

```text id="tc3abeh"
- Truncation can happen
- Overflow NOT checked at runtime
```

---

# **<span style="color:#8338ec">B. Float ↔ Integer</span>**

---

```go id="tc3b"
var f float64 = 3.7
var i int = int(f)
```

Output:

```text id="tc3bout"
i = 3  (decimal truncated)
```

---

## **<span style="color:#3a86ff">Best Practice</span>**

```text id="tc3bbp"
✔ Be careful: decimal part is LOST
```

---

# **<span style="color:#8338ec">C. String Conversions</span>**

---

## **<span style="color:#3a86ff">Number → String</span>**

```go id="tc3c"
import "strconv"

s := strconv.Itoa(10)
```

---

## **<span style="color:#3a86ff">String → Number</span>**

```go id="tc3c2"
n, _ := strconv.Atoi("123")
```

---

## **<span style="color:#8338ec">Important</span>**

```text id="tc3cimp"
String conversions require helper packages (strconv)
```

---

# **<span style="color:#8338ec">D. Byte, Rune, String Conversion</span>**

---

```go id="tc3d"
var s string = "A"
var b byte = s[0]
var r rune = rune(s[0])
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="tc3d2"
r := rune('A')
fmt.Println(string(r)) // "A"
```

---

# **<span style="color:#8338ec">E. Slice ↔ Array (Partial)</span>**

---

```go id="tc3e"
arr := [3]int{1, 2, 3}
slice := arr[:] // slice from array
```

---

## **<span style="color:#3a86ff">Important</span>**

```text id="tc3eimp"
Array → slice ✔
Slice → array ❌ (directly)
```

---

# **<span style="color:#8338ec">F. Interface Type Conversion (Type Assertion)</span>**

---

```go id="tc3f"
var i interface{} = "hello"

s := i.(string)
```

---

## **<span style="color:#3a86ff">Safe Version</span>**

```go id="tc3f2"
s, ok := i.(string)
```

---

# **<span style="color:#ff1744">4. Type Casting vs Type Assertion</span>**

---

| Concept         | Meaning                              |
| --------------- | ------------------------------------ |
| Type Conversion | Change value type                    |
| Type Assertion  | Extract concrete type from interface |

---

# **<span style="color:#ff1744">5. Internal Mechanics</span>**

---

## **<span style="color:#8338ec">Numeric Conversion</span>**

```text id="tc5num"
Value is reinterpreted into new type
(no runtime safety checks)
```

---

## **<span style="color:#8338ec">String Conversion</span>**

```text id="tc5str"
Uses library functions → parsing/formatting
```

---

## **<span style="color:#8338ec">Interface Assertion</span>**

```text id="tc5iface"
Checks underlying type at runtime
→ may panic if wrong
```

---

# **<span style="color:#ff1744">6. Use Cases of Type Conversion</span>**

---

# **<span style="color:#8338ec">Use Case 1: Arithmetic Compatibility</span>**

```go id="tc6a"
var a int = 10
var b float64 = 2.5

result := float64(a) + b
```

---

# **<span style="color:#8338ec">Use Case 2: Input Parsing</span>**

```go id="tc6b"
input := "123"
num, _ := strconv.Atoi(input)
```

---

# **<span style="color:#8338ec">Use Case 3: Working with Interfaces</span>**

```go id="tc6c"
var i interface{} = 42
num := i.(int)
```

---

# **<span style="color:#8338ec">Use Case 4: Byte/String Processing</span>**

```go id="tc6d"
b := []byte("hello")
s := string(b)
```

---

# **<span style="color:#ff1744">7. Common Pitfalls</span>**

---

## **<span style="color:#3a86ff">1. Data Loss</span>**

```go id="tc7a"
float64 → int  // loses decimals
```

---

## **<span style="color:#3a86ff">2. Overflow</span>**

```go id="tc7b"
int64 → int8  // may overflow silently
```

---

## **<span style="color:#3a86ff">3. Panic in Type Assertion</span>**

```go id="tc7c"
i.(int) // panic if not int
```

---

## **<span style="color:#3a86ff">4. Ignoring Errors</span>**

```go id="tc7d"
n, _ := strconv.Atoi("abc") // ❌ bad practice
```

---

# **<span style="color:#ff1744">8. Best Practices</span>**

---

## **<span style="color:#3a86ff">Do This</span>**

```text id="tc8do"
✔ Always convert explicitly
✔ Handle errors (strconv)
✔ Use safe type assertions
✔ Be aware of precision loss
```

---

## **<span style="color:#3a86ff">Avoid This</span>**

```text id="tc8avoid"
✘ Blind conversions
✘ Ignoring overflow risks
✘ Unsafe interface assertions
```

---

# **<span style="color:#ff1744">9. Combined Mental Model</span>**

---

```text id="tc9model"
Value → explicit conversion → new type
(no automatic magic)
```

---

# **<span style="color:#ff1744">10. Simple Exercise</span>**

---

## **<span style="color:#3a86ff">Task</span>**

```go id="tc10task"
package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := "50"

    num, _ := strconv.Atoi(s)

    result := float64(num) * 1.5

    fmt.Println(result)
}
```

---

## **<span style="color:#8338ec">Goal</span>**

```text id="tc10goal"
Practice:
- string → int
- int → float
```

---

# **<span style="color:#ff1744">Final Summary</span>**

---

### **Type Conversion**

```text id="tc11sum1"
Explicit transformation between types
```

---

### **Key Features**

```text id="tc11sum2"
No implicit conversion
Strict type safety
Manual control
```

---

### **Important Concepts**

```text id="tc11sum3"
Numeric conversion
String parsing
Type assertion
```

---

# **<span style="color:#ff1744">Key Insight</span>**

```text id="tc11insight"
Go forces you to be explicit with types,
so you stay in control of correctness and performance
```

---

If you want deeper next:

```text id="tc11next"
- "Interface internals (itab + dynamic type)"
- "Unsafe conversions using unsafe.Pointer"
- "Memory layout impact of type conversions"
```
