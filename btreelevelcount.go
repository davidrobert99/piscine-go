package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		left := 1 + BTreeLevelCount(root.Left)
		right := 1 + BTreeLevelCount(root.Right)

		if left > right {
			return left
		}
		return right
	}
} /* if (lDepth > rDepth)
    return(lDepth + 1);
else return(rDepth + 1);*/
