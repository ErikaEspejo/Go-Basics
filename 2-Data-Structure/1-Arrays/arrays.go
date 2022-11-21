package main

import "fmt"

func main() {
	// Arrays -> Store data of an specific data type. When the array has no information, it will store default values.
	var numbers [5]int
	fmt.Println(numbers)
	// Prints [0 0 0 0 0]

	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	fmt.Println(numbers)
	// Prints [1 2 3 4 5]
	fmt.Println(numbers[2])
	// Prints 3
	// Arrays are inmutable, it means, the size couldnt be changed. But the elements could be changed by its index.

	//Arrays with values
	names := [3]string{
		"Erika",
		"Maria",
		"José",
	}

	fmt.Println(names)
	// Prints [Erika Maria Jose]

	// Array without length
	colors := [...]string{
		"Red",
		"Green",
		"Black",
		"Blue",
	}
	fmt.Println(colors, len(colors)) // len function get the array lenght

	// Arrays with defined index
	currency := [...]string{
		0: "Dolar",
		2: "Euro",
		3: "Pesos",
		5: "Soles",
	}
	fmt.Println(currency, len(currency))

	// Get a piece of the array
	arrayPiece := currency[1:3] // originalArray[startIncluded : endExcluded]
	fmt.Println(arrayPiece)     //Prints ["" "Euro"]
	arrayPiece2 := colors[0:3]
	fmt.Println(arrayPiece2)        // Prints [Red Green Black]
	arrayPieceToEnd := currency[3:] // Without using an end index, it will split to the end of array
	fmt.Println(arrayPieceToEnd)
}
