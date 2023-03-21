package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			t := root.Right
			root = nil
			return t
		} else if root.Right == nil {
			t := root.Left
			root = nil
			return t
		}
		t := piscine.BTreeMin(root.Right)
		root.Data = t.Data
		root.Right = BTreeDeleteNode(root.Right, t)
	}
	return root
}
