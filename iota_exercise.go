package main

import "fmt"

const a = 42
const b = 43.4
const c = "Rohit Upmanyu"

const (
	d = iota
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T", a)
	fmt.Printf(" %T", b)
	fmt.Printf(" %T\n", c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T", d)
	fmt.Printf(" %T", e)
	fmt.Printf(" %T\n", f)
}
