// switch without expression act as cleaner if-else if-else ladder
package main

import "fmt"

func main(){
	age:=19
	switch{
	case age>=18:
		fmt.Println("Can Vote!")
	case age<18:
		fmt.Println("Cannot Vote!")
	}
}