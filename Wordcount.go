package main

import (
	"fmt"
	"os"
	"strings"
)

var file_path string
var wordcount int

func main() {

	for {
		fmt.Print("Input the file path \n")
		fmt.Scanln(&file_path)
		_, error := read_file(file_path)
		// extractstring := string(Content)

		if error == nil {
			break
		} else {
			fmt.Println("An error occurred while trying to", error)

		}

	}

	Content, _ := read_file(file_path)
	extractstring := string(Content)
	words := strings.Fields(extractstring)
	wordcount = len(words)

	fmt.Print("The number of words in this file is:", wordcount)

}

func read_file(file_path string) ([]byte, error) {
	fileContent, err := os.ReadFile(file_path)
	if err != nil {
		return fileContent, err
		// fmt.Println("Error opening this file:", err)
	}
	return fileContent, err
}
