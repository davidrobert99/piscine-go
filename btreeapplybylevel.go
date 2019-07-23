package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	for i := 1; i <= BTreeLevelCount(root); i++ {
		Level(root, i, f)
	}
}

func Level(root *TreeNode, nivel int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if nivel == 1 {
		f(root.Data)
	} else if nivel > 1 {
		Level(root.Right, nivel-1, f)
		Level(root.Left, nivel-1, f)
	}
}
