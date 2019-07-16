package piscine

func CountIf(f func(string) bool, tab []string) int {
	vet := make([]bool, len(tab))
	total := 0
	for i := 0; i < len(tab); i++ {
		vet[i] = f(tab[i])
		if vet[i] == true {
			total++
		}
	}
	return total
}
