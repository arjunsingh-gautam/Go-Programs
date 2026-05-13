// pointer to struct
package main

import "fmt"

func main(){
	type User struct{
		name string
		age int
	}
	ptr:=&User{
		name:"Sam",age:25,
	}
	fmt.Printf("Name:%v\n",ptr.name)
	fmt.Printf("Age:%d",ptr.age)
}