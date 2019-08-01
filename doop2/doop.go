package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 4 {
		if IsNumeric(arguments[1]) && IsNumeric(arguments[3]) {

		} else {
			fmt.Println(0)
		}
	}
}

func IsNumeric(str string) bool {
	aux := []rune(str)
	for i := 0; i < len(aux); i++ {
		if !(aux[i] >= 48 && aux[i] <= 57) {
			return false
		}
	}
	return true
}
