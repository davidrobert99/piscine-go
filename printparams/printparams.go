package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	for i := 1; i < len(arguments); i++ {
		fmt.Println(arguments[i])
	}
}
