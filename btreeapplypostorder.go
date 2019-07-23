package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	auxiliar := root
	if auxiliar != nil {
		BTreeApplyInorder(auxiliar.Left, f)
		BTreeApplyInorder(auxiliar.Right, f)
		f(auxiliar.Data)
	}
}
