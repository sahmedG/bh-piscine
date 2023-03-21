package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	replacement := node
	if replacement.Parent == nil {
		root = rplc
	} else if replacement == replacement.Parent.Left {
		replacement.Parent.Left = rplc
	} else {
		replacement.Parent.Right = rplc
	}
	if rplc != nil {
		rplc.Parent = replacement.Parent
	}
	return root
}
