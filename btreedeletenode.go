package piscine

/*
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
							} else {
								percorre.Left = nil
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
							} else {
								percorre.Right = nil
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

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node != nil {
		if root == node {
			apagador := root
			if root.Right != nil {
				aux := root.Right
				if aux.Left == nil {
					root = root.Right
					root.Parent = nil
					apagador = nil
					return root
				}
				troca := BTreeMin(aux)
				auxiliar := root.Data
				root.Data = troca.Data
				troca.Data = auxiliar
				aux = troca.Parent
				aux.Left = troca.Left
				if aux.Left != nil {
					aux.Left.Parent = aux
				}
				apagador.Parent = nil
				apagador = nil
			} else if root.Left != nil {
				root = root.Left
				root.Parent = nil
				apagador.Parent = nil
				apagador = nil
			} else {
				root = nil
				apagador.Parent = nil
				apagador = nil
			}
			return root
		}
		itera := root
		apagador := itera
		for itera != nil {
			if itera == node && itera != root {
				apagador = itera
				if itera.Right != nil {
					aux := itera.Right
					if aux.Left == nil {
						aux := itera.Parent
						tempLeft := itera.Left
						if itera.Left != nil {
							itera.Left.Parent = itera.Left
						}
						itera = itera.Right
						esquerda := itera.Left
						for esquerda != nil && esquerda.Left != nil {
							esquerda = esquerda.Left
						}
						if esquerda == nil {
							esquerda = itera
						}
						esquerda.Left = tempLeft
						if itera != nil {
							itera.Parent = aux
						}
						if aux.Left != nil && aux.Left == node {
							aux.Left = itera
						} else {
							aux.Right = itera
						}
						apagador.Parent = nil
						apagador = nil
						return root
					}
					troca := BTreeMin(aux)
					auxiliar := itera.Data
					itera.Data = troca.Data
					troca.Data = auxiliar
					aux = troca.Parent
					aux.Left = troca.Right
					if aux.Left != nil {
						aux.Left.Parent = aux
					}
					troca.Parent = nil
					troca = nil
				} else if itera.Left != nil {
					aux := itera.Left.Data
					itera.Data = aux
					tempLeft := itera.Right.Left
					tempRight := itera.Right.Right
					if tempLeft != nil {
						tempLeft.Parent = itera
					}
					if tempRight != nil {
						tempRight.Parent = itera
					}
					itera.Right = tempRight
					itera.Left = tempLeft
					apagador.Parent = nil
					apagador = nil
				} else {
					p := itera.Parent
					if p.Left != nil && p.Left.Data == node.Data {
						p.Left = nil
					} else {
						p.Right = nil
					}
					apagador.Parent = nil
					apagador = nil
					return root
				}
				return root
			} else if node.Data < itera.Data {
				itera = itera.Left
			} else {
				itera = itera.Right
			}
		}
	}
	return root
}
