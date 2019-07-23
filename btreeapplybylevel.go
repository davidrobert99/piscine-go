package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	aux := BTreeLevelCount(root)
	for i := 1; i <= aux; i++ {
		Level(root, i, f)
	}
}

func Level(root *TreeNode, nivel int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if nivel == 1 {
		f(root.Data)
	} else {
		Level(root.Left, nivel-1, f)
		Level(root.Right, nivel-1, f)
	}
}
