// Declaring variable with initialisers(initial value):
package main

import "fmt"
var i,j int = 1,2
func main() {
	var c,python,java=true,false,"no!" // if an initialiser present no need to specify type while variable declaration the compiler we allocate type by checking initialiser type
	fmt.Println(i,j,c,python,java)
}