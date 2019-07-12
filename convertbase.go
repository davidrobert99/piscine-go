package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	valor := AtoiBase(nbr, baseFrom)
	final := PrintNbrBaseSimples(valor, baseTo)
	return final
}

func PrintNbrBaseSimples(nbr int, base string) string {
	aux := []rune(base)
	nbr1 := nbr
	numeroDivisoes := 0
	for nbr > 0 {
		numeroDivisoes++
		nbr = nbr / len(aux)
	}
	posicao := 0
	vetor := make([]rune, numeroDivisoes)
	for nbr1 > 0 {
		vetor[posicao] = rune(aux[nbr1%len(aux)])
		nbr1 = nbr1 / len(aux)
		posicao++
	}
	final := StrRev(string(vetor))
	return final
}
