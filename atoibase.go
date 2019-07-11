package piscine

import (
	"math"
	"strings"
)

func AtoiBase(s string, base string) int {
	number := 0
	aux := []rune(base)
	if len(base) > 1 {
		if verificaRepetidos(aux) {
			numero := []rune(s)
			if verificaRepetidos(aux) {
				for i := 0; i < len(numero); i++ {
					number = number + strings.Index(string(aux), string(numero[i]))*int(math.Pow(float64(len(base)), float64(len(numero)-i-1)))
				}
			}
		}
	}
	return number
}
