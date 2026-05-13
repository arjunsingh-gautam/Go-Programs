# **<span style="color:#ff1744">Composition in Go — Complete Mental Model</span>**

---

# **<span style="color:#ff6f00">1. What is Composition in Go?</span>**

---

## **<span style="color:#3a86ff">Definition</span>**

```text id="cmp1def"
Composition means building complex types
by combining smaller types together
```

---

## **<span style="color:#3a86ff">Core Idea</span>**

```text id="cmp1core"
Instead of:
"is-a"

Go prefers:
"has-a"
```

---

## **<span style="color:#8338ec">Example</span>**

```go id="cmp1ex"
type Engine struct {
    Power int
}

type Car struct {
    Brand string
    Engine
}
```

---

## **Meaning</span>**

```text id="cmp1meaning"
Car HAS an Engine
```

---

# **<span style="color:#ff1744">2. Why Go Uses Composition Instead of Inheritance</span>**

---

# **<span style="color:#8338ec">Classical Inheritance (OOP)</span>**

```text id="cmp2inherit"
Child inherits parent behavior/data
```

---

## **Problem with Inheritance</span>**

```text id="cmp2prob"
- Tight coupling
- Deep hierarchies
- Fragile designs
- Hard maintenance
```

---

# **<span style="color:#8338ec">Go Philosophy</span>**

```text id="cmp2go"
Favor simple reusable components
instead of rigid inheritance trees
```

---

# **<span style="color:#ff1744">3. Composition vs Inheritance</span>**

---

| Concept      | Inheritance     | Composition     |
| ------------ | --------------- | --------------- |
| Relationship | IS-A            | HAS-A           |
| Coupling     | Tight           | Loose           |
| Reuse        | Parent-child    | Component reuse |
| Flexibility  | Lower           | Higher          |
| Complexity   | Can become deep | Simpler         |

---

# **<span style="color:#8338ec">Example</span>**

---

## **Inheritance Thinking</span>**

```text id="cmp3inherit"
Dog IS-A Animal
```

---

## **Composition Thinking</span>**

```text id="cmp3comp"
Car HAS-A Engine
User HAS-A Address
```

---

# **<span style="color:#ff1744">4. Syntax of Composition in Go</span>**

---

# **<span style="color:#8338ec">Embedded Struct Syntax</span>**

```go id="cmp4syntax"
type Child struct {
    Parent
}
```

---

## **Example</span>**

```go id="cmp4ex"
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

```go id="cmp4use"
u := User{
    Name: "Arj",
    Address: Address{
        City: "Mumbai",
    },
}

fmt.Println(u.City)
```

---

## **Important</span>**

```text id="cmp4important"
Embedded fields are PROMOTED automatically
```

---

# **<span style="color:#ff1744">5. Internal Mechanics of Composition</span>**

---

# **<span style="color:#8338ec">Very Important Insight</span>**

```text id="cmp5insight"
Go composition is NOT inheritance internally
```

---

## **Mechanics</span>**

```text id="cmp5mech"
Embedded struct becomes NORMAL field internally
```

---

# **<span style="color:#8338ec">Example</span>**

```go id="cmp5ex"
type Engine struct {
    Power int
}

type Car struct {
    Engine
}
```

Internally similar to:

```go id="cmp5internal"
type Car struct {
    Engine Engine
}
```

---

## **But Go Adds Field Promotion</span>**

```go id="cmp5promo"
c.Power
```

Internally becomes:

```text id="cmp5promo2"
c.Engine.Power
```

---

# **<span style="color:#ff1744">6. Method Promotion in Composition</span>**

---

# **<span style="color:#8338ec">Example</span>**

```go id="cmp6ex"
type Engine struct {
}

func (e Engine) Start() {
    fmt.Println("Engine started")
}

type Car struct {
    Engine
}
```

---

## **Usage</span>**

```go id="cmp6use"
c := Car{}

c.Start()
```

---

## **Mechanics</span>**

```text id="cmp6mech"
Go forwards method access
to embedded struct automatically
```

---

# **<span style="color:#ff1744">7. Detailed Real-World Example</span>**

---

# **<span style="color:#8338ec">Example: Smart Device System</span>**

```go id="cmp7ex"
package main

import "fmt"

type Battery struct {
    Capacity int
}

func (b Battery) ShowBattery() {
    fmt.Println("Battery:", b.Capacity)
}

type Camera struct {
    Megapixels int
}

func (c Camera) TakePhoto() {
    fmt.Println("Photo Taken")
}

type Smartphone struct {
    Brand string

    Battery
    Camera
}

func main() {

    phone := Smartphone{
        Brand: "GoPhone",

        Battery: Battery{
            Capacity: 5000,
        },

        Camera: Camera{
            Megapixels: 48,
        },
    }

    phone.ShowBattery()

    phone.TakePhoto()

    fmt.Println(phone.Brand)
}
```

---

# **<span style="color:#8338ec">Execution Flow</span>**

```text id="cmp7flow"
Smartphone contains:
- Battery
- Camera

Methods promoted automatically:
phone.ShowBattery()
phone.TakePhoto()
```

---

# **<span style="color:#ff1744">8. Composition + Interfaces (VERY IMPORTANT)</span>**

---

# **<span style="color:#8338ec">Powerful Go Pattern</span>**

```text id="cmp8pattern"
Composition + interfaces =
Go's alternative to inheritance polymorphism
```

---

## **Example</span>**

```go id="cmp8ex"
type Writer interface {
    Write()
}

type FileWriter struct {
}

func (f FileWriter) Write() {
}
```

---

## **Mental Model</span>**

```text id="cmp8model"
Behavior reuse through interfaces
Data reuse through composition
```

---

# **<span style="color:#ff1744">9. Advantages of Composition</span>**

---

# **<span style="color:#8338ec">1. Loose Coupling</span>**

```text id="cmp9a"
Components independent
```

---

# **<span style="color:#8338ec">2. Better Reusability</span>**

```text id="cmp9b"
Reuse small focused pieces
```

---

# **<span style="color:#8338ec">3. Flexible Designs</span>**

```text id="cmp9c"
Mix components dynamically
```

---

# **<span style="color:#8338ec">4. Easier Testing</span>**

```text id="cmp9d"
Mock components independently
```

---

# **<span style="color:#ff1744">10. Use Cases of Composition</span>**

---

# **<span style="color:#8338ec">1. Domain Modeling</span>**

```text id="cmp10a"
User HAS Address
Order HAS Payment
```

---

# **<span style="color:#8338ec">2. Middleware Systems</span>**

```text id="cmp10b"
HTTP handlers composed together
```

---

# **<span style="color:#8338ec">3. Shared Capabilities</span>**

```text id="cmp10c"
Logging
Caching
Database access
```

---

# **<span style="color:#8338ec">4. Microservice Components</span>**

```text id="cmp10d"
Reusable infrastructure modules
```

---

# **<span style="color:#ff1744">11. Best Practices</span>**

---

## **<span style="color:#3a86ff">1. Prefer Small Focused Structs</span>**

```text id="cmp11a"
Single responsibility
```

---

## **<span style="color:#3a86ff">2. Use Composition Over Giant Structs</span>**

```text id="cmp11b"
Split concerns cleanly
```

---

## **<span style="color:#3a86ff">3. Combine with Interfaces</span>**

```text id="cmp11c"
Highly flexible architecture
```

---

## **<span style="color:#3a86ff">4. Avoid Deep Embedding</span>**

```text id="cmp11d"
Too much nesting hurts readability
```

---

## **<span style="color:#3a86ff">5. Explicit Naming When Needed</span>**

```go id="cmp11e"
Engine Engine
```

Sometimes clearer than anonymous embedding

---

# **<span style="color:#ff1744">12. Precautions & Common Mistakes</span>**

---

# **<span style="color:#8338ec">1. Field Name Conflicts</span>**

```go id="cmp12a"
type A struct { Name string }
type B struct { Name string }
```

Embedded together:

```text id="cmp12aout"
Ambiguous selector errors
```

---

# **<span style="color:#8338ec">2. Overusing Embedding</span>**

```text id="cmp12b"
Not every relationship should be embedded
```

---

# **<span style="color:#8338ec">3. Assuming True Inheritance</span>**

```text id="cmp12c"
Go does NOT support:
- superclass
- override hierarchy
- protected/private inheritance
```

---

# **<span style="color:#8338ec">4. Hidden Coupling</span>**

```text id="cmp12d"
Excessive embedding can obscure ownership
```

---

# **<span style="color:#ff1744">13. Composition vs Method Override</span>**

---

# **<span style="color:#8338ec">Go Does NOT Override Like OOP</span>**

---

## **Example</span>**

```go id="cmp13ex"
type Animal struct {}

func (a Animal) Speak() {
    fmt.Println("Animal")
}

type Dog struct {
    Animal
}

func (d Dog) Speak() {
    fmt.Println("Dog")
}
```

---

## **Mechanics</span>**

```text id="cmp13mech"
Dog has its own Speak()
No virtual inheritance table like Java/C++
```

---

# **<span style="color:#ff1744">14. Combined Mental Model</span>**

---

```text id="cmp14model"
Composition =
building systems from reusable LEGO blocks
```

---

## **Flow</span>**

```text id="cmp14flow"
Small structs
↓
Combine together
↓
Promote fields/methods
↓
Build complex behavior
```

---

# **<span style="color:#ff1744">15. Simple Exercise</span>**

---

## **<span style="color:#3a86ff">Task</span>**

Create:

```text id="cmp15task"
1. Engine struct
2. GPS struct
3. Car struct embedding both
4. Methods:
   - Start()
   - Navigate()
```

---

## **Goal</span>**

```text id="cmp15goal"
Observe:
- method promotion
- field access
- reuse through composition
```

---

# **<span style="color:#ff1744">Final Summary</span>**

---

### **Composition**

```text id="cmp16sum1"
Combining structs to build larger systems
```

---

### **Go Philosophy**

```text id="cmp16sum2"
Prefer HAS-A over IS-A
```

---

### **Internal Mechanics**

```text id="cmp16sum3"
Embedded structs are normal fields
with automatic field/method promotion
```

---

### **Main Advantages**

```text id="cmp16sum4"
Flexibility
Loose coupling
Reusability
Simplicity
```

---

# **<span style="color:#ff1744">Key Insight</span>**

```text id="cmp16insight"
Go avoids complex inheritance hierarchies
by combining composition + interfaces
to build clean scalable systems
```
