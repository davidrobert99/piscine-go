package piscine

import (
	"fmt"
	"os"
)

func Raid3() {
	arguments := os.Args
	for i := range arguments {
		fmt.Print(arguments[i])
	}
}
