package piscine

import "sort"

func Unmatch(arr []int) int {
	sort.Ints(arr)
	for i := 0; i < len(arr); i += 2 {
		if i == len(arr)-1 {
			return arr[i]
		}
		if arr[i] != arr[i+1] {
			return arr[i]
		}
	}
	return -1
}

/*
func verifica(arr []int, posicao, numero int) bool {
	for i := range arr {
		if i != posicao {
			if arr[i] == numero {
				return true
			}
		}
	}
	return false
}*/
