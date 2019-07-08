package piscine

func StrLen(str string) int {
	size := 0
	var str2 = []rune(str)
	for i := 0; i < len(str2); i++ {
		size++
	}
	return size
}
