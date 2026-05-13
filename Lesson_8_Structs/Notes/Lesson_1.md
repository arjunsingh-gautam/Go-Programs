# **<span style="color:#ff1744">Structs in Go — Complete Mental Model</span>**

---

# **<span style="color:#ff6f00">1. What is a Struct in Go?</span>**

---

## **<span style="color:#3a86ff">Definition</span>**

```text id="st1def"
A struct is a user-defined composite data type
that groups multiple related fields into one unit
```

---

## **<span style="color:#3a86ff">Core Idea</span>**

```text id="st1core"
Struct =
logical grouping of related data
```

---

## **<span style="color:#8338ec">Example</span>**

```go id="st1ex"
type User struct {
    Name string
    Age  int
}
```

---

## **<span style="color:#8338ec">Analogy</span>**

```text id="st1ana"
Struct = real-world object blueprint

User:
- name
- age
- email
```

---

# **<span style="color:#ff1744">2. Why Structs Exist</span>**

---

## **<span style="color:#3a86ff">Problem Without Structs</span>**

```go id="st2bad"
name := "Arj"
age := 21
email := "a@gmail.com"
```

---

## **<span style="color:#8338ec">Problems</span>**

```text id="st2prob"
- Data scattered
- Hard to manage
- Poor abstraction
```

---

# **<span style="color:#8338ec">Solution</span>**

```go id="st2good"
type User struct {
    Name  string
    Age   int
    Email string
}
```

---

## **<span style="color:#3a86ff">Benefit</span>**

```text id="st2benefit"
Related data grouped together
```

---

# **<span style="color:#ff1744">3. Struct Syntax</span>**

---

# **<span style="color:#8338ec">Basic Syntax</span>**

```go id="st3syntax"
type StructName struct {
    FieldName Type
}
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="st3ex"
type Product struct {
    ID    int
    Name  string
    Price float64
}
```

---

# **<span style="color:#ff1744">4. Ways to Declare & Initialize Structs</span>**

---

# **<span style="color:#8338ec">A. Zero Value Struct</span>**

```go id="st4a"
var u User
```

---

## **<span style="color:#3a86ff">Result</span>**

```text id="st4aout"
Name = ""
Age  = 0
```

---

# **<span style="color:#8338ec">B. Positional Initialization</span>**

```go id="st4b"
u := User{"Arj", 21}
```

---

## **<span style="color:#8338ec">Precaution</span>**

```text id="st4bprec"
Order MUST match struct fields
```

---

# **<span style="color:#8338ec">C. Named Field Initialization (Best Practice)</span>**

```go id="st4c"
u := User{
    Name: "Arj",
    Age:  21,
}
```

---

## **<span style="color:#3a86ff">Advantages</span>**

```text id="st4cadv"
✔ More readable
✔ Safer
✔ Order-independent
```

---

# **<span style="color:#8338ec">D. Pointer to Struct</span>**

```go id="st4d"
u := &User{
    Name: "Arj",
}
```

---

## **<span style="color:#3a86ff">Meaning</span>**

```text id="st4dmean"
u is pointer to struct
```

---

# **<span style="color:#8338ec">E. Using new()</span>**

```go id="st4e"
u := new(User)
```

---

## **<span style="color:#3a86ff">Result</span>**

```text id="st4eout"
Returns *User (pointer)
```

---

# **<span style="color:#ff1744">5. Accessing Struct Fields</span>**

---

# **<span style="color:#8338ec">Dot Operator</span>**

```go id="st5a"
u.Name = "Arj"
fmt.Println(u.Age)
```

---

# **<span style="color:#8338ec">Pointer Struct Access</span>**

```go id="st5b"
p := &u

p.Name = "Go"
```

---

## **<span style="color:#3a86ff">Important</span>**

```text id="st5bimp"
Go automatically dereferences struct pointers
```

---

# **<span style="color:#ff1744">6. Internal Memory Mechanics of Structs</span>**

---

# **<span style="color:#8338ec">Struct Memory Layout</span>**

```text id="st6layout"
Struct fields stored sequentially in memory
```

---

## **<span style="color:#3a86ff">Example</span>**

```go id="st6ex"
type User struct {
    Age int
    Flag bool
}
```

---

## **<span style="color:#3a86ff">Visualization</span>**

```text id="st6vis"
Memory:
| Age | Flag | Padding |
```

---

# **<span style="color:#8338ec">Padding & Alignment (VERY IMPORTANT)</span>**

---

## **Why Padding Exists</span>**

```text id="st6padwhy"
CPU accesses aligned memory faster
```

---

## **Example</span>**

```go id="st6padex"
type Bad struct {
    A bool
    B int64
}
```

May contain padding bytes

---

## **Better Layout</span>**

```go id="st6good"
type Good struct {
    B int64
    A bool
}
```

---

## **Best Practice Insight</span>**

```text id="st6insight"
Field order affects memory efficiency
```

---

# **<span style="color:#ff1744">7. Stack vs Heap Allocation</span>**

---

# **<span style="color:#8338ec">Small Structs</span>**

```text id="st7small"
Usually allocated on stack
```

---

# **<span style="color:#8338ec">Escaping Structs</span>**

```text id="st7escape"
Moved to heap if referenced outside scope
```

---

## **Example</span>**

```go id="st7ex"
func create() *User {
    u := User{}
    return &u
}
```

---

## **Mechanics</span>**

```text id="st7mech"
u escapes →
allocated on heap
```

---

# **<span style="color:#ff1744">8. Go Struct vs C Struct</span>**

---

# **<span style="color:#8338ec">Similarities</span>**

---

## **<span style="color:#3a86ff">1. Group Data</span>**

```text id="st8sim1"
Both combine related fields
```

---

## **<span style="color:#3a86ff">2. Sequential Memory Layout</span>**

```text id="st8sim2"
Fields stored contiguously
```

---

# **<span style="color:#8338ec">Differences (VERY IMPORTANT)</span>**

---

# **<span style="color:#3a86ff">1. Methods on Structs</span>**

---

## **Go</span>**

```go id="st8go1"
func (u User) Print() {
}
```

✔ Supported

---

## **C</span>**

```text id="st8c1"
Structs cannot directly own methods
```

---

# **<span style="color:#3a86ff">2. Encapsulation via Capitalization</span>**

```go id="st8go2"
Name string // exported
age  int    // private
```

---

# **<span style="color:#3a86ff">3. Automatic Garbage Collection</span>**

```text id="st8gc"
Go manages memory automatically
```

---

# **<span style="color:#3a86ff">4. No Inheritance</span>**

```text id="st8inherit"
Go uses composition instead
```

---

# **<span style="color:#ff1744">9. Embedded Structs (Composition)</span>**

---

# **<span style="color:#8338ec">Example</span>**

```go id="st9ex"
type Address struct {
    City string
}

type User struct {
    Name string
    Address
}
```

---

## **Usage</span>**

```go id="st9use"
u.City
```

---

## **Mental Model</span>**

```text id="st9model"
Struct embedding =
composition-based reuse
```

---

# **<span style="color:#ff1744">10. Use Cases of Structs</span>**

---

# **<span style="color:#8338ec">1. Domain Models</span>**

```go id="st10a"
User, Product, Order
```

---

# **<span style="color:#8338ec">2. API Request/Response</span>**

```go id="st10b"
JSON serialization
```

---

# **<span style="color:#8338ec">3. Database Models</span>**

```go id="st10c"
ORM structures
```

---

# **<span style="color:#8338ec">4. Configuration Objects</span>**

```go id="st10d"
type Config struct {
}
```

---

# **<span style="color:#ff1744">11. Best Practices</span>**

---

## **<span style="color:#3a86ff">1. Use Named Field Initialization</span>**

```go id="st11a"
User{Name: "Arj"}
```

✔ Readable

---

## **<span style="color:#3a86ff">2. Keep Structs Focused</span>**

```text id="st11b"
One struct → one responsibility
```

---

## **<span style="color:#3a86ff">3. Use Pointer Receivers for Large Structs</span>**

```go id="st11c"
func (u *User) Update() {
}
```

---

## **<span style="color:#3a86ff">4. Optimize Field Ordering</span>**

```text id="st11d"
Reduce padding/memory waste
```

---

## **<span style="color:#3a86ff">5. Export Only Necessary Fields</span>**

```text id="st11e"
Encapsulation improves maintainability
```

---

# **<span style="color:#ff1744">12. Precautions & Common Mistakes</span>**

---

# **<span style="color:#8338ec">1. Large Struct Copies</span>**

```go id="st12a"
u2 := u1
```

❌ Full copy happens

---

## **Better</span>**

```go id="st12a_fix"
u2 := &u1
```

---

# **<span style="color:#8338ec">2. Excessive Nesting</span>**

```text id="st12b"
Deep structs become hard to manage
```

---

# **<span style="color:#8338ec">3. Overusing Empty Interface Fields</span>**

```go id="st12c"
Data interface{}
```

❌ weak typing

---

# **<span style="color:#8338ec">4. Unexported JSON Fields</span>**

```go id="st12d"
name string
```

❌ not serialized

---

# **<span style="color:#ff1744">13. Combined Mental Model</span>**

---

```text id="st13model"
Struct =
custom memory layout +
related data grouping +
behavior attachment
```

---

## **Lifecycle</span>**

```text id="st13life"
Define blueprint
↓
Create instance
↓
Allocate memory
↓
Access/update fields
↓
Pass/copy/reference
```

---

# **<span style="color:#ff1744">14. Simple Exercise</span>**

---

## **<span style="color:#3a86ff">Task</span>**

```go id="st14task"
type Book struct {
    Title string
    Pages int
}

func main() {

    b := Book{
        Title: "Go",
        Pages: 300,
    }

    fmt.Println(b.Title)
}
```

---

## **Experiment</span>**

```text id="st14exp"
1. Create pointer to Book
2. Modify fields
3. Observe behavior
```

---

# **<span style="color:#ff1744">Final Summary</span>**

---

### **Structs**

```text id="st15sum1"
Custom composite data types
```

---

### **Core Purpose**

```text id="st15sum2"
Group related data cleanly
```

---

### **Internal Mechanics**

```text id="st15sum3"
Sequential memory layout
Padding/alignment
Stack vs heap allocation
```

---

### **Go vs C**

```text id="st15sum4"
Go adds:
- methods
- GC
- encapsulation
- composition
```

---

# **<span style="color:#ff1744">Key Insight</span>**

```text id="st15insight"
Go structs are lightweight, memory-efficient building blocks
for modeling real-world and system-level data
```
