package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	Raid1a2()
}

func Raid1a2() os.FileInfo {

	arguments := os.Args
	x, _ := strconv.Atoi(arguments[1])
	y, _ := strconv.Atoi(arguments[2]) /*
		if x > 0 {
			limiteHorizontal := make([]rune, x)
			middle := make([]rune, x)
			for i := 0; i < x; i++ {
				if i == 0 || i == (x-1) {
					limiteHorizontal[i] = 'o'
				} else {
					limiteHorizontal[i] = '-'
				}
			}

			for i := 0; i < x; i++ {
				if i == 0 || i == (x-1) {
					middle[i] = '|'
				} else {
					middle[i] = ' '
				}
			}

			for i := 0; i < y; i++ {
				if i == 0 || i == (y-1) {
					fmt.Println(string(limiteHorizontal))
				} else {
					fmt.Println(string(middle))
				}
			}
		} */

	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	/*
		if fi.Mode()&os.ModeNamedPipe == 0 {
			fmt.Println("no pipe :(")
		} else {
			fmt.Println("hi pipe!")
		} */

	.00 + fmt.Print(x, y, arguments[0][2:])
	return fi
}
