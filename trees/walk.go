package trees

func WalkTree(tree *Tree, transformFunc func(*Tree)) {
	if tree == nil { return }
	transformFunc(tree)
	go WalkTree(tree.Left, transformFunc)
	go WalkTree(tree.Mid, transformFunc)
	go WalkTree(tree.Right, transformFunc)
}