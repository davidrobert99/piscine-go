package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else {
		percorre := root
		for percorre != nil && node != nil {
			if percorre.Data == node.Data {
				if percorre.Parent != nil { // se nao é a raiz
					percorre = percorre.Parent
					if percorre.Left != nil && percorre.Left.Data == node.Data { // se é o no da esquerda para apagar
						auxiliarApagar := percorre.Left //para apagar
						if percorre.Left.Left != nil {
							percorre.Left = percorre.Left.Left
							auxiliar := percorre.Left
							auxiliar.Parent = percorre
							if auxiliarApagar.Right != nil {
								for auxiliar.Right != nil {
									auxiliar = auxiliar.Right
								}
								auxiliar.Right = auxiliarApagar.Right
								auxiliar.Right.Parent = auxiliar
							}

						} else {
							if percorre.Left.Right != nil {
								percorre.Left = percorre.Left.Right
								percorre.Left.Parent = percorre //////////////////////////////
							}
						}
						auxiliarApagar = nil
					} else { // se é o no da direita para apagar
						auxiliarApagar := percorre.Right //para apagar
						if percorre.Right.Left != nil {
							percorre.Right = percorre.Right.Left
							auxiliar := percorre.Right
							auxiliar.Parent = percorre
							if auxiliarApagar.Right != nil {
								for auxiliar.Right != nil {
									auxiliar = auxiliar.Right
								}
								auxiliar.Right = auxiliarApagar.Right
								auxiliar.Right.Parent = auxiliar
							}
						} else {
							if percorre.Right.Right != nil {
								percorre.Right = percorre.Right.Right
								percorre.Right.Parent = percorre
							}
						}
						auxiliarApagar = nil
					}
					return root
				} else { // se e a raiz
					if root.Right != nil { //se à direita nao é nil
						if root.Left != nil { //se à esquerda nao é nil
							guardaDireita := root.Right
							root = root.Left
							root.Parent = nil
							auxiliar := root.Right
							if auxiliar != nil {
								for auxiliar.Right != nil {
									auxiliar = auxiliar.Right
								}
								auxiliar.Right = guardaDireita
								auxiliar.Right.Parent = auxiliar
							} else {
								root.Right = guardaDireita
								root.Right.Parent = root
							}

						} else {
							root = root.Right
							root.Parent = nil
						}
					} else {
						root = root.Left
						root.Parent = nil
					}
					return root
				}
			} else if node.Data < percorre.Data {
				percorre = percorre.Left
			} else {
				percorre = percorre.Right
			}
		}
	}
	return root
}

/*

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else { //fazer caso
		percorre := root
		for percorre != nil {
			if percorre.Data == node.Data {
				if percorre.Left != nil {
					if percorre.Right == nil {
						if percorre.Parent.Right.Data == percorre.Data {
							percorre = percorre.Parent
							percorre.Right = percorre.Right.Left
							percorre.Right.Parent = percorre
						} else {
							percorre = percorre.Parent
							percorre.Left = percorre.Left.Left
							percorre.Left.Parent = percorre
						}
					} else {
						if percorre.Left.Right != nil {
							auxiliar := percorre.Left.Right
							for auxiliar.Right.Right != nil {
								auxiliar = auxiliar.Right
							}
							percorre.Data = auxiliar.Right.Data
							auxiliar.Right = nil
						} else {
							percorre.Data = percorre.Left.Data
							percorre.Left = nil
						}
					}
				} else if percorre.Right != nil {
					if percorre.Parent.Right.Data == percorre.Data {
						percorre = percorre.Parent
						percorre.Right = percorre.Right.Right
						percorre.Right.Parent = percorre
					} else {
						percorre = percorre.Parent
						percorre.Left = percorre.Left.Right
						percorre.Left.Parent = percorre
					}
				} else {
					percorre = nil
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
}*/
