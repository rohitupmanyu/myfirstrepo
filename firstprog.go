package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello\n")
	foo()
	cnt, _ := fmt.Println("Some more stuff")

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println(retstring())
	bar()
	fmt.Printf("Printed %d bytes", cnt)
}

func foo() {
	fmt.Println("I am in foo")
}
func bar() {
	fmt.Println("We are done")
}
func retstring() string {
	return "My world"
}
