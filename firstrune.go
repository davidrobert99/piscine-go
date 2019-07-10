package piscine

func FirstRune(s string) rune {
	if s != "" {
		aux := []rune(s)
		return aux[0]
	}
	return rune(95)
}
