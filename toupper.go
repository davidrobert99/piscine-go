package piscine

func ToUpper(s string) string {
	aux := []rune(s)
	for i := 0; i < len(aux); i++ {
		if aux[i] >= 97 && aux[i] <= 122 {
			aux[i] = aux[i] - 32
		}
	}
	return string(aux)
}
