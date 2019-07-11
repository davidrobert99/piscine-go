package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if len(base) > 1 {
		aux := []rune(base)
		if verificaRepetidos(aux) {
			if nbr < 0 {
				z01.PrintRune('-')
				if nbr == -9223372036854775808 {
					fmt.Print("9223372036854775808")
				}
				nbr = -nbr
			}
			nbr1 := nbr
			numeroDivisoes := 0
			for nbr > 0 {
				numeroDivisoes++
				nbr = nbr / len(aux)
			}
			posicao := 0
			vetor := make([]rune, numeroDivisoes)
			for nbr1 > 0 {
				vetor[posicao] = rune(nbr1 % len(aux))
				nbr1 = nbr1 / len(aux)
				posicao++
			}
			for i := len(vetor) - 1; i >= 0; i-- {
				fmt.Print(string(aux[vetor[i]]))
			}
		} else {
			z01.PrintRune('N')
			z01.PrintRune('V')
		}
	} else {
		z01.PrintRune('N')
		z01.PrintRune('V')

	}
}

func verificaRepetidos(vetor []rune) bool {
	for i := 0; i < len(vetor); i++ {
		for j := i + 1; j < len(vetor); j++ {
			if vetor[i] == vetor[j] || vetor[i] == 45 || vetor[j] == 43 {
				return false
			}
		}
	}
	return true
}
