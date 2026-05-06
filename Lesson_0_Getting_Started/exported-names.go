// Exported names of package

/* We can only use the package variable outside the package which are exported.
Exported names of a package start with capital case */

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}