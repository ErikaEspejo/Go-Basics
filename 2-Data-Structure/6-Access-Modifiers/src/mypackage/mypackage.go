package mypackage

import "fmt"

// Public struct and public attributes
type CarPublic struct {
	Brand string
	Year  int
}

// private struct and private attributes
type carPrivate struct {
	brand string
	year  int
}

// Function with public access
func PrintMessage(text string) {
	fmt.Printf("Hola %s", text)
}
