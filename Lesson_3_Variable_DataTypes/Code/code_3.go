// Data Types in Go
package main

import "fmt"

type User struct{
	Name string
	Age string

}

func main(){
	var u User=User{Name:"John",Age:"30"}
	fmt.Printf("User details: %v\n",u)
	s :=[]int{1,2,3,4,5}
	fmt.Printf("Slice details: %v\n",s)
	m := map[string]int{"one":1,"two":2,"three":3}
	fmt.Printf("Map details: %v\n",m)
	

}