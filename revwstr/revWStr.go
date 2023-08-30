package main

import (
	"fmt"
	"os"
)

func reverseWords(input string) string {
	words := make([]string, 0)
	wordStart := -1

	for i, char := range input {
		if char != ' ' {
			if wordStart == -1 {
				wordStart = i
			}
		} else if wordStart != -1 {
			words = append(words, input[wordStart:i])
			wordStart = -1
		}
	}
	if wordStart != -1 {
		words = append(words, input[wordStart:])
	}

	reversed := ""
	for i := len(words) - 1; i >= 0; i-- {
		if i != len(words)-1 {
			reversed += " "
		}
		reversed += words[i]
	}

	return reversed
}

func main() {
	// Check if there is exactly one command-line argument
	if len(os.Args) != 2 {
		return
	}

	// Get the input string from the command-line argument
	input := os.Args[1]

	// Reverse the words and print the result
	reversed := reverseWords(input)
	fmt.Println(reversed)
}
