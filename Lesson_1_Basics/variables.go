// Variable declaration using var:

package main

import "fmt"

var c,python,java bool // package level scope

func main(){
	var i int // local/functional level scope
	fmt.Println(i,c,python,java) // Automatically initialise to 0 for bool whic false
}