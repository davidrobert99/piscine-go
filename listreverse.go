package piscine

func ListReverse(l *List) {
	auxiliar := l.Head
	if auxiliar.Next != nil {
		var vetor []interface{}
		for auxiliar != nil {
			vetor = append(vetor, auxiliar.Data)
			auxiliar = auxiliar.Next
		}
		link := &List{}
		link.Head = nil
		link.Tail = nil
		for i := len(vetor) - 1; i >= 0; i-- {
			ListPushBack(link, vetor[i])
		}
		*l = *link
	}
}
