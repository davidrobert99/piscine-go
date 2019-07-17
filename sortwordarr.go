package piscine

func SortWordArr(array []string) {
	n := len(array)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if array[i-1] > array[i] {
				array[i], array[i-1] = array[i-1], array[i]
				swapped = true
			}
		}
	}

}
