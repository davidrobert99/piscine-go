package main

import "fmt"

func main() {
	Raid1a(-4, -4)
}

func Raid1a(x, y int) {
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
	}
}
