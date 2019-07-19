package piscine

import (
	"fmt"
	"os"
)

func comcheck() {
	arguments := os.Args
	boolean := false
	for i := 1; i < len(arguments); i++ {
		if arguments[i] == "01" || arguments[i] == "galaxy" || arguments[i] == "galaxy 01" {
			boolean = true
		}
	}
	if boolean {
		fmt.Print("Alert!!!")
	}
	fmt.Print("\n")
}
