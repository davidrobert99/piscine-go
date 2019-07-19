package piscine

func Rot14(str string) string {
	vetor := []rune(str)
	resposta := make([]rune, len(str))
	for i := 0; i < len(str); i++ {
		valor := vetor[i]
		if valor >= 65 && valor <= 90 {
			valor = valor + 14
			if valor > 90 {
				valor = valor - (90 - 65 + 1)
			}
			resposta[i] = valor
		} else if valor >= 97 && valor <= 122 {
			valor = valor + 14
			if valor > 122 {
				valor = valor - (122 - 97 + 1)
			}
			resposta[i] = valor
		} else {
			resposta[i] = vetor[i]
		}
	}
	return string(resposta)
}
