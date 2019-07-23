package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	auxiliar := root
	if auxiliar != nil {
		BTreeApplyInorder(auxiliar.Left, f)
		f(auxiliar.Data)
		BTreeApplyInorder(auxiliar.Right, f)
	}
}
