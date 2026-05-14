// Arrays in Go
package main

import "fmt"

func main(){
	var nums [5]int
	fmt.Printf("Array is %v\n",nums)
	var strings [5]string=[5]string{"Arjun","John","Raj","Ash","Abhi"}
	for _,v:=range strings{
		fmt.Printf("String is %s\n",v)
	}
	var decimals [3]float32=[3]float32{0:4.56,1:7.83}
	fmt.Printf("Decimals:%v\n",decimals)
}