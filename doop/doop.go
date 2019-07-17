package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 4 {
		num1, er := strconv.Atoi(arguments[1])
		num2, er2 := strconv.Atoi(arguments[3])
		if er != nil || er2 != nil {
			fmt.Println(1)
		} else {
			if arguments[2] == "+" {
				result := num1 + num2
				fmt.Println(result)
			} else if arguments[2] == "-" {
				result := num1 - num2
				fmt.Println(result)
			} else if arguments[2] == "*" {
				result := num1 * num2
				fmt.Println(result)
			} else if arguments[2] == "%" {
				if num2 != 0 {
					result := num1 % num2
					fmt.Println(result)
				} else {
					fmt.Println("No Modulo by 0")
				}
			} else if arguments[2] == "/" {
				if num2 != 0 {
					result := num1 / num2
					fmt.Println(result)
				} else {
					fmt.Println("No division by 0")
				}
			} else {
				fmt.Println(0)
			}
		}
	}
}

func IsNumeric(str string) bool {
	aux := []rune(str)
	for i := 0; i < len(aux); i++ {
		if !(aux[i] >= 48 && aux[i] <= 57) {
			return false
		}
	}
	return true
}
