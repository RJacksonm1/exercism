package luhn

import (
	"strconv"
	"strings"
)

// Valid checks whether a card number is valid as per the Luhn algorithm
func Valid(numberString string) bool {
	numberString = strings.Replace(numberString, " ", "", -1)
	numberStringSlice := strings.Split(numberString, "")
	numberStringLength := len(numberStringSlice)

	// Not a Luhn number if it doesnt have a check digit
	if numberStringLength < 2 {
		return false
	}

	// We double every second digit from the right
	doublingOffset := numberStringLength % 2

	var sum int
	for i, n := range numberStringSlice {
		nInt, err := strconv.Atoi(n)

		// Non-numerics not allowed
		if err != nil {
			return false
		}

		if i%2 == doublingOffset {
			nInt *= 2
		}

		if nInt > 9 {
			nInt -= 9
		}

		sum += nInt
	}

	return sum%10 == 0
}
