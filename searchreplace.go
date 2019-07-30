package piscine

import (
	"fmt"
	"os"
)

func Searchreplace() {
	arguments := os.Args
	if len(arguments) > 1 && len(arguments) <= 3 {
		resultado := arguments[1]
		for i := range arguments[1] {
			if arguments[1][i] == arguments[1][0] {
				resultado[i] = arguments[2][0]
			}
		}
	} else {
		fmt.Println("")
	}
}
