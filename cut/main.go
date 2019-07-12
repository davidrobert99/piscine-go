package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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
			fmt.Println(err)
		}
		fmt.Print(char)
	} else {
		for i := 1; i < len(arguments); i++ {
			dat, err := ioutil.ReadFile(arguments[i])
			check(err)
			fmt.Println(string(dat))
			fmt.Println("")
		}
	}
}