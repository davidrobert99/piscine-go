package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for i := len(arguments) - 1; i > 0; i-- {
		s := []rune(arguments[i])
		for j := 0; j < len(s); j++ {
			z01.PrintRune(s[j])
		}
		z01.PrintRune('\n')
	}
}
