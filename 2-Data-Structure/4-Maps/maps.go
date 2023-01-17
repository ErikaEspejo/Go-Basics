package main

import "fmt"

func main() {

	// map is used to create maps [key]value
	ages := make(map[string]int) // make is used to create any variable datatype.

	ages["Jose"] = 14
	ages["Pepito"] = 20

	fmt.Println(ages)

	// It also can be declared like this:
	temperature := map[string]int{
		"China":    15,
		"Colombia": 14,
		"Canada":   3,
	}

	fmt.Println(temperature)

	// Print each item of the map
	for key, value := range ages {
		fmt.Println(key, value)
	}

	// Get value from a given key
	// There is a second to get a boolean which indicates if that key actually exists.
	// Normally it's declared as "ok", but it could get any name.
	// Maps are more efficient than slice or arrays, because it implements native concurrence.
	joseAge, yes := ages["Jose"]
	josepAge, ok := ages["Josep"]

	fmt.Println(joseAge, yes)
	fmt.Println(josepAge, ok) // In case the key doesnt exist, it will get a zero value

}
