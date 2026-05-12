// if with initialisation
package main

import (
	"fmt"
	"math/rand"
)

func main(){
	if num:=rand.Intn(100);num%2==0{
		fmt.Printf("%d is even",num)
	}else{
		fmt.Printf("%d is odd",num)
	}
}