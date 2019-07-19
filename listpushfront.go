package piscine

func ListPushFront(l *List, data interface{}) {
	if l.Head == nil {
		n := &NodeL{Data: data, Next: nil}
		l.Head = n
	} else {
		n := &NodeL{Data: data, Next: l.Head}
		l.Head = n
	}

}
