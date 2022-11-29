package main

import "fmt"

func main() {
	// Conditional FOR
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// For While
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For in range
	evenNumberList := []int{2, 4, 6, 8}

	for i, number := range evenNumberList {
		fmt.Printf("Position: %d, Even number: %d \n", i, number)
	}

	// For FOREVER
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}

}
