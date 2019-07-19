package piscine

func ListSize(l *List) int {
	if l.Head == nil {
		return 0
	} else {
		copia := l.Head
		size := 1
		for copia.Next != nil {
			copia = copia.Next
			size++
		}
		return size
	}
}
