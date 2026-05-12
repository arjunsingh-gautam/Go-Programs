// continue in go
package main

import "fmt"

func main(){
	for _,v := range []int{1,2,3,4,5,6,7,8,9,10}{
		if v%2==0{
			continue
		}
		fmt.Printf("%d\n",v)
	}
}