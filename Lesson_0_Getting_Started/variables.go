// Variable declaration using var:

package main

import "fmt"

//a:=4 syntax error: Outside function every statement begins with keyword
var c,python,java bool // package level scope

func main(){
	a:=7.97 // short assignment statement
	var i int // local/functional level scope
	i=7
	fmt.Println(i,c,python,java,a) // Automatically initialise to 0 for bool which false
}