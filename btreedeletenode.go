package piscine

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
