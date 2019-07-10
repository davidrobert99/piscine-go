package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if len(base) > 1 {
		aux := []rune(base)
		inverteu := false
		nbr1 := nbr
		if nbr < 0 {
			nbr = -nbr
			inverteu = true
		}
		numeroDivisoes := 0
		for nbr > 0 {
			numeroDivisoes++
			nbr = nbr / len(aux)
		}
		posicao := 0
		vetor := make([]rune, numeroDivisoes)
		for nbr1 > 0 {
			vetor[posicao] = rune(nbr1 % len(aux))
			nbr1 = nbr1 % len(aux)
			posicao++
		}
		if inverteu {
			z01.PrintRune('-')

		}
		//faz print
		fmt.Println(vetor)
	} else {
		z01.PrintRune('N')
		z01.PrintRune('B')
		z01.PrintRune('\n')
	}
}
