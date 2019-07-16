package piscine

func Map(f func(int) bool, arr []int) []bool {
	vet := make([]bool, len(arr))
	for i := 0; i < len(arr); i++ {
		vet[i] = f(arr[i])
	}
	return vet
}
