package main

import "fmt"

var BMImethod string

var BMIvalue, weight, height, FinalBMIValue float64

func main() {
	validinput := []string{"A", "B"}
	//Error handling to ensure does not input an Integer when program expects string
	for {
		fmt.Println("Input A to calculate BMI in Standard unit i.e weight in pounds and height in feet and inches")
		fmt.Println("Input B to calculate BMI in Metric unit i.e weight in kg and height in centimeters")
		fmt.Scanln(&BMImethod)

		if CheckInputValidity(BMImethod, validinput) {
			break
		} else {
			fmt.Println("You entered an invalid input. Try again with an option between A and B")
		}
	}
	switch BMImethod {

	case "A":
		//Error handling to ensure does not input a string when program expects integer or float
		for {
			fmt.Println("What do you weight in pound")
			// fmt.Scanln(&weight)
			_, err := fmt.Scan(&weight)

			if err == nil {
				break
			} else {
				fmt.Println("Input is not an integer or float")
			}
		}

		for {
			fmt.Println("What is your height?")
			fmt.Println("Please input in feet and inches e.g 5.11, 5.7")
			// fmt.Scanln(&height)
			_, err := fmt.Scan(&height)

			if err == nil {
				break
			} else {
				fmt.Println("Input is not an integer or float")
			}
		}
		FinalBMIValue = BMICalculator(weight, height)

		fmt.Printf("Your BMI is %.2f\n", FinalBMIValue)

		Comment := checkBMIcategory(FinalBMIValue)
		fmt.Println(Comment)

	case "B":

		for {
			fmt.Println("What do you weight in kilogram")
			// fmt.Scanln(&weight)
			_, err := fmt.Scan(&weight)

			if err == nil {
				break
			} else {
				fmt.Println("Input is not an integer or float")
			}
		}

		for {
			fmt.Println("What is your height?")
			fmt.Println("Please input in centimeter e.g 125, 150")
			// fmt.Scanln(&height)
			_, err := fmt.Scan(&height)

			if err == nil {
				break
			} else {
				fmt.Println("Input is not an integer")
			}

		}
		FinalBMIValue = BMICalculator(weight, height)

		fmt.Printf("Your BMI is %.2f\n", FinalBMIValue)

		Comments := checkBMIcategory(FinalBMIValue)
		fmt.Println(Comments)

	}

}

// function to calculate the BMI using standard or metric unit giving user flexibility in measurement choices.
func BMICalculator(weight float64, height float64) float64 {

	switch BMImethod {
	case "A":
		feetinInches := int(height) * 12
		Totalinches := float64(feetinInches) + float64(height)
		heightsquared := Totalinches * Totalinches
		Value := weight / heightsquared
		BMIvalue = Value * 703
	case "B":
		CentemeterToMeteres := height * 0.01
		heightsquared := CentemeterToMeteres * CentemeterToMeteres
		BMIvalue = weight / heightsquared
	}
	return BMIvalue
}

//This function is a conditional statement to check if user input one of the options in our slices of strings
func CheckInputValidity(input string, validinput []string) bool {
	for _, inputed := range validinput {
		if inputed == input {
			return true
		}
	}
	return false
}

//This function helps  to check which category the user BMI falls
func checkBMIcategory(BMI float64) string {
	var category string
	if BMI >= 18.0 {
		category = ("Your BMI indicate that you are underweight")
	}
	if BMI > 18.5 && FinalBMIValue <= 24.9 {
		category = ("Your BMI indicate that you are of normal weight")
	}
	if BMI >= 25 {
		category = ("Your BMI indicate that you overweight")
	}
	return category
}
