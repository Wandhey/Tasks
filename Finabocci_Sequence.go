package main

import "fmt"

var Userinput int
var Finabocci_Sequence, linear_series []int

func main() {
	//Error handling to check for wrong input
	var final_fib []int
	for {
		fmt.Println("Please input a number")
		_, err := fmt.Scan(&Userinput)

		if err == nil {
			break
		} else {
			fmt.Println("Invalid input")
		}
	}
	//We iterate over the slice of integers created by our linear_sequence function for the fib value
	linear_series = linearsequence(Userinput)
	for _, number := range linear_series {
		final_fib = append(final_fib, adjusted_fibonacci(number))
	}

	fmt.Println("The fibonacci sequence:", final_fib)

}

//This function helps to create a slice of integers up to nth number(number inputted by user)
func linearsequence(n int) []int {
	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		sequence[i] = i * 1
	}
	return sequence
}

//My fib function
func Finabocci(n int) int {
	var result int

	if n == 0 {
		result = 0
	}
	if n == 1 {
		result = 1
	}
	if n > 1 {
		result = (n - 1) + (n - 2)
	}
	return result

}

func adjusted_fibonacci(n int) int {
	var result int

	if n == 0 {
		result = 0
	}
	if n == 1 {
		result = 1
	}
	if n > 1 {
		result = Finabocci(n-1) + (n - 2)
	}
	return result
}
