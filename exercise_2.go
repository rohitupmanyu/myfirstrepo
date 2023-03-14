package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Println(x, y, z) // Printing the zero values
	x = 2
	y = "Hello"
	z = true 
	fmt.Println(x, y, z)
}

