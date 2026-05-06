// In this module we will learn about package in go
package main

import (
	"fmt"
	"math/rand"
)
func main() {
	fmt.Println("My favorite number is:",rand.Intn(100))
}