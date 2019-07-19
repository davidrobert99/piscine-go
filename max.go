package piscine

func Max(arr []int) int {
	if len(arr) > 0 {
		maximo := arr[0]
		for i := range arr {
			if maximo < arr[i] {
				maximo = arr[i]
			}
		}
		return maximo
	}
	return 0
}
