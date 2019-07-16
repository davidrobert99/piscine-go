package main

import (
	"fmt"
	"os"
)

/*
func check(e error) {
	if e != nil {
		panic(e)
	}
}*/

func main() {
	arguments := os.Args
	if len(arguments) > 2 {
		fmt.Println("Too many arguments")
	} else if len(arguments) == 1 {
		fmt.Println("File name missing")
	} else {
		/*
			dat, err := ioutil.ReadFile(arguments[1])
			check(err)
			fmt.Println(string(dat))
		*/
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
