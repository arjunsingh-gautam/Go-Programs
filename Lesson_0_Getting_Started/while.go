// For is Go's While

// for behaves like C's while statement if we remove semicolon syntax and write like c style while

package main

import "fmt"

func main(){
	sum:=0
	var i int = 1
	for i<10{
		sum+=i
		i++
	}
	fmt.Printf("Sum is:%d",sum)
}