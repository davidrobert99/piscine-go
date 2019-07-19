package piscine

func ListMerge(l1 *List, l2 *List) {
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
