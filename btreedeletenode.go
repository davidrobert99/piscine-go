package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else {
		percorre := root
		for percorre != nil {
			if percorre.Data == node.Data {
				if percorre.Parent != nil { // se nao é a raiz
					percorre = percorre.Parent
					if percorre.Left.Data == node.Data {
						auxiliarApagar := percorre.Left //para apagar
						if percorre.Left.Left != nil {
							percorre.Left = percorre.Left.Left
							auxiliar := percorre.Left
							if auxiliar != nil {
								for auxiliar.Right != nil {
									auxiliar = auxiliar.Right
								}
							}
							auxiliar = auxiliarApagar.Right
							auxiliarApagar = nil
						} else {
							percorre.Left = percorre.Left.Right
						}

					} else {
						auxiliarApagar := percorre.Right //para apagar
						if percorre.Right.Left != nil {
							percorre.Right = percorre.Right.Left
							auxiliar := percorre.Right
							if auxiliar != nil {
								for auxiliar.Right != nil {
									auxiliar = auxiliar.Right
								}
							}
							auxiliar = auxiliarApagar.Right
							auxiliarApagar = nil
						} else {
							percorre = percorre.Left.Left
						}
					}
				} else { // se e a raiz
					if root.Right != nil {
						if root.Left != nil {
							root = root.Left
							auxiliar := root.Left
							if auxiliar != nil {
								for auxiliar.Right != nil {
									auxiliar = auxiliar.Right
								}
							}
							root.Right = percorre.Right
						} else {
							root = root.Right
						}
					} else {
						root = root.Left
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
