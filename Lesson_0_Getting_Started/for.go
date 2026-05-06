// For Iterative statement:

package main

import "fmt"

func main(){
	sum:=0
	for i:=0;i<10;i++{
		sum+=i
	}
	sum2:=1
	for ;sum2<3; { // for sum2=1;sum2<3;sum2++
		sum2+=sum2
	}
	fmt.Printf("Sum is:%d",sum)
}