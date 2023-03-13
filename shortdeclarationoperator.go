package main

import (
	"fmt"
)

func main() {
	x := 42 //declare and assign at the same time
	fmt.Println(x)
	x = 99 // Just assign as the variable is already declared with "short declaration operator"
	fmt.Println(x)
	y := 100 + 24
	fmt.Println(y)
}
