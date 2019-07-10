package piscine

import (
	"github.com/01-edu/z01"
)

func EightQueens() {
	var NaoPodeTabuleiro [8]int
	var solucao [8]int
	solucao = [...]int{0, 0, 0, 0, 0, 0, 0, 0}
	for i := 1; i <= 8; i++ {
		solucao[0] = i
		NaoPodeTabuleiro[0] = i
		verificaSolucoes(NaoPodeTabuleiro, solucao, 1)
	}
}

func verificaSolucoes(vetorNaoPode [8]int, VetorNumeroAtual [8]int, numeroAtual int) {
	if numeroAtual == 8 { //Ultima posicao do array
		imprime := true
		for i := 0; i <= 7; i++ {
			if VetorNumeroAtual[i] == 0 {
				imprime = false
			}
		}
		if imprime {
			final(VetorNumeroAtual)
		}
	} else {
		var auxiliar [8]int
		auxiliar = criaVetorAuxiliar(vetorNaoPode, numeroAtual)
		for i := 0; i <= 8; i++ {
			if verificaSePodeInserir(auxiliar, i) {
				VetorNumeroAtual[numeroAtual] = i
				vetorNaoPode[numeroAtual] = i
				verificaSolucoes(vetorNaoPode, VetorNumeroAtual, numeroAtual+1)
			}
		}
	}

}

func final(resposta [8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(resposta[i] + 48))
	}
	z01.PrintRune('\n')
}

func verificaSePodeInserir(vetorNaoPode [8]int, numero int) bool {
	for i := 0; i < 8; i++ {
		if vetorNaoPode[i] == numero {
			return false
		}
	}
	return true
}

func criaVetorAuxiliar(vetorNaoPode [8]int, posicao int) [8]int {
	var novo [8]int //Diagonal Descendente
	novo = [...]int{0, 0, 0, 0, 0, 0, 0, 0}
	var novo2 [8]int //Diagonal Ascendente
	novo2 = [...]int{0, 0, 0, 0, 0, 0, 0, 0}
	var final [8]int
	final = [...]int{0, 0, 0, 0, 0, 0, 0, 0}
	for e := 0; e <= 7; e++ {
		if vetorNaoPode[e] != 0 {
			novo[e] = vetorNaoPode[e]
			novo2[e] = vetorNaoPode[e]
			final[e] = vetorNaoPode[e]
		}
	}
	for e := 0; e <= posicao; e++ {
		if novo[e] != 0 {
			if ((posicao - e) + vetorNaoPode[e]) <= 8 {
				novo[e] = (posicao - e) + vetorNaoPode[e]
			} else {
				novo[e] = 0
			}
		}

	}
	for e := 0; e <= posicao; e++ {
		if novo2[e] != 0 {
			if (-(posicao - e) + vetorNaoPode[e]) >= 0 {
				novo2[e] = -(posicao - e) + vetorNaoPode[e]
			} else {
				novo2[e] = 0
			}
		}

	}
	for e := 0; e <= 7; e++ {
		if final[e] == 0 {
			for i := 0; i <= 7; i++ {
				if !numeroEstaVetor(final, novo[i]) && e < 8 {
					final[e] = novo[i]
					e++
				}
				if !numeroEstaVetor(final, novo2[i]) && e < 8 {
					final[e] = novo2[i]
					e++
				}
			}
		}
	}
	return final
}

func numeroEstaVetor(vetor [8]int, numero int) bool {
	for i := 0; i < 8; i++ {
		if vetor[i] == numero {
			return true
		}
	}
	return false
}
