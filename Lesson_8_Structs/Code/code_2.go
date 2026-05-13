// Composition in struct
package main

import (
	"fmt"
)
func main(){
	type Address struct{
		city string
		pincode int
	}
	type User struct{
		name string
		age int
		Address
	}
	u1:=User{}
	u1.name="Arjun"
	u1.age=21
	u1.city="Adampur"
	u1.pincode=411923
	fmt.Printf("User is:%v",u1)
}