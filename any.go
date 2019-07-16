package piscine

func Any(f func(string) bool, arr []string) bool {
	vet := make([]bool, len(arr))
	for i := 0; i < len(arr); i++ {
		vet[i] = f(arr[i])
		if vet[i] == true {
			return true
		}
	}
	return false
}
