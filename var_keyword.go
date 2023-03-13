package main

import (
	"fmt"
)
var y int = 43
func main() {
	// declare a variable and assign a value
	// short declaration operator
	x := 42
	fmt.Println(x)
	// Using 'var' keyword this time
	var y = 43
	fmt.Printf("Local y is %d\n",y);
	foo()
}

func foo() {
	y = 11
	fmt.Printf("Global y is %d\n",y);
}
