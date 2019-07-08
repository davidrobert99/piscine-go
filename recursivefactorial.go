package piscine

func IterativeFactorial(int nb) int {
	if nb == 0 {
		return 1
	} else {
		return nb * IterativeFactorial(nb-1)
	}
}
