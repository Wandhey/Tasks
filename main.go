package main

import "fmt"

var BaseCurrency, ConvertToCurrency string

var Amount, DollarEquivalent, FinalValue float64

func main() {
	CurrencyList := []string{"Naira", "Euro", "Dollar", "Pound", "Yuan", "Riyal", "Ruble", "Cedis", "Rupee", "Rand"}

	for {
		fmt.Println("Input a base currency from the list:\n Naira\n Euro\n Dollar\n Pound\n Yuan\n Riyal\n Ruble\n Cedis\n Rupee\n Rand")
		fmt.Scanln(&BaseCurrency)

		if CheckCurrencyValidity(BaseCurrency, CurrencyList) {
			break
		} else {
			fmt.Println("Invalid Currency, Please input a valid Currency")
		}
	}

	for {
		fmt.Println("Convert to what currency:\n Naira\n Euro\n Dollar\n Pound\n Yuan\n Riyal\n Ruble\n Cedis\n Rupee\n Rand")
		fmt.Scanln(&ConvertToCurrency)

		if CheckCurrencyValidity(ConvertToCurrency, CurrencyList) {
			break
		} else {
			fmt.Println("Invalid Currency, Please input a valid Currency")
		}
	}

	fmt.Printf("Input %v amount:", BaseCurrency)

	fmt.Scanln(&Amount)

	// Variable DollarEquivalent is the base currency equivalent in dollar
	DollarEquivalent = ConvertBasetoDollar(BaseCurrency)

	// ConvertBasetoDollar is function that helps us first conver the base currency to dollar
	FinalValue = DollarToDesiredCurrency(ConvertToCurrency)

	//DollarToDesiredCurrency helps convert the dollar equivalent of the base currency to the desired currency
	fmt.Printf("The value in %s is %.2f", ConvertToCurrency, FinalValue)

}

func ConvertBasetoDollar(BaseCurrency string) float64 {
	var DollarValue float64
	switch BaseCurrency {
	case "Naira":
		DollarValue = Amount / 1635
	case "Euro":
		DollarValue = Amount / 0.92
	case "Dollar":
		DollarValue = Amount
	case "Pound":
		DollarValue = Amount / 0.78
	case "Yuan":
		DollarValue = Amount / 7.16
	case "Riyal":
		DollarValue = Amount / 3.75
	case "Ruble":
		DollarValue = Amount / 84.97
	case "Cedis":
		DollarValue = Amount / 15.41
	case "Rand":
		DollarValue = Amount / 18.28
	}
	return DollarValue

}

func DollarToDesiredCurrency(ConvertToCurrency string) float64 {
	var ToCurrencyValue float64
	switch ConvertToCurrency {
	case "Naira":
		ToCurrencyValue = DollarEquivalent * 1635
	case "Euro":
		ToCurrencyValue = DollarEquivalent * 0.92
	case "Dollar":
		ToCurrencyValue = DollarEquivalent
	case "Pound":
		ToCurrencyValue = DollarEquivalent * 0.78
	case "Yuan":
		ToCurrencyValue = DollarEquivalent * 7.16
	case "Riyal":
		ToCurrencyValue = DollarEquivalent * 3.75
	case "Ruble":
		ToCurrencyValue = DollarEquivalent * 84.97
	case "Cedis":
		ToCurrencyValue = DollarEquivalent * 15.41
	case "Rand":
		ToCurrencyValue = DollarEquivalent * 18.28
	}
	return ToCurrencyValue
}

//This function helps check if the user inputed a valid currency
func CheckCurrencyValidity(BaseCurrency string, CurrencyList []string) bool {
	for _, currency := range CurrencyList {
		if currency == BaseCurrency {
			return true
		}
	}
	return false
}
