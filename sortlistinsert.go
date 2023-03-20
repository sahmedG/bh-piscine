package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	l = ListPushBack(l, data_ref)
	l = ListSort(l)
	return l
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
}

func ListSort(l *NodeI) *NodeI {
	Head := l
	if Head == nil {
		return nil
	}
	Head.Next = ListSort(Head.Next)
	if Head.Next != nil && Head.Data > Head.Next.Data {
		Head = exchange(Head)
	}
	return Head
}

func exchange(l *NodeI) *NodeI {
	a := l
	b := l.Next
	c := b
	for b != nil && l.Data > b.Data {
		a = b
		b = b.Next
	}
	a.Next = l
	l.Next = b
	return c
}
