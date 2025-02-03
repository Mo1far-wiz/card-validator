package utils

import (
	"strings"
	"unicode"
)

func CheckLuhn(s string) bool {
	// use only numbers in case of something like 1111-1111-1111-1111
	numbersStr := strings.ReplaceAll(s, "-", "")
	if numbersStr == "" {
		return false
	}

	digits := len(numbersStr)

	numbers := make([]int, digits)
	for idx, val := range numbersStr {
		if !unicode.IsDigit(val) {
			return false
		}
		numbers[idx] = int(val - '0')
	}

	var sum int = 0
	for i, second := digits-1, false; i >= 0; i, second = i-1, !second {
		var d int = numbers[i]
		if second {
			d = d * 2
		}

		sum += d / 10
		sum += d % 10
	}

	return sum%10 == 0
}
