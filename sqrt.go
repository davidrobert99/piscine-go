package piscine

func Sqrt(nb int) int {
	if nb > 0 {
		valorAux := 0
		valorfinal := valorAux
		for valorAux*valorAux <= nb {
			valorfinal = valorAux
			valorAux++
		}
		if valorfinal*valorfinal == nb {
			return valorfinal
		} else {
			return 0
		}
	} else {
		return 0
	}
}
