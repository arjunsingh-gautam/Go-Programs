// Type conversion :
package main

import (
	"fmt"
)

func main() {
	var x,y float64=3.14,4.24
	var f int =int(x*x+y*y)
	i:=uint(f)
	fmt.Println(x,y,f,i)

}