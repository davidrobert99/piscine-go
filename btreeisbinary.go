package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		if root.Parent != nil && root.Parent.Left == root {
			if root.Parent.Data < root.Data {
				return false
			} else {
				return true && BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
			}
		} else if root.Parent != nil && root.Parent.Right == root {
			if root.Parent.Data > root.Data {
				return false
			} else {
				return true && BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
			}
		}
		return true && BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
	}
}
