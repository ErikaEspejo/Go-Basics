package main

import (
	"fmt"
)

// Function without parameters and return
func helloWorld() {
	fmt.Println("Hello World")
}

// Function without return, and with parameters
func greeting(name string) {
	fmt.Println("I'm", name)
}

// Function with parameters and one return.
// In a function with return, a return data type is required. It is written after the parameters and before the body of the function
func sum(a, b int) int { // When there is more than one parameter of the same data type, you can write the variables and then the data type.
	return a + b
}

// Function with parameters and more than one return
func doubleReturn(a, b int) (c, d int) { // c an d are the return variables in memory
	return a * b, a / b
}

// MAIN FUNCTION
func main() {

	// Call and execture the functions
	helloWorld()
	greeting("Erika")
	result := sum(10, 20) // Save the return data into a variable
	fmt.Println(result)

	multiplication, division := doubleReturn(6, 2) // Go requires to use all the data, so if you get two returns you need to store both into different variables
	fmt.Println("Multiplication:", multiplication)
	fmt.Println("Division:", division)

	// When you only need to use one of the returns and ignore the other ones, you could use the "_" character to escape.
	// It should be in the place of the return you need to ignore.
	multplication2, _ := doubleReturn(5, 3)
	fmt.Println("Multiplication Without Division:", multplication2)

	_, division2 := doubleReturn(6, 2)
	fmt.Println("Division without Multiplication:", division2)
}
