package piscine

func Concat(str1 string, str2 string) string {
	aux1 := []rune(str1)
	aux2 := []rune(str2)
	final := make([]rune, (len(aux1) + len(aux2)))
	for i := 0; i < len(final); i++ {
		if i < len(aux1) {
			final[i] = aux1[i]
		} else {
			final[i] = aux2[i-len(aux1)]
		}
	}
	return string(final)
}
