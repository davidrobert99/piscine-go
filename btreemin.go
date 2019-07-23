package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	auxiliar := root
	for auxiliar.Left != nil {
		auxiliar = auxiliar.Left
	}
	return auxiliar
}
