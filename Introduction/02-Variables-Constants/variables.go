package main

import "fmt"

func main() {

	// VARIABLES -> Declaration and init

	var name1 string          // It's declared
	name1 = "Erika"           // It's initialized
	var name2 string = "Alex" // It can be declared and initialized at the same time
	name3 := "Juan"           // We avoid to specify the variable and data type. Go detects wich data type it is.

	fmt.Println(name1, name2, name3)

	// VARIABLES -> Declared without initialization, this way it have a default value.

	var a int     // Default -> 0
	var b float32 // Default -> 0
	var c string  // Default -> ""
	var d bool    // Default -> false

	fmt.Println(a, b, c, d)

	// CONSTANTS

	const pi = 3.1416

	fmt.Println(pi)
}
