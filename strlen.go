package piscine

func StrLen(str string) int {
	size := 0
	for i := 0; i < len(str); i++ {
		size++
	}
	return size
}
