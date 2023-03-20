package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	l = ListPushBack(l, data_ref)
	l = ListSort(l)
	return l
}
