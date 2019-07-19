package piscine

//Dar sort e verificar 2 a 2
func Unmatch(arr []int) int {
	for i := range arr {
		if !verifica(arr, i, arr[i]) {
			return arr[i]
		}
	}
	return -1
}

func verifica(arr []int, posicao, numero int) bool {
	for i := range arr {
		if i != posicao {
			if arr[i] == numero {
				return true
			}
		}
	}
	return false
}
