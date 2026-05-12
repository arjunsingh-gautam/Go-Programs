// LIFO internal deferred stack working
package main

import "fmt"
func main(){
	fmt.Println("Start count...")
	for _,v:=range []int{1,2,3,4,5,6,7,8,9,10}{
		defer fmt.Printf("%d\n",v)
	}
	fmt.Println("Done counting")
}