package armstrong

import (
	"math"
	"strconv"
)

func NumberMagnitude(n int) int {
	magnitude := 0
	for n >= int(math.Pow(10, float64(magnitude))) {
		magnitude++
	}
	return magnitude
}

func SepareDigits(n int) ([]int, error) {
	numStr := strconv.Itoa(n)
	digitos := make([]int, len(numStr))

	for i, char := range numStr {
		digito, err := strconv.Atoi(string(char))
		if err != nil {
			return nil, err
		}
		digitos[i] = digito
	}

	return digitos, nil
}

func IsNumber(n int) bool {
	if n == 0 {
		return true
	}
	digits, _ := SepareDigits(n)

	magnitude := NumberMagnitude(n)
	sum := 0

	for _, digit := range digits {
		pot := math.Pow(float64(digit), float64(magnitude))
		sum += int(pot)
	}

	return sum == n
}
