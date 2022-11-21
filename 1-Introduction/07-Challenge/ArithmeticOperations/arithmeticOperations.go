package main

import "fmt"

/*
1. Statement: Given two integers, find the sum. To solve this problem, the user is required to enter two integers and the system
performs the respective calculation to find the sum.

2. Statement: Find the quotient and remainder of two integers. For the solution of this problem, the user is required to enter
two integers and the system performs the respective calculation to find the quotient and remainder.
*/

func sum(a, b int) int {
	return a + b
}

func quotient(a, b int) float32 {
	return float32(a) / float32(b)
}

func module(a, b int) int {
	return a % b
}

func arithmeticOperations(a, b int) {
	fmt.Print("Please enter a number:")
	fmt.Scanln(&a)
	fmt.Print("Please enter another number:")
	fmt.Scanln(&b)

	fmt.Println("Numbers -> x:", a, "; y:", b)
	fmt.Println("Sum:", sum(a, b))
	fmt.Println("Quotient:", quotient(a, b))
	fmt.Println("Module:", module(a, b))
}

func main() {
	var a, b int
	arithmeticOperations(a, b)
}
