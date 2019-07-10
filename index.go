package piscine

func Index(s string, toFind string) int {
	aux := []rune(s)
	aux2 := []rune(toFind)
	if len(aux) > 0 && len(aux2) > 0 {
		if len(aux) < len(aux2) {

			return -1
		}
		for i := 0; i < len(aux); i++ {
			if aux[i] == aux2[0] {
				valor := i
				for j := 0; j < len(aux2); j++ {
					if aux[i] != aux2[j] {
						return -1
					}
					if i <= len(aux)-1 && aux[i] == aux2[j] {
						i++
					}

				}
				return valor
			}
		}
	}

	return -1
}
