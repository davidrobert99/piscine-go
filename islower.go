package piscine

func IsLower(str string) bool {
	aux := []rune(str)
	for i := 0; i < len(aux); i++ {
		if !(aux[i] >= 97 && aux[i] <= 122) {
			return false
		}
	}
	return true
}
