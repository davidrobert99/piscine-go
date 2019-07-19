package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	} else {
		copia := l.Head
		size := 1
		for copia.Next != nil {
			copia = copia.Next
			size++
		}
		return copia.Data
	}
}
