package main

import "fmt"

func main() {
  // DEFER -> 
  /* Execute the line at last time, just before the program ends. 
     You may use only one defer in a function.
     It's used to close database connections by instance.
  */
  
  defer fmt.Println("Hello") // It should be executed at first, but now it is executed at last.
  fmt.Println("World") // It is now executed at first.
  fmt.Println("Hey!") // It should be executed at last but now it is executed before the defer line.

  // CONTINUE AND BREAK
  for i := 0; i < 10; i++ {
    fmt.Println(i)

    // continue
    if i == 2 {
      fmt.Println("It's 2")
      continue 
      fmt.Println("This code shouldn't run, because there's a continue before this line, so the loop for should continue.")
    }

    // break
    if i == 8 {
      fmt.Println("Break")
      break // It interrupts a flow, in this case when the loop for finds a break, it is interrupted and exits
    }
  }
}