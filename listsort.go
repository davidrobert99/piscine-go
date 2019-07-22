package piscine

import "sort"

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	var numbers []int
	percorre := l
	for percorre != nil {
		numbers = append(numbers, percorre.Data)
		percorre = percorre.Next
	}
	//SortIntegerTable(numbers)
	sort.Ints(numbers)
	l = nil
	for i := 0; i < len(numbers); i++ {
		l = listPushBack(l, numbers[i])
	}
	return l
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

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
