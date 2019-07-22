package piscine

func ListReverse(l *List) {
	anterior := &NodeL{}
	anterior = nil
	seguinte := &NodeL{}
	seguinte = nil
	atual := l.Head
	l.Head = l.Tail
	l.Tail = atual
	for atual != nil {
		seguinte = atual.Next
		atual.Next = anterior
		anterior = atual
		atual = seguinte
	}
}
