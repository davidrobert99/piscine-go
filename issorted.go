package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	for i := 1; i < len(tab)-1; i++ {
		if f(tab[i-1], tab[i]) >= 0 {
			return false
		}
	}
	return true
}
