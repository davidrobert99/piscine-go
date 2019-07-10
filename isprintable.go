package piscine

func IsPrintable(str string) bool {
	aux := []rune(str)
	for i := 0; i < len(aux); i++ {
		if aux[i] <= 31 {
			return false
		}
	}
	return true
}
