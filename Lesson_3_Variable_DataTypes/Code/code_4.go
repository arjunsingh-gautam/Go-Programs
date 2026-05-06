// Type conversion in Go
package main

import "fmt"

func main(){
	var x int=5
	var y float64=float64(x)
	fmt.Printf("Value of x is:%v and type of x is:%T\n",x,x)
	fmt.Printf("Value of y is:%v and type of y is:%T\n",y,y)
}