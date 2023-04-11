package main

import (
	"fmt"
)

type rohittype int
var x rohittype
var y int
func main() {
	fmt.Println(x)
	fmt.Printf("type is %T\n", x)
	x = 42
	fmt.Printf("%v\n", x)
	y = int(x)
	fmt.Printf("%v", y)
	fmt.Printf("%T", y)
}

