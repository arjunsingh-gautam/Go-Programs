// If statement in Go

package main

import (
	"fmt"
)

func isEven(a int) bool {
	if a%2==0{
		return true
	} else{
		return false
	}

}

func main() {
	var num int
	fmt.Printf("Enter a number:")
	fmt.Scanf("%d",&num)
	fmt.Printf("%v is even:%v",num,isEven(num))
}