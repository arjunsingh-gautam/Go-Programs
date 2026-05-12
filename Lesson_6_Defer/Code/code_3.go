// context is registered when defer statement is first seen
package main

import "fmt"
func main(){
	x:="Deferred"
	defer fmt.Println(x)
	x="Not Deferred"
	fmt.Println(x)
}