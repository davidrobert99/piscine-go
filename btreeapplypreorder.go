package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	auxiliar := root
	if auxiliar == nil {
		return
	}
	f(auxiliar.Data)
	BTreeApplyPreorder(auxiliar.Left, f)
	BTreeApplyPreorder(auxiliar.Right, f)
}
