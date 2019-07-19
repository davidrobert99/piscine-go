package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = n
	} else {
		copia := l.Head
		for copia.Next != nil {
			copia = copia.Next
		}
		copia.Next = n
		l.Tail = n
	}
}
