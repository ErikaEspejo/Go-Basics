package main

import "fmt"

func inputData() (user, password string) {
  var userInput, passwordInput string
  fmt.Println("-------------------")
  fmt.Print("Enter your username: ")
  fmt.Scanln(&userInput)
  fmt.Print("Enter your password: ")
  fmt.Scanln(&passwordInput)
  return userInput, passwordInput
}

func main() {
  var user, userLogin, password, passwordLogin string
  
  fmt.Println("Please register your user!")
  user, password = inputData()
  fmt.Println("\n----User succesfully created!----")

  fmt.Println("Please login!")
  userLogin, passwordLogin = inputData()
  if userLogin == user && passwordLogin == password {
    fmt.Println("\n----User succesfully logged!----")
  } else {
    fmt.Println("\n----User or Password incorrect!----")
  }
  
}