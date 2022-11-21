package main

import "fmt"

func main() {
	hello := "hello"
	world := "world"

	// Prints without line breaks
	fmt.Print(hello)
	fmt.Print(world)

	// Prints with line breaks
	fmt.Println(hello)
	fmt.Println(world)

	// Prints replacing variable values -> Follow the link to get more information about the variables https://pkg.go.dev/fmt
	name := "Erika"
	age := 26

	fmt.Printf("Hello %s, your age is %d \n", name, age) // name replaces the first character and age replace the second one
	fmt.Printf("Hello %v, your age is %v \n", age, name) // age replaces the first character and name replace the second one.

	// Saves into a variable
	messages := fmt.Sprintf("Hello %v, your age is %v", age, name)
	fmt.Println(messages)

	// Print data type of a variable
	fmt.Printf("Name: %T \n", name)

	// Read and capture data from console
	fmt.Print("Please enter your name: ")
	fmt.Scanln(&name) // Capture console data and save into a variable which already exists. It only capture one word data (without any space)
	fmt.Println("El nombre modificado es:", name)
}
