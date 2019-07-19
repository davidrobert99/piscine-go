package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	percorre := n1
	for percorre.Next != nil {
		percorre = percorre.Next
	}
	percorre.Next = n2
	n1 = ListSort(n1)
	return n1
}
