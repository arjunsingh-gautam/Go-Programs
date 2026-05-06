// In this module we will understand defer statements
// Defer functions are inserted in stack but return value from LIFO(Last in first out) when surrounding function return values
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
