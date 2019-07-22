package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
	} else {
		percorre := l1.Head
		for percorre.Next != nil {
			percorre = percorre.Next
		}
		percorre.Next = l2.Head
		for percorre.Next != nil {
			percorre = percorre.Next
		}
		l1.Tail = percorre
	}
}
