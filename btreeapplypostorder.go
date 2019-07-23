package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	auxiliar := root
	if auxiliar != nil {
		BTreeApplyPostorder(auxiliar.Left, f)
		BTreeApplyPostorder(auxiliar.Right, f)
		f(auxiliar.Data)
	}

}
