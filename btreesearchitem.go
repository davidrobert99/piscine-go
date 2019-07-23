package piscine

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
