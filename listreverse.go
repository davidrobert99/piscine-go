package piscine

import "fmt"

func ListReverse(l *List) {
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
}
