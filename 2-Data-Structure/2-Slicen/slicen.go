package main

import "fmt"

func main() {
	// Slicen -> Dynamic Array
	numbers := []int{1, 2, 3} //[empty] it means it is a slicen
	fmt.Println(numbers, len(numbers))

	// Add elements to slicen
	numbers = append(numbers, 4, 5, 6)
	fmt.Println(numbers, len(numbers))

	// Sub slicen
	subSlicen := numbers[:2]
	fmt.Println(subSlicen)

	/* Subslicen are considered as children of a father slicen, so when there is any modification to the father slicen,
	the children also are going to be modified */

	numbers[0] = 100 // it's been modified after the subslicen is created

	fmt.Println(subSlicen) // Prints [100 2]
	fmt.Println(numbers)   // Prints [100 2 3 4 5 6]

	/* Both were affected by the modification because subSlicen is a child of numbers. Even when the modification
	has been made after it was created. */
}
