package piscine

import "strings"

func SplitWhiteSpaces(str string) []string {
	answer := strings.Fields(str)
	return answer
}
