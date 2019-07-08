package main

import "fmt"

func main() {
	arg1 := 4
	arg2 := 3
	fmt.Println(IterativePower(arg1, arg2))
}

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	valor := 1
	for power > 0 {
		valor = nb * valor
		power--
	}
	return valor

}
