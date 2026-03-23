// Basic Go Data Types

/* 1. string
2. bool
3. int,int8,int16,int32,int64
4. rune
5. float32,float64
6. complex64,complex128
7. uint,uint8,uint16,uint32,uint64,uintptr
8. byte */

package main

import (
	"fmt"
)

var(
	b bool=false
	u uint=8
	z complex128=7+8i
	s string="Arjun"
	d float64=7.8282
)

func main(){
	fmt.Printf("Type:%T Value:%v\n",b,b)
	fmt.Printf("Type:%T Value:%v\n",u,u)
	fmt.Printf("Type:%T Value:%v\n",z,z)
	fmt.Printf("Type:%T Value:%v\n",s,s)
	fmt.Printf("Type:%T Value:%v\n",d,d)
}