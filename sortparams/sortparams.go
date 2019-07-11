package main

import (
	"os"
	"sort"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	sort.Strings(arguments)
	for i := 1; i < len(arguments); i++ {
		if len(arguments[i]) == 1 {
			z01.PrintRune(rune(arguments[i][0]))
			z01.PrintRune('\n')
		}
	}
}
