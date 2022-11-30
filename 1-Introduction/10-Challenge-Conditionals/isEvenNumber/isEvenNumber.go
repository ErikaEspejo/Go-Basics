package main

import "fmt"

func isEven(number int) bool {
  return number % 2 == 0
}

func main() {
  var number int
  
  fmt.Print("Please enter an integer number: ")
  fmt.Scanln(&number)

  if isEven(number) {
    fmt.Printf("%d is an even number.\n", number)
  } else {
    fmt.Printf("%d is an odd number.\n", number)
  }
  
}