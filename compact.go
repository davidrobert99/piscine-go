package piscine

func Compact(ptr *[]string) int {
	conta := 0
	vetorAuxiliar := *ptr
	var final []string
	for i, _ := range vetorAuxiliar {
		if vetorAuxiliar[i] != "" {
			final = append(final, vetorAuxiliar[i])
			conta++
		}
	}
	*ptr = final
	return conta
}
