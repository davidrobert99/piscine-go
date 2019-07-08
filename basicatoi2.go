package piscine

import (
	"math"
)

func BasicAtoi2(s string) int {
	num := 0
	valor := 1
	str2 := []rune(s)
	for i := 0; i <= len(str2)-1; i++ {
		if str2[i] >= 48 && str2[i] <= 57 {
			num = num + (int(str2[i]-48))*int(math.Pow(10, float64(len(str2)-i-1)))
		} else {
			valor = 0
		}
	}
	return num * valor
}
