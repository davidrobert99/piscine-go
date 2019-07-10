package piscine

func LastRune(s string) rune {
	if s != "" {
		aux := []rune(s)
		return aux[len(aux)-1]
	}
	return rune(95)
}
