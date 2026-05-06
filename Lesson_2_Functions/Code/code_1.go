// Functions in  Go basic example

package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	result := add(5, 3)
	fmt.Printf("The sum of 5 and 3 is: %d\n", result)
}