package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	} else {
		res := root
		lres := BTreeMax(root.Left)
		rres := BTreeMax(root.Right)
		if lres.Data > res.Data {
			res = lres
		} else if rres.Data > res.Data {
			res = rres
		}

		return res
	}
}
