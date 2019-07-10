package piscine

func IsAlpha(str string) bool {
	aux := []rune(str)
	for i := 0; i < len(aux); i++ {
		if !((aux[i] >= 97 && aux[i] <= 122) || (aux[i] >= 65 && aux[i] <= 90) || (aux[i] >= 48 && aux[i] <= 57)) {
			return false
		}
	}
	return true
}
