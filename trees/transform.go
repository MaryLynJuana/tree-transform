package trees

func getValue(tree *Tree) int {
	if tree == nil { return 0 }
	return tree.Value
}

func max(a, b, c int) int {
	maxValue := a
	if b > maxValue {
		maxValue = b
	}
	if c > maxValue {
		return c
	}
	return maxValue
}

func TransformTree(tree *Tree) {
	tree.Value = max(
		getValue(tree.Left), getValue(tree.Mid), getValue(tree.Right),
	)
}
