package piscine

func CollatzCountdown(start int) int {
	valor := 1
	for start != 1 {
		if start%2 == 0 {
			start = start / 2
			valor++
		} else {
			start = start*3 + 1
			valor++
		}
	}
	return valor
}
