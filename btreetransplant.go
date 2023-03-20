package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	repl := node
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		repl.Parent.Left = rplc
	} else if node == node.Parent.Right {
		repl.Parent.Right = rplc
	}
	repl.Parent = node.Parent
	return root
}
