package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
	percorre := l.Head
	anterior := l.Head
	for percorre != nil {
		if percorre.Data == data_ref {
			anterior.Next = percorre.Next
			percorre = percorre.Next
		} else {
			anterior = percorre
			percorre = percorre.Next
		}
	}

}
