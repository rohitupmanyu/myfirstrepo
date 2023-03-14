package main
import "fmt"

var y int = -42  // y is an int explicitly and default
var z = "Time pass" // z becomes a string by default

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	z = "Time pass karo" 
	//z = 42 wont work as types wont match
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	foo()
}

func foo() {
	fmt.Printf(z)
}
