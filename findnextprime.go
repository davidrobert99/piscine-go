package piscine

import "github.com/01-edu/z01"

func FindNextPrime(nb int) int {
	if nb > 0 {
		encontrou := false
		i := nb
		valor := nb
		for !encontrou {
			valor = i
			encontrou = IsPrime(i)
			i++
			if i >= z01.MaxInt {
				return 0
			}
		}
		return valor
	} else {
		return 2
	}
}
