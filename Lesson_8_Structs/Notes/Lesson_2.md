# **<span style="color:#ff1744">Methods on Structs in Go — Complete Mental Model</span>**

---

# **<span style="color:#ff6f00">1. What is a Method in Go?</span>**

---

## **<span style="color:#3a86ff">Definition</span>**

```text id="m1def"
A method is a function attached to a specific type
(usually a struct)
```

---

## **<span style="color:#3a86ff">Core Idea</span>**

```text id="m1core"
Method =
function + receiver
```

---

## **<span style="color:#8338ec">Example</span>**

```go id="m1ex"
type User struct {
    Name string
}

func (u User) Greet() {
    fmt.Println("Hello", u.Name)
}
```

---

## **<span style="color:#8338ec">Usage</span>**

```go id="m1use"
u := User{Name: "Arj"}

u.Greet()
```

---

# **<span style="color:#ff1744">2. Why Methods Exist</span>**

---

## **<span style="color:#3a86ff">Problem Without Methods</span>**

```go id="m2bad"
func greet(u User) {
}
```

---

## **Problems</span>**

```text id="m2prob"
- Logic disconnected from data
- Poor organization
- Harder abstraction
```

---

# **<span style="color:#8338ec">Solution</span>**

```go id="m2good"
func (u User) Greet() {
}
```

---

## **Benefit</span>**

```text id="m2benefit"
Behavior grouped with related data
```

---

# **<span style="color:#ff1744">3. Syntax of Methods in Go</span>**

---

# **<span style="color:#8338ec">General Syntax</span>**

```go id="m3syntax"
func (receiver ReceiverType) MethodName(parameters) returnType {
}
```

---

## **Breakdown</span>**

```text id="m3break"
receiver → object receiving method
MethodName → function attached to type
```

---

# **<span style="color:#ff1744">4. Receiver — Most Important Concept</span>**

---

# **<span style="color:#8338ec">Value Receiver</span>**

```go id="m4value"
func (u User) Print() {
}
```

---

## **Meaning</span>**

```text id="m4valuemean"
Method gets COPY of struct
```

---

# **<span style="color:#8338ec">Pointer Receiver</span>**

```go id="m4ptr"
func (u *User) Update() {
}
```

---

## **Meaning</span>**

```text id="m4ptrmean"
Method gets POINTER to original struct
```

---

# **<span style="color:#ff1744">5. Methods vs Python Instance/Class Methods</span>**

---

# **<span style="color:#8338ec">Similarity to Python Instance Methods</span>**

---

## **Python</span>**

```python id="m5py"
class User:
    def greet(self):
        pass
```

---

## **Go Equivalent</span>**

```go id="m5go"
func (u User) Greet() {
}
```

---

## **Similarity</span>**

```text id="m5sim"
Both attach behavior to objects
```

---

# **<span style="color:#8338ec">Differences from Python</span>**

---

# **<span style="color:#3a86ff">1. No Classes in Go</span>**

```text id="m5diff1"
Go uses structs + methods
NOT classes
```

---

# **<span style="color:#3a86ff">2. Receiver Explicitly Declared</span>**

```go id="m5diff2"
func (u User)
```

Python:

```python id="m5diff2py"
def greet(self)
```

---

# **<span style="color:#3a86ff">3. No True Class Methods</span>**

```text id="m5diff3"
Go has no static/class methods like Python
```

---

## **Closest Equivalent</span>**

```go id="m5equiv"
func CreateUser() User {
}
```

---

# **<span style="color:#ff1744">6. Value Receiver vs Pointer Receiver (VERY IMPORTANT)</span>**

---

# **<span style="color:#8338ec">A. Value Receiver</span>**

---

## **Example</span>**

```go id="m6a"
type Counter struct {
    Count int
}

func (c Counter) Increment() {
    c.Count++
}
```

---

## **Usage</span>**

```go id="m6ause"
c := Counter{}

c.Increment()

fmt.Println(c.Count)
```

---

## **Output</span>**

```text id="m6aout"
0
```

---

## **Why?</span>**

```text id="m6awhy"
Method received COPY
Original unchanged
```

---

# **<span style="color:#8338ec">B. Pointer Receiver</span>**

---

## **Example</span>**

```go id="m6b"
func (c *Counter) Increment() {
    c.Count++
}
```

---

## **Output</span>**

```text id="m6bout"
1
```

---

## **Why?</span>**

```text id="m6bwhy"
Method modifies ORIGINAL struct
```

---

# **<span style="color:#ff1744">7. Internal Mechanics of Methods</span>**

---

# **<span style="color:#8338ec">Compiler Transformation</span>**

---

## **Method Call</span>**

```go id="m7call"
u.Greet()
```

Internally becomes:

```text id="m7internal"
Greet(u)
```

---

## **Pointer Receiver</span>**

```go id="m7ptr"
u.Update()
```

Internally:

```text id="m7ptrinternal"
Update(&u)
```

---

# **<span style="color:#8338ec">Key Insight</span>**

```text id="m7insight"
Methods are syntactic sugar over functions
with receiver as first parameter
```

---

# **<span style="color:#ff1744">8. Good Real-World Example</span>**

---

# **<span style="color:#8338ec">Bank Account Example</span>**

```go id="m8ex"
package main

import "fmt"

type BankAccount struct {
    Owner   string
    Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
    b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) {

    if amount > b.Balance {
        fmt.Println("Insufficient balance")
        return
    }

    b.Balance -= amount
}

func (b BankAccount) ShowBalance() {
    fmt.Println("Balance:", b.Balance)
}

func main() {

    acc := BankAccount{
        Owner:   "Arj",
        Balance: 1000,
    }

    acc.Deposit(500)

    acc.Withdraw(300)

    acc.ShowBalance()
}
```

---

# **<span style="color:#8338ec">Execution Flow</span>**

```text id="m8flow"
Create struct
↓
Deposit modifies original balance
↓
Withdraw modifies original balance
↓
ShowBalance prints current value
```

---

# **<span style="color:#ff1744">9. Method Sets in Go (Advanced Important Concept)</span>**

---

# **<span style="color:#8338ec">Rule</span>**

```text id="m9rule"
Value type can call:
✔ value receiver methods

Pointer type can call:
✔ value receiver methods
✔ pointer receiver methods
```

---

## **Example</span>**

```go id="m9ex"
u := User{}
u.Update()
```

✔ Go automatically takes address

---

# **<span style="color:#ff1744">10. Use Cases of Methods</span>**

---

# **<span style="color:#8338ec">1. Encapsulation</span>**

```text id="m10a"
Attach behavior to data
```

---

# **<span style="color:#8338ec">2. Data Validation</span>**

```go id="m10b"
func (u *User) Validate()
```

---

# **<span style="color:#8338ec">3. Business Logic</span>**

```go id="m10c"
Deposit()
Withdraw()
```

---

# **<span style="color:#8338ec">4. Interface Implementation</span>**

```go id="m10d"
Area()
String()
Read()
```

---

# **<span style="color:#ff1744">11. Best Practices</span>**

---

## **<span style="color:#3a86ff">1. Use Pointer Receivers for Large Structs</span>**

```text id="m11a"
Avoid expensive copies
```

---

## **<span style="color:#3a86ff">2. Use Pointer Receivers for Mutation</span>**

```text id="m11b"
If method modifies struct
```

---

## **<span style="color:#3a86ff">3. Keep Receiver Names Short</span>**

```go id="m11c"
func (u User)
```

✔ Good

---

## **<span style="color:#3a86ff">4. Be Consistent</span>**

```text id="m11d"
Avoid mixing value/pointer receivers randomly
```

---

## **<span style="color:#3a86ff">5. Methods Should Represent Behavior</span>**

```text id="m11e"
Struct = data
Method = action on data
```

---

# **<span style="color:#ff1744">12. Precautions & Common Mistakes</span>**

---

# **<span style="color:#8338ec">1. Accidentally Copying Structs</span>**

```go id="m12a"
func (u User) Update()
```

❌ changes lost

---

# **<span style="color:#8338ec">2. Inconsistent Receiver Types</span>**

```text id="m12b"
Can confuse interfaces/method sets
```

---

# **<span style="color:#8338ec">3. Large Struct Copies</span>**

```text id="m12c"
Value receiver may hurt performance
```

---

# **<span style="color:#8338ec">4. Treating Go Like OOP Classes</span>**

```text id="m12d"
Go is composition-oriented,
NOT classical inheritance OOP
```

---

# **<span style="color:#ff1744">13. Combined Mental Model</span>**

---

```text id="m13model"
Struct =
data

Method =
behavior operating on data
```

---

## **Execution Flow</span>**

```text id="m13flow"
Struct instance
↓
Method called
↓
Receiver passed internally
↓
Method executes
```

---

# **<span style="color:#ff1744">14. Simple Exercise</span>**

---

## **<span style="color:#3a86ff">Task</span>**

Create:

```text id="m14task"
Struct: Rectangle
Methods:
- Area()
- Scale()
```

---

## **Example</span>**

```go id="m14ex"
type Rectangle struct {
    Width  int
    Height int
}

func (r Rectangle) Area() int {
    return r.Width * r.Height
}

func (r *Rectangle) Scale(x int) {
    r.Width *= x
    r.Height *= x
}
```

---

# **<span style="color:#ff1744">Final Summary</span>**

---

### **Methods**

```text id="m15sum1"
Functions attached to structs/types
```

---

### **Similar to Python?**

```text id="m15sum2"
Yes → like instance methods
BUT Go has no classes
```

---

### **Key Concept**

```text id="m15sum3"
Receiver controls:
- copy behavior
- mutation
- performance
```

---

### **Most Important Rule**

```text id="m15sum4"
Use pointer receivers when modifying structs
```

---

# **<span style="color:#ff1744">Key Insight</span>**

```text id="m15insight"
Go methods bring behavior close to data,
without the complexity of classical OOP systems
```
