package piscine

import "fmt"

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	node := &TreeNode{Left: nil, Right: nil, Parent: nil, Data: data}
	if root == nil {
		return node
	} else {

		auxiliar := root
		anterior := &TreeNode{}

		for auxiliar != nil {
			fmt.Println(data)
			fmt.Println(auxiliar)
			if auxiliar.Data > data {
				fmt.Println("left")
				anterior = auxiliar
				auxiliar = auxiliar.Left
			} else {
				fmt.Println("right")
				anterior = auxiliar
				auxiliar = auxiliar.Right
			}
		}
		auxiliar = node
		node.Parent = anterior
		return root

	}
}
