package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	value1 := 1
	value2 := 2

	if value1 == 1 && value2 == 2 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("No es verdad")
	}

  // Convert a string to number
  value, err := strconv.Atoi("53") // Atoi is like ParseInt, and strconv is a package to convert the data type from or to a string. It returns the number and an error if it cannot be done.
  if err != nil { // nil is the zero value for pointers, interfaces, maps, slices, channels and function types, representing an uninitialized value. An error is an interface, so a nil error means no error.
    log.Fatal(err) // Log creates logs, and fatal is like Print() but it also throws an Exit(1), it means that it interrupts totally the program.
  }
  fmt.Println("Value:", value)

}