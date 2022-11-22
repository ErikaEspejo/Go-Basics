package main

import (
	"fmt"
	"math"
)

func main() {
	a := 20
	b := 10

	//SUM
	result := a + b // := allows you to declare and initilize at the same time
	fmt.Println("Sum:", result)

	// SUBSTRACTION
	result = a - b // It's used "=" because the variable is already declared, so it is only being reused.
	fmt.Println("Substraction:", result)

	// MULTIPLICATION
	result = a * b
	fmt.Println("Multiplication", result)

	// DIVISION
	// The only way to get an exact division is with data of float type (Ex: 2.0, 1.5,...). Otherwise you just get an integer division.
	result = a / b
	fmt.Println("Integer Division:", result)

	result2 := 2 / 3
	fmt.Println("Integer Division:", result2)

	result3 := 2.0 / 3.0
	fmt.Println("Exact Division:", result3)

	// MODULE
	result = a % b
	fmt.Println("Module:", result)

	result = 5 % 2
	fmt.Println("Module:", result)

	// INCREMENTAL
	a++
	fmt.Println("Incremental:", a)

	// DECREMENTAL
	a--
	fmt.Println("Decremental:", a)

	// Challenge -> Calculate area of circle, square, rectangle and trapezium
	fmt.Println("\n---- CHALLENGES ----")
	// SQUARE
	fmt.Println("-- Square --")
	side := 10
	squareArea := side * side
	fmt.Println("Side:", side)
	fmt.Println("Square Area:", squareArea)

	// CIRCLE
	fmt.Println("-- Circle --")
	radius := 10
	circleArea := math.Phi * math.Pow(float64(radius), 2)
	fmt.Println("Radius:", radius)
	fmt.Println("Circle Area:", circleArea)

	// RECTANGLE
	fmt.Println("-- Rectangle --")
	base := 10
	height := 20
	rectangleArea := base * height
	fmt.Println("Base:", base)
	fmt.Println("Height:", height)
	fmt.Println("Rectangle Area:", rectangleArea)

	// TRAPEZIUM
	fmt.Println("-- Trapezium --")
	minorBase := 10
	majorBase := 20
	height = 20
	trapeziumArea := (minorBase + majorBase) * height / 2
	fmt.Println("Minor Base:", minorBase)
	fmt.Println("Major Base:", majorBase)
	fmt.Println("Height:", height)
	fmt.Println("Trapezium Area:", trapeziumArea)

}
