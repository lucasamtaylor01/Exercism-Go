
package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	// Tratamento 00: retirando espaços em branco
	id = strings.ReplaceAll(id, " ", "")

	// Tratamento 01: Verificando o tamanho da string
	if len(id) <= 1 {
		return false
	}

	// Tratamento 02: Verificando se não há letras
	for i := 0; i < len(id); i++ {
		if id[i] < '0' || id[i] > '9' {
			return false
		}
	}

	// Tratamento 03: Convertendo para int e colocando no slice
	var code []int
	for i := 0; i < len(id); i++ {
		num, _ := strconv.Atoi(string(id[i]))
		code = append(code, num)
	}

	// Fase 0
	for i := 0; i < len(code); i++ {
		if i%2 == 0 {
			code[i] *= 2
		}
	}

	// Fase 1
	for i := 0; i < len(code); i++ {
		if code[i] > 9 {
			code[i] -= 9
		}
	}

	sum := 0
	for i := 0; i < len(code); i++ {
		sum += code[i]
	}

	if sum%10 == 0 {
		return true
	} else {
		return false
	}
}