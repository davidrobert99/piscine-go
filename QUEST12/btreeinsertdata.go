package piscine

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
			if auxiliar.Data > data {
				anterior = auxiliar
				auxiliar = auxiliar.Left
			} else {
				anterior = auxiliar
				auxiliar = auxiliar.Right
			}
		}
		if anterior.Data > data {
			anterior.Left = node
			node.Parent = anterior
		} else {
			anterior.Right = node
			node.Parent = anterior
		}
		return root

	}
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	auxiliar := root
	if auxiliar != nil {
		BTreeApplyInorder(auxiliar.Left, f)
		f(auxiliar.Data)
		BTreeApplyInorder(auxiliar.Right, f)
	}
}

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	auxiliar := root
	if auxiliar != nil {
		BTreeApplyInorder(auxiliar.Left, f)
		BTreeApplyInorder(auxiliar.Right, f)
		f(auxiliar.Data)
	}
}
