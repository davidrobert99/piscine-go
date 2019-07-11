package piscine

func Capitalize(s string) string {
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
}
