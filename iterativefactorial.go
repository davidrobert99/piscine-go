package piscine

func IterativeFactorial(nb int) int {
	if nb <= 0 || nb >= 21 {
		return 0
	}
	valor := 1
	for nb > 0 {
		valor = valor * nb
		nb--
	}
	return valor
}
