package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	vet := make([]int, n)
	for i := 0; i < n; i++ {
		vet[i] = i + 48
	}
	pos := n - 1
	auxiliarincrementa(n, vet, pos)
	auxiliarimprime(vet)
	z01.PrintRune(rune('\n'))
}

func auxiliarincrementa(n int, vet []int, pos int) {
	if vet[0] == (58 - n) {
		return
	} else {
		auxiliarimprime(vet)
		z01.PrintRune(rune(','))
		z01.PrintRune(rune(' '))
		if vet[pos] != (9 + 48) {
			vet[pos]++
			auxiliarincrementa(n, vet, pos)
		} else {
			pos--
			if pos >= 0 {
				for pos >= 0 && vet[pos] == (57-(n-pos-1)) {
					pos--
				}
				vet[pos]++
			}
			for pos < n-1 && pos >= 0 {
				vet[pos+1] = vet[pos] + 1
				pos++
			}
			auxiliarincrementa(n, vet, pos)
		}
	}
}

func auxiliarimprime(vet []int) {
	for i := 0; i < len(vet); i++ {
		z01.PrintRune(rune(vet[i]))
	}
}
