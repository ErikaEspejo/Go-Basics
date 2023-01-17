package main

import "fmt"

// Struct: Allows you to create something like classes of POO
// Struct declaration
type Car struct {
	brand string // atributes
	year  int
}

func main() {
	myCar := Car{brand: "Ford", year: 2020} // Instance of the struct (Object)
	fmt.Println(myCar)
	fmt.Println(myCar.brand)
	fmt.Println(myCar.year)

	var yourCar Car
	yourCar.brand = "Ferrari"
	fmt.Println(yourCar)
	fmt.Println(yourCar.brand)
	fmt.Println(yourCar.year) // When it wasn't stored an attribute, it will take a zero value

}
