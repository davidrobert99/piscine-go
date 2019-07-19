package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	auxiliar := l
	posicao := 0
	for auxiliar != nil {
		if posicao == pos {
			return auxiliar
		} else if posicao > pos {
			return nil
		} else {
			posicao++
			auxiliar = auxiliar.Next
		}
	}
	return nil
}
