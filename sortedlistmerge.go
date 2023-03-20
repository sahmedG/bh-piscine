package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	node1 := ListSort(n1)
	node2 := ListSort(n2)
	if node1 == node2 {
		return node2
	}
	if node2 == nil {
		return node1
	}
	if node1.Data <= node2.Data {
		node1.Next = SortedListMerge(node1.Next, node2)
		return node1
	} else {
		node2.Next = SortedListMerge(node1, node2.Next)
		return node2
	}
}
