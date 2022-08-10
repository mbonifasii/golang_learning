package main

import "fmt"

func main() {

	// Declaring constances
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// Declaring variables
	// base := 12
	// var altura int = 14
	// var area int

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)
	// fmt.Println(base, altura, area)

	// Square area
	const squareBase = 10
	squareArea := squareBase * squareBase

	fmt.Println(squareArea)
	// ------

	x := 10
	y := 50

	result := x + y
	fmt.Println("Suma:", result)

	result = y - x
	fmt.Println("Resta: ", result)

	result = x * y
	fmt.Println("Multiplicacion: ", result)

	result = y / x
	fmt.Println("Division: ", result)

	// modulo
	result = y % x
	fmt.Println("Modulo:", result)

	// Incrementar
	x++
	fmt.Println("Incremental: ", x)

	// Decremental
	x--
	fmt.Println("Decremental: ", x)

	// Retos
	// - area cuadrado
	// - area trapecio
	// - area circulo

	// Area cuadrado
	var width float64 = 10
	var heigth float64 = 20
	var area float64

	area = heigth * width
	fmt.Println("Area de rectangulo: ", area)

	// Area trapecio
	var mayorBase float64
	var minorBase float64

	var trapecioArea = ((mayorBase + minorBase) / 2) * heigth
	fmt.Println("Area de trapecio: ", trapecioArea)

	// Area circulo
	var radio float64 = 2.2
	var areaCircle float64 = pi * (radio * radio)

	fmt.Println("Area de circulo: ", areaCircle)
}
