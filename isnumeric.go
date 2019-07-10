package piscine

func IsNumeric(str string) bool {
	aux := []rune(str)
	for i := 0; i < len(aux); i++ {
		if !(aux[i] >= 48 && aux[i] <= 57) {
			return false
		}
	}
	return true
}
