package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	l = ListPushBackNode(l, data_ref)
	l = ListSort(l)
	return l
}

func ListPushBackNode(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data, Next: nil}
	if l == nil {
		return n
	}
	a := l
	for a.Next != nil {
		a = a.Next
	}
	a.Next = n
	return l
}
