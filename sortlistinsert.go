package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	l = ListPushBackNode(l, data_ref)
	l = ListSort(l)
	return l
}

func ListPushBackNode(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data, Next: nil}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

// a lista esta ordenada, basta verificar os elementos
