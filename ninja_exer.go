package main

import "fmt"

var y = 32
var rohi rune = 5

const (
	aa      = 22
	bb int8 = 8
	cc      = iota
	dd      = iota
)

func main() {
	x := 42
	fmt.Printf("%d\t%b\t%x\ty=%v\n", x, x, x, y)
	a := (32 <= 42)
	fmt.Printf("(32 <= 42) %v\n", a)
	a = (32 >= 42)
	fmt.Printf("(32 >= 42) %v\n", a)
	a = (32 != 42)
	fmt.Printf("(32 != 42) %v\n", a)
	a = (32 < 42)
	fmt.Printf("(32 < 42) %v\n", a)
	a = (32 > 42)
	fmt.Printf("%v\n", aa)
	fmt.Printf("%b\n", bb)
	fmt.Printf("%b\n", bb<<1)
	ss := `Here you go hey 
	h
	ey bl ah blah`
	fmt.Println(ss)
	fmt.Println(cc)
	fmt.Println(rohi)
	fmt.Printf("%T\n", rohi)

}
