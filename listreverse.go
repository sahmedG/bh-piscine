package piscine

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}
	t := l.Head
	c := l.Head
	prev := nil
	for c != nil {
		next := c.next
		c.Next = prev
		prev = c
		c = next
	}
	l.Head = prev
	l.Tail = t
	l.Tail.Next = nil
}
