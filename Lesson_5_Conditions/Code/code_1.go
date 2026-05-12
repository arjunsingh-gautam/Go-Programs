// In this module we will learn about if statement
package main

import "fmt"

func main(){
	var age int
	fmt.Println("Enter your age:")
	fmt.Scan(&age)
	if age>=18{
		fmt.Println("You can vote!")
	}
	if 0<age && age<18{
		fmt.Println("You cannot vote")
	}

}