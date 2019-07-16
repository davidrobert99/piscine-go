package piscine

import "fmt"

func ForEach(f func(int), arr []int) {
	for i := 0; i < len(arr); i++ {
		f(arr[i])
		if i == len(arr)-1 {
			fmt.Println("")
		}
	}
}
