// In this module we will learn about for loop
package main

import "fmt"

func main(){
	sum:=0
	for i:=0;i<10;i++{
		sum+=i
	}
	fmt.Printf("Sum is:%d",sum)
}