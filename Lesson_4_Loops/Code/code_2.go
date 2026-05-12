// In this module we will learn about while loop syntax in go
package main

import "fmt"

func main(){
	sum:=0
	i:=0
	for i<14{ // while loop syntax
		sum+=i
		i++
	}
	fmt.Printf("Sum is:%d",sum)
}