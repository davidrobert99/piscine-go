package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		reader := bufio.NewReader(os.Stdin)
		char, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Print(char)
	} else {
		for i := 1; i < len(arguments); i++ {

			file, err := os.Open(arguments[i])
			if err != nil {
				fmt.Println(err.Error())
				break
			} else {

				fi, erro := file.Stat()
				if erro != nil {
					fmt.Println(err.Error())
					break
				} else {
					arr := make([]byte, fi.Size())
					file.Read(arr)
					fmt.Println(string(arr))
					file.Close()
				}

			}
		}
	}
}

/*
	dat, err := ioutil.ReadFile(arguments[i])
	check(err)
	fmt.Println(string(dat))
	fmt.Println("")
*/
