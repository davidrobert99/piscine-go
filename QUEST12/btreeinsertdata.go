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

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	auxiliar := root
	if auxiliar != nil {
		f(auxiliar.Data)
		BTreeApplyInorder(auxiliar.Left, f)
		BTreeApplyInorder(auxiliar.Right, f)
	}
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return root
	}
	auxiliar := root
	for auxiliar != nil {
		if auxiliar.Data == elem {
			return auxiliar
		}
		if auxiliar.Data < elem {
			auxiliar = auxiliar.Right
		} else if auxiliar.Data > elem {
			auxiliar = auxiliar.Left
		}
	}
	return nil
}

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		left := 1 + BTreeLevelCount(root.Left)
		right := 1 + BTreeLevelCount(root.Right)

		if left > right {
			return left
		}
		return right
	}
} /* if (lDepth > rDepth)
    return(lDepth + 1);
else return(rDepth + 1);*/

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		if root.Parent != nil && root.Parent.Left == root {
			if root.Parent.Data < root.Data {
				return false
			} else {
				return true && BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
			}
		} else if root.Parent != nil && root.Parent.Right == root {
			if root.Parent.Data > root.Data {
				return false
			} else {
				return true && BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
			}
		}
		return true && BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
	}
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	for i := 1; i <= BTreeLevelCount(root); i++ {
		Level(root, i, f)
	}
}

func Level(root *TreeNode, nivel int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if nivel == 1 {
		f(root.Data)
	} else if nivel > 1 {
		Level(root.Right, nivel-1, f)
		Level(root.Left, nivel-1, f)
	}
}

func BTreeMax(root *TreeNode) *TreeNode {
	auxiliar := root
	for auxiliar.Right != nil {
		auxiliar = auxiliar.Right
	}
	return auxiliar
}

func BTreeMin(root *TreeNode) *TreeNode {
	auxiliar := root
	for auxiliar.Left != nil {
		auxiliar = auxiliar.Left
	}
	return auxiliar
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	auxiliar := root
	for auxiliar != nil {
		if auxiliar.Data == node.Data {
			if auxiliar.Parent != nil {
				auxiliar = auxiliar.Parent
				if auxiliar.Left.Data == node.Data {
					auxiliar.Left = rplc
				} else {
					auxiliar.Right = rplc
				}
			} else {
				auxiliar = rplc
			}
		}
		if node.Data < auxiliar.Data {
			auxiliar = auxiliar.Left
		} else {
			auxiliar = auxiliar.Right
		}
	}
	return root
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	percorre := root
	if root.Data == node.Data {
		root = root.Left
		auxiliar := root.Left
		if auxiliar != nil {
			for auxiliar.Right != nil {
				auxiliar = auxiliar.Right
			}
		}
		root.Right = percorre.Right
	} else {
		for percorre != nil {
			if percorre.Data == node.Data {
				if percorre.Parent != nil {
					percorre = percorre.Parent
					if percorre.Left.Data == node.Data {
						auxiliarApagar := percorre.Left //para apagar
						percorre.Left = percorre.Left.Left
						auxiliar := percorre.Left
						if auxiliar != nil {
							for auxiliar.Right != nil {
								auxiliar = auxiliar.Right
							}
						}
						auxiliar = auxiliarApagar.Right
					} else {
						auxiliarApagar := percorre.Right //para apagar
						percorre.Right = percorre.Right.Left
						auxiliar := percorre.Right
						if auxiliar != nil {
							for auxiliar.Right != nil {
								auxiliar = auxiliar.Right
							}
						}
						auxiliar = auxiliarApagar.Right
					}
				}
			}
			if node.Data < percorre.Data {
				percorre = percorre.Left
			} else {
				percorre = percorre.Right
			}
		}
	}
	return root
}
