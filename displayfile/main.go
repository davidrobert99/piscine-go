package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) > 2 {
		fmt.Println("Too many arguments")
	} else if len(arguments) == 1 {
		fmt.Println("File name missing")
	} else {
		file, e := os.Open(arguments[1])
		if e != nil {
			fmt.Println(e.Error())
		} else {
			fi, _ := file.Stat()
			arr := make([]byte, fi.Size())
			file.Read(arr)
			fmt.Println(string(arr))
			file.Close()
		}

	}
}
