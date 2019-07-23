package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	auxiliar := root
	for auxiliar.Right != nil {
		auxiliar = auxiliar.Right
	}
	return auxiliar
}
