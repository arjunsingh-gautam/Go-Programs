// Pointer to Pointer in Go Higher order pointer in Go
package main

import "fmt"

func main(){
	x:=89
	p:=&x
	pp:=&p
	ppp:=&pp
	fmt.Printf("Addres of x:%p\n",p)
	fmt.Printf("Address of p:%p\n",pp)
	fmt.Printf("Address of pp:%p\n",ppp)
	fmt.Printf("Value of x is:%d,%d,%d,%d\n",x,*p,**pp,***ppp)
}

