package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var input string

func main() {
	//had to use bufio as fmt.scanln was only picking the first word of phrases
	fmt.Println("Please input a word or a phrase")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	if PalindromChecker(input) {
		fmt.Println(input, ":is a Palindrome")
	} else {
		fmt.Println(input, ":is not a Palindrome")
	}

}

func PalindromChecker(input string) bool {
	var correctedinput []rune
	for _, character := range input {
		if unicode.IsLetter(character) || unicode.IsDigit(character) {
			correctedinput = append(correctedinput, unicode.ToLower(character))
		}
	}

	cleanedcharacter := string(correctedinput)
	n := len(cleanedcharacter)

	for i := 0; i < n/2; i++ {
		if cleanedcharacter[i] != cleanedcharacter[n-1-i] {
			return false

		}

	}

	return true
}
