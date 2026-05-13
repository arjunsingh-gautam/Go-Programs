// Basic of Pointers in Go
package main

import "fmt"

func main(){
	x:=39
	p:=&x
	fmt.Printf("Address of %d is:%p\n",x,p)
}