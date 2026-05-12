// Infinite Loop:
package main

import "fmt"

func main(){
	sum:=0
	i:=0
	for{
		sum+=i
		if sum>=100{
			break
		}
		i++
	}
	fmt.Printf("Sum is:%d",sum)
}
