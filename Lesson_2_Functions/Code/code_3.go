// Definition of a functions where return are initialized with zero values of return type
package main

import "fmt"

func split(value int) (x int, y int) {
	x = value * 4 / 9
	y = value - x
	return
}

func main() {
	fmt.Println(split(17))
}