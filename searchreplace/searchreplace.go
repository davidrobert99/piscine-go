package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 4 {
		resposta := Split(arguments[1], string(arguments[2][0]), string(arguments[3][0]))
		fmt.Println(resposta)
	} else {
		fmt.Println("")
	}
}

func Split(str, s, replace string) string {
	return strings.Replace(str, s, replace, -1)
}
