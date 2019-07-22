package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n2 != nil && n1 != nil {
		percorre := n1
		for percorre.Next != nil {
			percorre = percorre.Next
		}
		percorre.Next = n2
		n1 = ListSort(n1)
		return n1
	} else if n1 == nil {
		return n2
	} else {
		return n1
	}
}
