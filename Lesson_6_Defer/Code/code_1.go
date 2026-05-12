// Defer statement
package main

import "fmt"
func main(){
	defer fmt.Println("Hello") // deferred execution executed after are neighbouring functions return 
	// defer statements executed after main() return defer statements start poping from deferred stack
	fmt.Println("World")
}