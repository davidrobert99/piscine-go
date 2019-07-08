package piscine

import (
	"math"
)

func BasicAtoi(s string) int {
	num := 0
	str2 := []rune(s)
	for i := 0; i <= len(str2)-1; i++ {
		num = num + (int(str2[i]-48))*int(math.Pow(10, float64(len(str2)-i-1)))
	}
	return num
}
