package main

import "fmt"

func main() {

	// How to declare variables in Go
	// var name type = expression
	var a int = 1
	// default value is assigned to 0
	var b int

	var c, d int
	// declaration without types
	var e, f = 1, 2

	var integer, string = 1, "string"

	// short declaration:
	// name := expression
	bool := false

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c, d)
	fmt.Println(e, f)
	fmt.Println(integer, string)
	fmt.Println(bool)
	fmt.Println("===================================")

	// ** Pointers **
	x := 1
	p := &x // & = address of expression

	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println("===================================")

	// declaration of types
	// ** Types **
	// fahrenheit & celsius
	type fahrenheit int
	type celsius int

	var fa fahrenheit = 32
	var ce celsius = 0

	ce = celsius((fa - 32) * 5 / 9)

	fmt.Println(fa, ce)

}
