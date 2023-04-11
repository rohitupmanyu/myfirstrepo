package main

import (
	"fmt"
)

type rohittype int
var x rohittype

func main() {
	fmt.Println(x)
	fmt.Printf("type is %T\n", x)
	x = 42
	fmt.Printf("%v\n", x)
}

