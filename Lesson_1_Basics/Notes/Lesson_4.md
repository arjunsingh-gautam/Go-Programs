# **<span style="color:#ff1744">`go build` vs `go run` — Complete Mental Model</span>**

---

# **<span style="color:#ff6f00">1. Short Definition</span>**

---

## **<span style="color:#3a86ff">go run</span>**

```text
Compiles + runs your program immediately (temporary execution)
```

---

## **<span style="color:#3a86ff">go build</span>**

```text
Compiles your program into a binary executable (no execution)
```

---

# **<span style="color:#ff1744">2. Core Difference</span>**

---

| Command  | What it does                       |
| -------- | ---------------------------------- |
| go run   | Compile → execute → discard binary |
| go build | Compile → create binary file       |

---

# **<span style="color:#ff1744">3. Example</span>**

---

## **<span style="color:#3a86ff">Code</span>**

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello")
}
```

---

## **<span style="color:#3a86ff">Using go run</span>**

```bash
go run main.go
```

Output:

```text
Hello
```

No binary created

---

## **<span style="color:#3a86ff">Using go build</span>**

```bash
go build
```

Creates:

```text
main.exe   (Windows)
./main     (Linux/Mac)
```

Run manually:

```bash
./main
```

---

# **<span style="color:#ff1744">4. Internal Mechanics (VERY IMPORTANT)</span>**

---

# **<span style="color:#8338ec">A. go run — Execution Flow</span>**

```text
1. Read source files
2. Resolve imports (go.mod)
3. Compile code → temporary binary
4. Execute binary
5. Delete temporary binary
```

---

## **<span style="color:#3a86ff">Key Property</span>**

```text
No persistent output
```

---

# **<span style="color:#8338ec">B. go build — Execution Flow</span>**

```text
1. Read source files
2. Resolve imports (go.mod)
3. Compile code
4. Link dependencies
5. Generate binary executable
6. Store on disk
```

---

## **<span style="color:#3a86ff">Key Property</span>**

```text
Produces reusable executable
```

---

# **<span style="color:#ff1744">5. Analogy</span>**

---

## **<span style="color:#3a86ff">go run</span>**

```text
Cooking instant noodles:
Cook → eat → nothing stored
```

---

## **<span style="color:#3a86ff">go build</span>**

```text
Cooking meal and storing it:
Cook → store → reuse later
```

---

# **<span style="color:#ff1744">6. Performance Insight</span>**

---

## **<span style="color:#3a86ff">go run</span>**

```text
Slower for repeated execution
(because it compiles every time)
```

---

## **<span style="color:#3a86ff">go build</span>**

```text
Compile once → run many times
(faster overall)
```

---

# **<span style="color:#ff1744">7. When to Use Which</span>**

---

## **<span style="color:#3a86ff">Use go run When</span>**

```text
✔ Quick testing
✔ Small scripts
✔ Learning/debugging
✔ One-time execution
```

---

## **<span style="color:#3a86ff">Use go build When</span>**

```text
✔ Production code
✔ Deployment
✔ CLI tools
✔ Performance-sensitive usage
```

---

# **<span style="color:#ff1744">8. Important Flags & Variations</span>**

---

## **<span style="color:#3a86ff">Custom binary name</span>**

```bash
go build -o app.exe
```

---

## **<span style="color:#3a86ff">Run multiple files</span>**

```bash
go run main.go utils.go
```

---

## **<span style="color:#3a86ff">Build specific package</span>**

```bash
go build ./...
```

---

# **<span style="color:#ff1744">9. Key Internal Difference (Critical)</span>**

---

## **<span style="color:#8338ec">Temporary Binary Location (go run)</span>**

```text
Stored in system temp directory
Deleted after execution
```

---

## **<span style="color:#8338ec">Persistent Binary (go build)</span>**

```text
Stored in your current directory
```

---

# **<span style="color:#ff1744">10. Combined Mental Model</span>**

---

### **go run**

```text
Write → Compile → Run → Discard
```

---

### **go build**

```text
Write → Compile → Save → Run anytime
```

---

# **<span style="color:#ff1744">11. Common Mistakes</span>**

---

## **<span style="color:#3a86ff">Mistake 1</span>**

```text
Expecting go run to create binary
```

---

## **<span style="color:#3a86ff">Mistake 2</span>**

```text
Running go build but forgetting to execute binary
```

---

## **<span style="color:#3a86ff">Mistake 3</span>**

```text
Using go run in production repeatedly
```

---

# **<span style="color:#ff1744">12. Simple Exercise</span>**

---

## **<span style="color:#3a86ff">Goal</span>**

```text
Observe difference in behavior
```

---

## **<span style="color:#3a86ff">Steps</span>**

---

### **1. Create file**

```go
package main

import "fmt"

func main() {
    fmt.Println("Testing build vs run")
}
```

---

### **2. Run using go run**

```bash
go run main.go
```

✔ Observe:

```text
- Output appears
- No executable file created
```

---

### **3. Run using go build**

```bash
go build
```

✔ Observe:

```text
- Executable file created
```

---

### **4. Run binary**

```bash
./main   (or main.exe)
```

---

## **<span style="color:#8338ec">Learning Outcome</span>**

```text
You clearly see:
go run = temporary execution
go build = persistent executable
```

---

# **<span style="color:#ff1744">Final Summary</span>**

---

### **go run**

```text
Fast for development
No binary output
Compiles every time
```

---

### **go build**

```text
Creates executable
Reusable
Used in production
```

---

# **<span style="color:#ff1744">Key Insight</span>**

```text
go run = convenience
go build = control + performance
```

---

If you want deeper system-level understanding next:

```text
- "Go compilation pipeline (lexer → parser → SSA → binary)"
- "Static vs dynamic linking in Go"
- "How Go produces single static binaries"
```
