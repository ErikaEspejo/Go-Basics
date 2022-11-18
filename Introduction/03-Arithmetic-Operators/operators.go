package main

import "fmt"

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

}
