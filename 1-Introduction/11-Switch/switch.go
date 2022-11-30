package main

import "fmt"

func main() {

  // Switch when you need a specific value
	module := 4 % 2
  switch module {
  // It also can be written as:
  // switch module := 5 % 2; module {
    case 0:
      fmt.Println("Es par")
    default:
      fmt.Println("Es impar")
  }

  // Switch as conditional if
  value := 50
  switch {
    case value > 100: // condition
      fmt.Println("Es mayor a 100")
    case value < 100: // condition
      fmt.Println("Es menor a 100")
    default: // default condition
      fmt.Println("Es 100")
  }
}
