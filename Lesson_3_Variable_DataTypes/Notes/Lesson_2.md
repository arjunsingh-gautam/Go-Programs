# **<span style="color:#ff1744">Data Types in Go — Complete Mental Model</span>**

---

# **<span style="color:#ff6f00">1. What is a Data Type in Go?</span>**

---

## **<span style="color:#3a86ff">Definition</span>**

```text id="d1def"
A data type defines:
- What kind of value a variable holds
- How much memory it uses
- What operations are allowed
```

---

## **<span style="color:#3a86ff">Core Idea</span>**

```text id="d1core"
Type = structure + behavior + memory rules
```

---

## **<span style="color:#8338ec">Analogy</span>**

```text id="d1ana"
Data type = blueprint of a container
(int box, string box, struct box)
```

---

# **<span style="color:#ff1744">2. Categories of Data Types in Go</span>**

---

## **<span style="color:#3a86ff">Main Classification</span>**

```text id="d2types"
1. Basic Types
2. Aggregate Types
3. Reference Types
4. Interface Types
5. Other Special Types
```

---

# **<span style="color:#ff1744">3. Basic Data Types</span>**

---

# **<span style="color:#8338ec">A. Integer Types</span>**

```go id="d3int"
var a int = 10
var b int8 = 127
var c uint = 20
```

---

## **<span style="color:#3a86ff">Types</span>**

```text id="d3inttypes"
int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64
uintptr
```

---

## **<span style="color:#3a86ff">Best Practices</span>**

```text id="d3intbp"
✔ Use int for general purpose
✔ Use specific size only when needed (memory/interop)
✔ Avoid mixing signed & unsigned
```

---

# **<span style="color:#8338ec">B. Floating Point</span>**

```go id="d3float"
var pi float64 = 3.14
```

---

## **<span style="color:#3a86ff">Types</span>**

```text id="d3floattypes"
float32, float64
```

---

## **<span style="color:#3a86ff">Best Practices</span>**

```text id="d3floatbp"
✔ Use float64 by default
✔ Avoid equality comparison (precision issues)
```

---

# **<span style="color:#8338ec">C. Boolean</span>**

```go id="d3bool"
var flag bool = true
```

---

## **<span style="color:#3a86ff">Best Practices</span>**

```text id="d3boolbp"
✔ Use clear naming (isValid, hasAccess)
```

---

# **<span style="color:#8338ec">D. String</span>**

```go id="d3string"
var name string = "Go"
```

---

## **<span style="color:#3a86ff">Important</span>**

```text id="d3stringimp"
Strings are IMMUTABLE
```

---

## **<span style="color:#3a86ff">Best Practices</span>**

```text id="d3stringbp"
✔ Use string builder for heavy concatenation
✔ Be aware of UTF-8 encoding
```

---

# **<span style="color:#ff1744">4. Aggregate Data Types</span>**

---

# **<span style="color:#8338ec">A. Arrays</span>**

```go id="d4array"
var arr [3]int = [3]int{1, 2, 3}
```

---

## **<span style="color:#3a86ff">Properties</span>**

```text id="d4arrayprop"
Fixed size
Value type (copied on assignment)
```

---

## **<span style="color:#3a86ff">Best Practices</span>**

```text id="d4arraybp"
✔ Rarely used directly
✔ Use slices instead
```

---

# **<span style="color:#8338ec">B. Structs</span>**

```go id="d4struct"
type User struct {
    Name string
    Age  int
}
```

---

## **<span style="color:#3a86ff">Usage</span>**

```go id="d4structuse"
u := User{Name: "Arj", Age: 21}
```

---

## **<span style="color:#3a86ff">Best Practices</span>**

```text id="d4structbp"
✔ Use for grouping related data
✔ Keep fields exported only if needed
✔ Use tags for JSON/db
```

---

# **<span style="color:#ff1744">5. Reference Data Types</span>**

---

# **<span style="color:#8338ec">A. Slices</span>**

```go id="d5slice"
s := []int{1, 2, 3}
```

---

## **<span style="color:#3a86ff">Properties</span>**

```text id="d5sliceprop"
Dynamic size
Reference to underlying array
```

---

## **<span style="color:#3a86ff">Best Practices</span>**

```text id="d5slicebp"
✔ Use slices instead of arrays
✔ Be careful with append (capacity changes)
```

---

# **<span style="color:#8338ec">B. Maps</span>**

```go id="d5map"
m := map[string]int{"a": 1}
```

---

## **<span style="color:#3a86ff">Properties</span>**

```text id="d5mapprop"
Key-value store
Reference type
```

---

## **<span style="color:#3a86ff">Best Practices</span>**

```text id="d5mapbp"
✔ Always check key existence
✔ Initialize before use
```

---

# **<span style="color:#8338ec">C. Pointers</span>**

```go id="d5ptr"
x := 10
p := &x
```

---

## **<span style="color:#3a86ff">Usage</span>**

```go id="d5ptruse"
*p = 20
```

---

## **<span style="color:#3a86ff">Best Practices</span>**

```text id="d5ptrbp"
✔ Use pointers to avoid copying large structs
✔ Avoid unnecessary pointer usage
```

---

# **<span style="color:#ff1744">6. Interface Type</span>**

---

## **<span style="color:#3a86ff">Definition</span>**

```text id="d6def"
Interface defines behavior (methods), not data
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="d6ex"
type Shape interface {
    Area() float64
}
```

---

## **<span style="color:#3a86ff">Implementation</span>**

```go id="d6impl"
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}
```

---

## **<span style="color:#8338ec">Key Insight</span>**

```text id="d6insight"
Go uses implicit interface implementation
(no "implements" keyword)
```

---

## **<span style="color:#3a86ff">Best Practices</span>**

```text id="d6bp"
✔ Keep interfaces small
✔ Define interfaces where used
```

---

# **<span style="color:#ff1744">7. Other Important Data Types</span>**

---

# **<span style="color:#8338ec">A. Function Type</span>**

```go id="d7func"
var f func(int) int
```

---

# **<span style="color:#8338ec">B. Channel (Concurrency)</span>**

```go id="d7chan"
ch := make(chan int)
```

---

# **<span style="color:#8338ec">C. Rune & Byte</span>**

```go id="d7rune"
var r rune = 'A'
var b byte = 255
```

---

## **<span style="color:#3a86ff">Difference</span>**

```text id="d7diff"
rune = int32 (Unicode)
byte = uint8
```

---

# **<span style="color:#ff1744">8. Combined Mental Model</span>**

---

```text id="d8model"
Basic types → simple values
Aggregate → group values
Reference → point to shared data
Interface → define behavior
```

---

# **<span style="color:#ff1744">9. Common Mistakes</span>**

---

## **<span style="color:#3a86ff">Mistake 1</span>**

```text id="d9m1"
Using array instead of slice
```

---

## **<span style="color:#3a86ff">Mistake 2</span>**

```text id="d9m2"
Ignoring nil map initialization
```

---

## **<span style="color:#3a86ff">Mistake 3</span>**

```text id="d9m3"
Overusing pointers
```

---

## **<span style="color:#3a86ff">Mistake 4</span>**

```text id="d9m4"
Large interfaces (bad design)
```

---

# **<span style="color:#ff1744">10. Simple Exercise</span>**

---

## **<span style="color:#3a86ff">Task</span>**

```go id="d10task"
package main

import "fmt"

type User struct {
    Name string
    Age  int
}

func main() {
    u := User{"Arj", 21}

    s := []int{1, 2, 3}

    m := map[string]int{"a": 1}

    fmt.Println(u, s, m)
}
```

---

## **<span style="color:#8338ec">Goal</span>**

```text id="d10goal"
Identify:
- struct → aggregate
- slice/map → reference
- string/int → basic
```

---

# **<span style="color:#ff1744">Final Summary</span>**

---

### **Basic Types**

```text id="d11sum1"
int, float, bool, string
```

---

### **Aggregate Types**

```text id="d11sum2"
array, struct
```

---

### **Reference Types**

```text id="d11sum3"
slice, map, pointer
```

---

### **Advanced Types**

```text id="d11sum4"
interface, function, channel
```

---

# **<span style="color:#ff1744">Key Insight</span>**

```text id="d11insight"
Go types are simple but powerful,
designed for performance, safety, and clarity
```

---

If you want next deep dive:

```text id="d11next"
- "Slice internals (len, cap, underlying array)"
- "Interface internals (itab, dynamic dispatch)"
- "Memory layout of structs and alignment"
```
