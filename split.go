package piscine

import (
	"strings"
)

func Split(str, charset string) []string {
	straux := strings.Replace(str, charset, " ", -1)
	return SplitWhiteSpaces(straux)
}
