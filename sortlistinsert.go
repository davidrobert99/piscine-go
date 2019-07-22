package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	ListPushBackNode(l, data_ref)
	l = ListSort(l)
	return l
}

func ListPushBackNode(node *NodeI, data int) {
	n := &NodeI{Data: data, Next: nil}
	if node == nil {
		node = nil
	} else {
		copia := node
		for copia != nil {
			copia = copia.Next
		}
		copia = n
	}
}
