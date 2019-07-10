package piscine

func NRune(s string, n int) rune {
	if s != "" {
		aux := []rune(s)
		if n <= len(aux) && n > 0 {
			return aux[n-1]
		}
	}
	return rune(95)
}
