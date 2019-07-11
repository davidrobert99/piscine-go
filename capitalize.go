package piscine

/*func Capitalize(s string) string {
	aux := []rune(s)
	letraMaiuscula := true
	for i := 0; i < len(aux); i++ {
		if aux[i] >= 97 && aux[i] <= 122 && !letraMaiuscula {
			letraMaiuscula = true
		} else if aux[i] >= 97 && aux[i] <= 122 && letraMaiuscula {
			aux[i] = aux[i] - 32
			letraMaiuscula = false
		} else if aux[i] >= 65 && aux[i] <= 90 && !letraMaiuscula {
			aux[i] = aux[i] + 32
			letraMaiuscula = true
		} else if aux[i] >= 65 && aux[i] <= 90 && letraMaiuscula {
			letraMaiuscula = false
		} else {
			letraMaiuscula = true
		}
	}
	return string(aux)
}*/

func Capitalize(s string) string {
	aux := []rune(s)
	capitalize := true
	sai := true
	for i := 0; i < len(aux); i++ {
		if !(aux[i] >= 97 && aux[i] <= 122) && !(aux[i] >= 65 && aux[i] <= 90) && !(aux[i] >= 48 && aux[i] <= 57) {
			capitalize = true
		} else if !(aux[i] >= 97 && aux[i] <= 122) && capitalize == true {
			capitalize = false
			sai = true
			for j := i + 1; j < len(aux) && sai; j++ {
				if aux[j] >= 97 && aux[j] <= 122 {
					i++
				} else if aux[j] >= 65 && aux[j] <= 90 {
					aux[j] = aux[j] + 32
					i++
				} else {
					capitalize = true
					sai = false
				}
			}
		} else if aux[i] >= 97 && aux[i] <= 122 {
			if capitalize {
				aux[i] = aux[i] - 32
				capitalize = false
			}
		}
	}
	return string(aux)
}
