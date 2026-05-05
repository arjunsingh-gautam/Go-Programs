// Demonstrating why use capital case for function name in Go if we want to use it outside the package

package main

import (
	"fmt"
	"go_practice/utils"
)

func main() {
    fmt.Println(utils.Add(2, 3))
    // fmt.Println(utils.subtract(5, 2)) // try uncommenting because in utils package subtract function is not exported because it starts with small case letter
}