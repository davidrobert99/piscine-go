package piscine

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	n := len(array)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if f(array[i-1], array[i]) > 0 {
				array[i], array[i-1] = array[i-1], array[i]
				swapped = true
			}
		}
	}

}
