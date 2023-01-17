package main

// A variable with private access needs to begin in minusc.
// A variable with public access needs to begin in mayusc.

import (
	"fmt"

	pk "./mypackage"
)

func main() {
	var myPublicCar pk.CarPublic
	myPublicCar.Brand = "Toyota"
	myPublicCar.Year = 2020
	fmt.Println(myPublicCar)
	fmt.Println(myPublicCar.Brand)
	fmt.Println(myPublicCar.Year)

	/* You cannot call a private struct
	var myPrivateCar pk.carPrivate
	fmt.Println(myPrivateCar)
	*/
	pk.PrintMessage(myPublicCar.Brand)
}
