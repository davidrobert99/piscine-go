package piscine

/*func ListReverse(l *List) {
	auxiliar := l
	if ListSize(l) > 1 {
		var vetor []interface{}
		for auxiliar.Head != nil {
			vetor = append(vetor, auxiliar.Head.Data)
			auxiliar.Head = auxiliar.Head.Next
			fmt.Println(auxiliar)
		}
		link := &List{}
		link.Head = nil
		link.Tail = nil
		for i := len(vetor) - 1; i >= 0; i-- {
			ListPushBack(link, vetor[i])
		}
		*l = *link

	} else {
		fmt.Print("ahha")
	}
}*/

func ListReverse(l *List) {
	if ListSize(l) > 1 {
		if ListSize(l) == 2 {

		} else {

		}
	}
}

func reverse ( l *List ){
	anterior := nil
	seguinte := nil
	atual := l.Head
	auxiliar := l.Head
	l.head = l.tail 
	l.tail = auxiliar
	for atual!= nil{
		seguinte = atual.next
		atual.next = anterior
		anterior = atual
		atual = next
	}

}  

