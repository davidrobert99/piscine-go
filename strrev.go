package piscine

func StrRev(s string) string {
	reverse := []rune(s)
	for i, j := 0, len(reverse)-1; i < len(reverse)/2; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}
	return string(reverse)
}
