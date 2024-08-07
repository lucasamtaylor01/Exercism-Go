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

    romanNumerals := []struct {
        Value  int
        Symbol string
    }{
        {1000, "M"},
        {900, "CM"},
        {500, "D"},
        {400, "CD"},
        {100, "C"},
        {90, "XC"},
        {50, "L"},
        {40, "XL"},
        {10, "X"},
        {9, "IX"},
        {5, "V"},
        {4, "IV"},
        {1, "I"},
    }

    romanNumber := ""

    for _, numeral := range romanNumerals {
        for input >= numeral.Value {
            input -= numeral.Value
            romanNumber += numeral.Symbol
        }
    }

    return romanNumber, nil
}
