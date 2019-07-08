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

	n := BasicAtoi2(s)
	n2 := BasicAtoi2(s2)
	n3 := BasicAtoi2(s3)
	n4 := BasicAtoi2(s4)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
}

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
