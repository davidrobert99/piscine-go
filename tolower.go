package piscine

func ToLower(s string) string {
	aux := []rune(s)
	for i := 0; i < len(aux); i++ {
		if aux[i] >= 65 && aux[i] <= 90 {
			aux[i] = aux[i] + 32
		}
	}
	return string(aux)
}
