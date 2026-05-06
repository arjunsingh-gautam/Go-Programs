// In this module we will understand variable declaration using var keyword and data types in go
package main

import "fmt"

var global_var1 int // Initialised with 0 value
var global_var2 int = 4 
var global_var3=8.75 // Data-type inferred as float and assigned once


func main() {

	var local_var1 bool //Initialise 0 value i.e False
	var local_var2="Arjun"// inferred as string using initialiser
	local_var3:=3.14 // Can be used inside a functon only both inferred and initialised at same time

	fmt.Printf("Type of variable:%T and value stored at variable:%v \n",global_var1,global_var1)
	fmt.Printf("Type of variable:%T and value stored at variable:%v \n",global_var2,global_var2)
	fmt.Printf("Type of variable:%T and value stored at variable:%v \n",global_var3,global_var3)
	fmt.Printf("Type of variable:%T and value stored at variable:%v \n",local_var1,local_var1)
	fmt.Printf("Type of variable:%T and value stored at variable:%v \n",local_var2,local_var2)
	fmt.Printf("Type of variable:%T and value stored at variable:%v \n",local_var3,local_var3)
}