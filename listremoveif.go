package piscine

import "fmt"

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head != nil {
		percorre := l.Head
		anterior := l.Head
		for percorre != nil {
			if percorre.Data == data_ref {
				anterior.Next = percorre.Next
				if percorre.Next == nil {
					l.Tail = anterior
					if anterior != nil && anterior.Data == data_ref {
						l.Tail = nil
						if l.Head.Data == data_ref {
							l.Head = l.Head.Next
						}
						return
					}
				}
				percorre = percorre.Next
			} else {
				anterior = percorre
				percorre = percorre.Next
			}
		}
	}
	if l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
}

func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

/*
	for percorre != nil {
		if percorre.Data == data_ref {
			anterior.Next = percorre.Next
			if percorre.Next == nil {
				l.Tail = anterior
			}
			percorre = percorre.Next
		} else {
			anterior = percorre
			percorre = percorre.Next
		}
	}
} */
