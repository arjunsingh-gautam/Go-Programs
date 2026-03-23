// In shortened variable the type is infered by RHS expression

package main

import "fmt"

func main() {
	v:="Arjun"
	fmt.Printf("%v is of type %T",v,v)
}