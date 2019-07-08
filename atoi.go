package piscine

import (
	"math"
)

func Atoi(s string) int {
	num := 0
	valor := 1
	auxiliar := 0
	str2 := []rune(s)
	if len(str2) > 0 {
		if len(str2) > 1 {
			if str2[0] == 45 {
				valor = -1
				auxiliar = 1
			} else if str2[0] == 43 {
				auxiliar = 1
			}
			for i := auxiliar; i <= len(str2)-1; i++ {
				if str2[i] >= 48 && str2[i] <= 57 {
					num = num + (int(str2[i]-48))*int(math.Pow(10, float64(len(str2)-i-1)))
				} else {
					valor = 0
				}
			}
			return num * valor
		} else {
			if str2[0] >= 48 && str2[0] <= 57 {
				num = num + (int(str2[0] - 48))
			}
			return num
		}
	}
	return 0
}
