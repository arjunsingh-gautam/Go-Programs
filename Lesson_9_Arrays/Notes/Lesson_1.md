# **<span style="color:#ff1744">Arrays in Go — Complete Mental Model</span>**

---

# **<span style="color:#ff6f00">1. What are Arrays?</span>**

---

## **<span style="color:#3a86ff">Definition</span>**

```text id="arr1def"
An array is a fixed-size ordered collection
of elements of the SAME type
```

---

## **<span style="color:#3a86ff">Core Idea</span>**

```text id="arr1core"
Array =
continuous block of memory
containing same-type values
```

---

## **<span style="color:#8338ec">Example</span>**

```go id="arr1ex"
var nums [5]int
```

Meaning:

```text id="arr1meaning"
Array of 5 integers
```

---

## **<span style="color:#8338ec">Analogy</span>**

```text id="arr1ana"
Array = row of fixed lockers

Index:
0 1 2 3 4
```

---

# **<span style="color:#ff1744">2. Important Properties of Arrays in Go</span>**

---

## **<span style="color:#3a86ff">Properties</span>**

```text id="arr2props"
✔ Fixed size
✔ Same data type
✔ Contiguous memory
✔ Index-based access
✔ Value type
```

---

## **<span style="color:#8338ec">Very Important</span>**

```text id="arr2important"
Array size is PART OF THE TYPE
```

---

## **Example</span>**

```go id="arr2type"
[5]int ≠ [10]int
```

These are different types

---

# **<span style="color:#ff1744">3. Syntax of Arrays in Go</span>**

---

# **<span style="color:#8338ec">General Syntax</span>**

```go id="arr3syntax"
var arrayName [size]Type
```

---

# **<span style="color:#ff1744">4. Ways to Declare, Define & Initialize Arrays</span>**

---

# **<span style="color:#8338ec">A. Declaration Only</span>**

```go id="arr4a"
var nums [5]int
```

---

## **Result</span>**

```text id="arr4aout"
[0 0 0 0 0]
```

---

## **Why?</span>**

```text id="arr4awhy"
Arrays get zero values automatically
```

---

# **<span style="color:#8338ec">B. Full Initialization</span>**

```go id="arr4b"
var nums [5]int = [5]int{1, 2, 3, 4, 5}
```

---

# **<span style="color:#8338ec">C. Short Syntax</span>**

```go id="arr4c"
nums := [5]int{1, 2, 3, 4, 5}
```

---

# **<span style="color:#8338ec">D. Compiler-Inferred Size</span>**

```go id="arr4d"
nums := [...]int{1, 2, 3}
```

---

## **Meaning</span>**

```text id="arr4dmean"
Compiler counts size automatically
```

---

# **<span style="color:#8338ec">E. Partial Initialization</span>**

```go id="arr4e"
nums := [5]int{1, 2}
```

---

## **Output</span>**

```text id="arr4eout"
[1 2 0 0 0]
```

---

# **<span style="color:#8338ec">F. Indexed Initialization</span>**

```go id="arr4f"
nums := [5]int{
    0: 10,
    3: 40,
}
```

---

## **Output</span>**

```text id="arr4fout"
[10 0 0 40 0]
```

---

# **<span style="color:#ff1744">5. Accessing & Modifying Arrays</span>**

---

# **<span style="color:#8338ec">Access Elements</span>**

```go id="arr5a"
fmt.Println(nums[0])
```

---

# **<span style="color:#8338ec">Modify Elements</span>**

```go id="arr5b"
nums[1] = 100
```

---

# **<span style="color:#8338ec">Index Rules</span>**

```text id="arr5rules"
Index starts at 0
Valid range:
0 → len(array)-1
```

---

# **<span style="color:#ff1744">6. Operations on Arrays</span>**

---

# **<span style="color:#8338ec">A. Iteration</span>**

```go id="arr6a"
for i := 0; i < len(nums); i++ {
    fmt.Println(nums[i])
}
```

---

# **<span style="color:#8338ec">B. range Loop</span>**

```go id="arr6b"
for i, v := range nums {
    fmt.Println(i, v)
}
```

---

# **<span style="color:#8338ec">C. Length</span>**

```go id="arr6c"
len(nums)
```

---

# **<span style="color:#8338ec">D. Comparison</span>**

```go id="arr6d"
a := [3]int{1,2,3}
b := [3]int{1,2,3}

fmt.Println(a == b)
```

✔ Allowed if same type

---

# **<span style="color:#8338ec">E. Copying</span>**

```go id="arr6e"
b := a
```

---

## **Important</span>**

```text id="arr6eimportant"
Entire array gets copied
```

---

# **<span style="color:#ff1744">7. Internal Memory Mechanics of Arrays</span>**

---

# **<span style="color:#8338ec">Contiguous Memory Layout</span>**

```text id="arr7layout"
Elements stored sequentially in memory
```

---

## **Example</span>**

```go id="arr7ex"
nums := [3]int{10,20,30}
```

---

## **Visualization</span>**

```text id="arr7vis"
Address:
1000 → 10
1008 → 20
1016 → 30
```

---

# **<span style="color:#8338ec">Why Arrays are Fast</span>**

```text id="arr7fast"
Index calculation is direct:
base + offset
```

---

# **<span style="color:#ff1744">8. Arrays are Value Types (VERY IMPORTANT)</span>**

---

# **<span style="color:#8338ec">Example</span>**

```go id="arr8ex"
a := [3]int{1,2,3}

b := a

b[0] = 100
```

---

## **Output</span>**

```text id="arr8out"
a = [1 2 3]
b = [100 2 3]
```

---

## **Why?</span>**

```text id="arr8why"
Assignment COPIES entire array
```

---

# **<span style="color:#ff1744">9. Does Go Array Decay to Pointer Like C?</span>**

---

# **<span style="color:#8338ec">C Behavior</span>**

```text id="arr9c"
Array name decays into pointer
```

---

## **Example in C</span>**

```c id="arr9cex"
int arr[5];

arr → pointer to first element
```

---

# **<span style="color:#8338ec">Go Behavior</span>**

```text id="arr9go"
NO automatic array-to-pointer decay
```

---

## **Important Difference</span>**

```text id="arr9diff"
Go arrays remain full VALUE TYPES
```

---

# **<span style="color:#8338ec">Example</span>**

```go id="arr9goex"
func process(a [3]int) {
}
```

Passing array:

```text id="arr9copy"
Entire array copied
```

---

# **<span style="color:#8338ec">Pointer to Array</span>**

```go id="arr9ptr"
func process(a *[3]int) {
}
```

✔ Explicit pointer needed

---

# **<span style="color:#ff1744">10. Arrays vs Slices</span>**

---

| Feature            | Array       | Slice            |
| ------------------ | ----------- | ---------------- |
| Size               | Fixed       | Dynamic          |
| Type includes size | Yes         | No               |
| Assignment         | Copies data | Shares reference |
| Flexibility        | Low         | High             |

---

# **<span style="color:#8338ec">Important Insight</span>**

```text id="arr10insight"
Slices are built ON TOP OF arrays
```

---

# **<span style="color:#ff1744">11. Constraints of Arrays</span>**

---

# **<span style="color:#8338ec">1. Fixed Size</span>**

```text id="arr11a"
Cannot resize after creation
```

---

# **<span style="color:#8338ec">2. Expensive Copies</span>**

```text id="arr11b"
Large arrays copied entirely
```

---

# **<span style="color:#8338ec">3. Less Flexible</span>**

```text id="arr11c"
Poor for dynamic collections
```

---

# **<span style="color:#8338ec">4. Rarely Used Directly</span>**

```text id="arr11d"
Slices preferred in most Go programs
```

---

# **<span style="color:#ff1744">12. Use Cases of Arrays</span>**

---

# **<span style="color:#8338ec">1. Fixed-Size Data</span>**

```text id="arr12a"
RGB colors
Days of week
Coordinates
```

---

# **<span style="color:#8338ec">2. Performance-Critical Systems</span>**

```text id="arr12b"
Avoid allocations
Predictable memory layout
```

---

# **<span style="color:#8338ec">3. Underlying Storage for Slices</span>**

```text id="arr12c"
Slices internally use arrays
```

---

# **<span style="color:#8338ec">4. Embedded/Low-Level Programming</span>**

```text id="arr12d"
Fixed memory control
```

---

# **<span style="color:#ff1744">13. Best Practices</span>**

---

## **<span style="color:#3a86ff">1. Prefer Slices for General Usage</span>**

```text id="arr13a"
Slices are more flexible
```

---

## **<span style="color:#3a86ff">2. Use Arrays Only for Fixed-Size Data</span>**

```text id="arr13b"
When size is truly constant
```

---

## **<span style="color:#3a86ff">3. Pass Large Arrays by Pointer</span>**

```go id="arr13c"
func process(a *[1000]int)
```

✔ avoids copying

---

## **<span style="color:#3a86ff">4. Use range for Readability</span>**

```go id="arr13d"
for _, v := range arr
```

---

# **<span style="color:#ff1744">14. Precautions & Common Mistakes</span>**

---

# **<span style="color:#8338ec">1. Accidental Array Copies</span>**

```go id="arr14a"
b := a
```

❌ copies whole array

---

# **<span style="color:#8338ec">2. Out-of-Bounds Access</span>**

```go id="arr14b"
arr[10]
```

❌ runtime panic

---

# **<span style="color:#8338ec">3. Confusing Arrays with Slices</span>**

```go id="arr14c"
[5]int ≠ []int
```

---

# **<span style="color:#8338ec">4. Using Huge Arrays on Stack</span>**

```text id="arr14d"
Can increase stack memory usage
```

---

# **<span style="color:#ff1744">15. Combined Mental Model</span>**

---

```text id="arr15model"
Array =
fixed-size contiguous memory block
```

---

## **Lifecycle</span>**

```text id="arr15flow"
Declare size
↓
Allocate continuous memory
↓
Store values sequentially
↓
Access via index math
```

---

# **<span style="color:#ff1744">16. Simple Exercise</span>**

---

## **<span style="color:#3a86ff">Task</span>**

```go id="arr16task"
arr := [5]int{1,2,3,4,5}

copyArr := arr

copyArr[0] = 100

fmt.Println(arr)
fmt.Println(copyArr)
```

---

## **Goal</span>**

```text id="arr16goal"
Observe value-copy behavior
```

---

# **<span style="color:#ff1744">Final Summary</span>**

---

### **Arrays**

```text id="arr17sum1"
Fixed-size homogeneous collections
```

---

### **Key Properties**

```text id="arr17sum2"
Contiguous memory
Value type
Fast indexed access
```

---

### **Major Difference from C**

```text id="arr17sum3"
Go arrays DO NOT decay into pointers
```

---

### **Main Limitation**

```text id="arr17sum4"
Fixed size and costly copying
```

---

# **<span style="color:#ff1744">Key Insight</span>**

```text id="arr17insight"
Arrays are low-level predictable memory structures,
while slices provide Go’s practical dynamic collection abstraction
```
