package main

import "fmt"

/*
3. Statement: Given the sale value of products, find the IGV (18%) and the sale price. To solve this problem, it is required that
the user enters the sale value of the product and the system performs the respective calculation to find the IGV and the sale price.
*/

func IGV(price float32) float32 {
	return price * 0.18
}

func totalPrice(price float32) float32 {
	return price + IGV(price)
}

func printSaleValues(price float32) {
	fmt.Print("Please enter a sale value:")
	fmt.Scanln(&price)
	fmt.Println("IGV:", IGV(price))
	fmt.Println("Sale Price:", totalPrice(price))
}

func main() {
	var price float32
	printSaleValues(price)

}
