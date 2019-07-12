package main

import (
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
	if len(arguments) > 2 {
		fmt.Println("Too many arguments")
	} else if len(arguments) == 1 {
		fmt.Println("File name missing")
	} else {
		dat, err := ioutil.ReadFile(arguments[1])
		check(err)
		fmt.Println(string(dat))
	}
}
