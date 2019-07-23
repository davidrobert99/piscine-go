package piscine

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
