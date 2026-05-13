// Basic structure in Go
package main

import "fmt"
func main(){
	type User struct{
		name string
		age int
	}
	u1:=User{"Arjun",19}
	u2:=User{name:"Ash",age:23}
	var u3 User;
	fmt.Printf("%v\n",u1)
	fmt.Printf("%v\n",u2)
	fmt.Printf("%v\n",u3)
	u1.name="Sara"
	fmt.Printf("%v\n",u1)
}