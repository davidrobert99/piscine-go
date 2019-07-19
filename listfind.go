package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	copia := l.Head
	for copia != nil {
		if comp(copia.Data, ref) {
			return &copia.Data
		} else {
			copia = copia.Next
		}
	}
	return nil
}
