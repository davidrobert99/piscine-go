package piscine

func Abort(a, b, c, d, e int) int {
	values := [5]int{a, b, c, d, e}
	values = SortInteger(values)
	return values[2]
}

func SortInteger(table [5]int) [5]int {
	n := len(table)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if table[i-1] > table[i] {
				table[i], table[i-1] = table[i-1], table[i]
				swapped = true
			}
		}
	}
	return table
}
