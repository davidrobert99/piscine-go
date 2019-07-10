package piscine

func IsUpper(str string) bool {
	aux := []rune(str)
	for i := 0; i < len(aux); i++ {
		if !(aux[i] >= 65 && aux[i] <= 90) {
			return false
		}
	}
	return true
}
