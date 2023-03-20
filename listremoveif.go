package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	temp := l.Head
	prev := l.Head
	for temp != nil && temp.Data == data_ref {
		l.Head = temp
		temp = temp.Next
	}
	if temp == nil {
		return
	}
	prev.Next = temp.Next
	temp = prev.Next
}
