package main

import (
	"fmt"
	"strings"
)

func isPalyndrome(word string) bool {
	var wordReversed string

	for i := len(word) - 1; i >= 0; i-- {
		wordReversed += string(word[i])
	}

	return strings.ToLower(word) == strings.ToLower(wordReversed)
}

func main() {
	var word string
	fmt.Print("Please enter a word: ")
	fmt.Scanln(&word)

	if isPalyndrome(word) {
		fmt.Printf("The word %s is a palyndrome.\n", word)
	} else {
		fmt.Printf("The word %s isn't a palyndrome.\n", word)
	}
}
