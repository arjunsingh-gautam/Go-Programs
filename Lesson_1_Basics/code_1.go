// importing external package and managing it using go mod

package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
    color.Green("Hello, Go Modules!")
    fmt.Println("Done")
}