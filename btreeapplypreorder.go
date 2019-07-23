package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	auxiliar := root
	if auxiliar != nil {
		f(auxiliar.Data)
		BTreeApplyInorder(auxiliar.Left, f)
		BTreeApplyInorder(auxiliar.Right, f)
	}
}
