package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Luhn algorithm
func LuhnCheck(cardNumber string) bool {
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")
	cardNumber = strings.ReplaceAll(cardNumber, "-", "")

	// keep track of the sum
	sum := 0
	alternate := false

	// loop from right to left
	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')

		// dkip non-numeric characters
		if !unicode.IsDigit(rune(cardNumber[i])) {
			continue
		}

		// double every second digit
		if alternate {
			digit *= 2
			// if doubling the digit results in a number greater than 9, subtract 9
			if digit > 9 {
				digit -= 9
			}
		}

		// add the digit to the sum
		sum += digit

		// alternate for the next iteration
		alternate = !alternate
	}

	// Check if the sum modulo 10 is 0
	return sum%10 == 0
}

func main() {
	// Example usage
	cardNumber := "4539 1488 0343 6467"
	if LuhnCheck(cardNumber) {
		fmt.Println("Card number is valid!")
	} else {
		fmt.Println("Card number is invalid!")
	}
}

