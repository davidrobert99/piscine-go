package piscine

import "strings"

func Split(str, charset string) []string {
	str = strings.Replace(str, charset, " ", -1)
	return SplitWhiteSpaces(str)
}
