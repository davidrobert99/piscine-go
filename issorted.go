package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	return menorParaMaior(f, tab) || maiorParaMenor(f, tab)
}

func menorParaMaior(f func(a, b int) int, tab []int) bool {
	for i := 1; i < len(tab); i++ {
		if f(tab[i-1], tab[i]) >= 0 {
			return false
		}
	}
	return true
}

func maiorParaMenor(f func(a, b int) int, tab []int) bool {
	for i := len(tab) - 1; i > 0; i-- {
		if f(tab[i], tab[i-1]) > 0 {
			return false
		}
	}
	return true
}
