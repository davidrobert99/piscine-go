package main

import (
	"fmt"
	"math"
)

func main() {
	s := "12345"
	s2 := "0000000012345"
	s3 := "012 345"
	s4 := "Hello World!"
	s5 := "+1234"
	s6 := "-1234"
	s7 := "++1234"
	s8 := "1"

	n := Atoi(s)
	n2 := Atoi(s2)
	n3 := Atoi(s3)
	n4 := Atoi(s4)
	n5 := Atoi(s5)
	n6 := Atoi(s6)
	n7 := Atoi(s7)
	n8 := Atoi(s8)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
	fmt.Println(n8)
}
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
