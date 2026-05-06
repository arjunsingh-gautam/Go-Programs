// In this module we will learn about shadowing in Go
package main

import "fmt"

func main(){
	x:=5

	if true{
		x:=x+10 //shadowing
		fmt.Printf("Value of x inner is:%v\n",x)

	}
	fmt.Printf("Value of x outter:%v",x)
}