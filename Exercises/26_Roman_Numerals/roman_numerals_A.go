package romannumerals

import (
	"errors"
	"fmt"
)

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 {
		return "", errors.New("input must be a positive integer")
	}

	if input > 3999 {
		return "", errors.New(fmt.Sprintf("%d is out of range", input))
	}

	roman_number := ""

	for input != 0 {
		if input-1000 >= 0 {
			input -= 1000
			roman_number += "M"
		} else if input-900 >= 0 {
			input -= 900
			roman_number += "CM"
		} else if input-500 >= 0 {
			input -= 500
			roman_number += "D"
		} else if input-400 >= 0 {
			input -= 400
			roman_number += "CD"
		} else if input-100 >= 0 {
			input -= 100
			roman_number += "C"
		} else if input-90 >= 0 {
			input -= 90
			roman_number += "XC"
		} else if input-50 >= 0 {
			input -= 50
			roman_number += "L"
		} else if input-40 >= 0 {
			input -= 40
			roman_number += "XL"
		} else if input-10 >= 0 {
			input -= 10
			roman_number += "X"
		} else if input-9 >= 0 {
			input -= 9
			roman_number += "IX"
		} else if input-5 >= 0 {
			input -= 5
			roman_number += "V"
		} else if input-4 >= 0 {
			input -= 4
			roman_number += "IV"
		} else {
			input--
			roman_number += "I"
		}
	}
	return roman_number, nil
}
