package piscine

import "fmt"

func Raid1d(x, y int) {
	if x > 0 {
		limiteHorizontalsuperior := make([]rune, x)
		limiteHorizontalinferior := make([]rune, x)
		middle := make([]rune, x)
		for i := 0; i < x; i++ {
			if i == 0 {
				limiteHorizontalsuperior[i] = 'A'
			} else if i == (x - 1) {
				limiteHorizontalsuperior[i] = 'C'
			} else {
				limiteHorizontalsuperior[i] = 'B'
			}
		}

		for i := 0; i < x; i++ {
			if i == 0 {
				limiteHorizontalinferior[i] = 'A'
			} else if i == (x - 1) {
				limiteHorizontalinferior[i] = 'C'
			} else {
				limiteHorizontalinferior[i] = 'B'
			}
		}

		for i := 0; i < x; i++ {
			if i == 0 || i == (x-1) {
				middle[i] = 'B'
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
