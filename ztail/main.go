package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) > 1 {
		for i := 2; i < len(arguments); i++ {
			file, err := os.Open(arguments[i])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				fi, err := file.Stat()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
					// nao obteve o ficheiro
				} else {
					valor, _ := strconv.Atoi(arguments[1])
					if valor < 0 {
						valor = valor * -1
					}
					corte := int(fi.Size()) - valor
					if corte < 0 {
						corte = 0
					}
					arr := make([]byte, fi.Size())
					file.Read(arr)
					if len(arguments) == 3 {
						fmt.Print(string(arr[corte:]))
					} else if valor != 0 {
						fmt.Print("==> ")
						fmt.Print(arguments[i])
						fmt.Println(" <==")
						fmt.Print(string(arr[corte:]))
					}

				}
				file.Close()
			}
		}
	} else { //nao foi passado argumentos
		os.Exit(1)
	}
}
