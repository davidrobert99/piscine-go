package main

/*
type List = piscine.List
type Node = piscine.NodeL

func PrintElem(node *piscine.NodeL) {
	fmt.Println(node.Data)
}

func StringToInt(node *piscine.NodeL) {
	node.Data = 2
}

func PrintList(l *piscine.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, "->")
		it = it.Next
	}
	fmt.Println(nil)
}

func PrintList2(l *piscine.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}
func listPushBack(l *piscine.NodeI, data int) *piscine.NodeI {
	n := &piscine.NodeI{Data: data}

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

func main() {
	var link *piscine.NodeI
	var link2 *piscine.NodeI

	link = listPushBack(link, 3)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)

	link2 = listPushBack(link2, -2)
	link2 = listPushBack(link2, 9)

	PrintList2(piscine.SortedListMerge(link2, link))
}*/

import (
	"fmt"

	piscine ".."
)

func main() {
	arr := []int{1, 2, 3, 1, 2, 3, 4, 1}
	unmatch := piscine.Unmatch(arr)
	fmt.Println(unmatch)
}
