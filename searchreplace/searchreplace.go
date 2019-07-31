package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 4 {
		aux1 := []rune(arguments[2])
		aux2 := []rune(arguments[3])
		resposta := Split(arguments[1], string(aux1[0]), string(aux2[0]))
		fmt.Println(resposta)
	} else {
		fmt.Println("")
	}
}

func Split(str, s, replace string) string {
	return strings.Replace(str, s, replace, -1)
}
