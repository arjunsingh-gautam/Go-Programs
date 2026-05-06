// Returning multiple values from a function in Go
package main

import "fmt"

func divideAndRemainder(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {
	a := 10
	b := 3	
	quotient, remainder := divideAndRemainder(a, b)
	fmt.Printf("When dividing %d by %d, quotient is %d and remainder is %d\n", a, b, quotient, remainder)
}