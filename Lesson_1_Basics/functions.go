// Functions in go

package main

import "fmt"

func add(x int,y int)int { // function declaration similar to C
	return x+y
}

func sub(x,y int)int { // we can shorten x int,y int -> x,y int if consecutive arguements contain same type omit type for all such consecutive arguement but the last
	return x-y
}

// Function can return multiple values
func swap(x,y string) (string,string) {
	return y,x // return multiple values
}

// Name return values: return variables define at top of the function using only return(naked)
func split(sum int)(x,y int){
	x=sum*4/9
	y=sum -x
	return	
}

func main(){
	fmt.Printf("Sum is:%d\n",add(43,12))
	fmt.Printf("Difference is:%d\n",sub(23,12))
	a,b:=swap("Hello","World")
	fmt.Printf("a:%s\nb:%s\n",a,b)
	fmt.Println(split(17))
}