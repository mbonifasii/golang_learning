package main

import "fmt"

func main() {

	// Declaring constances
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// Declaring variables
	base := 12
	var altura int = 14
	var area int

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)
	fmt.Println(base, altura, area)

	// Square area
	const squareBase = 10
	squareArea := squareBase * squareBase

	fmt.Println(squareArea)
}
