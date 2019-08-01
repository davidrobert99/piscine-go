package main

import (
	"fmt"
	"os"
)

func estaVetor(vetor []rune, verifica rune) bool {
	for i := range vetor {
		if vetor[i] == verifica {
			return false
		}
	}
	return true
}

func main() {
	arguments := os.Args
	if len(arguments) == 3 {
		primeira := []rune(arguments[1])
		segunda := []rune(arguments[2])
		resposta := make([]rune, 0)
		for i := 0; i < len(primeira); i++ {
			for j := i + 1; j < len(segunda); j++ {
				if primeira[i] == segunda[j] {
					resposta = append(resposta, primeira[i])
				}
			}
		}
		aux := make([]rune, 0)
		for i := range resposta {
			if estaVetor(aux, resposta[i]) {
				aux = append(aux, resposta[i])
			}
		}
		fmt.Println(string(aux))

	} else {
		fmt.Println("")
	}
}
