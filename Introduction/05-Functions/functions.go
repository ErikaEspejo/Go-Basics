package main

import "fmt"

// Function without parameters and return
func helloWorld() {
	fmt.Println("Hello World")
}

// Function without return, and with parameters
func greeting(name string) {
	fmt.Println("I'm", name)
}

// Function with parameters and return
// In a function with return, a return data type is required. It is written after the parameters and before the body of the function
func sum(a, b int) int { // When there is more than one parameter of the same data type, you can write the variables and then the data type.
	return a + b
}

// MAIN FUNCTION
func main() {

	// Call and execture the functions
	helloWorld()
	greeting("Erika")
	result := sum(10, 20) // Save the return data into a variable
	fmt.Println(result)
}
