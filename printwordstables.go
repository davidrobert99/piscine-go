package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(table []string) {
	for i := 0; i < len(table); i++ {
		vetor := []rune(table[i])
		for j := 0; j < len(vetor); j++ {
			z01.PrintRune(vetor[j])
			if j == len(vetor)-1 {
				z01.PrintRune('\n')
			}
		}
	}
}
