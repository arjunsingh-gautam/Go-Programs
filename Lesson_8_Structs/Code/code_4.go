// object bound methods
package main

import "fmt"

// Can define object bound methods at package scope only
type Rectangle struct{
		Length float64
		Breadth float64
	}

func  (ptr *Rectangle) Area() float64 {
		return ptr.Length*ptr.Breadth
	}

func main(){
	

	r:=Rectangle{Length:3.5,Breadth:4.2}
	p:=&r
	fmt.Printf("Area of rectangle:%.2f",p.Area())
}