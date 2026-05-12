// In this module we will learn about if-else statement
package main

import "fmt"

func main(){
	var marks int
	fmt.Println("Enter your marks:")
	fmt.Scan(&marks)
	if marks>90 && marks<=100{
		fmt.Println("Grade-A")
	}else if marks>80 && marks<=90{
		fmt.Println("Grade-B")
	}else if marks>70 && marks<=80{
		fmt.Println("Grade-C")
	}else{
		fmt.Println("Fail")
	} 

}