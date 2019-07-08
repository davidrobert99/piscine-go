package main

import "fmt"

func main() {
	Raid1a(1, 5)
}

func Raid1a(x, y int) {
	if x > 0 {
		limiteHorizontalsuperior := make([]rune, x)
		limiteHorizontalinferior := make([]rune, x)
		middle := make([]rune, x)
		for i := 0; i < x; i++ {
			if i == 0 {
				limiteHorizontalsuperior[i] = '/'
			} else if i == (x - 1) {
				limiteHorizontalsuperior[i] = '\\'
			} else {
				limiteHorizontalsuperior[i] = '*'
			}
		}

		for i := 0; i < x; i++ {
			if i == 0 {
				limiteHorizontalinferior[i] = '\\'
			} else if i == (x - 1) {
				limiteHorizontalinferior[i] = '/'
			} else {
				limiteHorizontalinferior[i] = '*'
			}
		}

		for i := 0; i < x; i++ {
			if i == 0 || i == (x-1) {
				middle[i] = '*'
			} else {
				middle[i] = ' '
			}
		}

		for i := 0; i < y; i++ {
			if i == 0 {
				fmt.Println(string(limiteHorizontalsuperior))
			} else if i == (y - 1) {
				fmt.Println(string(limiteHorizontalinferior))
			} else {
				fmt.Println(string(middle))
			}
		}
	}
}
