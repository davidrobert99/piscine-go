package piscine

import "fmt"

/*
type rainha struct {
	x int
	y int
}

func EightQueens() {
	rainha1 := rainha{3, 5}
	var tabuleiro [8][8]int
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			tabuleiro[i][j] = 0
		}
	}
	tabuleiro = posiciona(tabuleiro, rainha1)
	imprimeTabuleiro(tabuleiro)
}

func imprimeTabuleiro(tabuleiro [8][8]int) {
	for i := 0; i < 8; i++ {
		fmt.Println(tabuleiro[i])
	}
}

func posiciona(tabuleiro [8][8]int, r rainha) [8][8]int {
	x1 := r.x
	y1 := r.y
	for x1 != 0 && y1 != 0 {
		x1--
		y1--
	}
	x2 := r.x
	y2 := r.y
	fmt.Print(x2)
	fmt.Println(y2)
	for y2 != 7 && x2 != 0 {

		x2--
		y2++
		fmt.Print(x2)
		fmt.Println(y2)
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if i == x1 && j == y1 {
				tabuleiro[i][j] = 2
				x1++
				y1++
			}
			fmt.Print(x2)
			fmt.Print(y2)
			fmt.Print(i)
			fmt.Println(j)
			if i == x2 && j == y2 {
				tabuleiro[i][j] = 2
				x2++
				y2--
			}
			if i == (r.x - 1) {
				tabuleiro[i][j] = 2
			}
			if j == (r.y - 1) {
				tabuleiro[i][j] = 2
			}
		}
	}
	return tabuleiro
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}*/

func EightQueens() {
	var NaoPodeTabuleiro [8]int
	var solucao [8]int
	solucao = [...]int{0, 0, 0, 0, 0, 0, 0, 0}
	for i := 1; i <= 1; i++ { //<=8
		solucao[0] = i
		NaoPodeTabuleiro[0] = i
		verificaSolucoes(NaoPodeTabuleiro, solucao, 1)
	}
}

func verificaSolucoes(vetorNaoPode [8]int, VetorNumeroAtual [8]int, numeroAtual int) {
	if numeroAtual == 8 { //Ultima posicao do array
		//fmt.Println(VetorNumeroAtual)
	} else {
		for i := 1; i <= 3; i++ { //<=8
			var vetorAuxiliarNaoPode = criaVetorAuxiliar(vetorNaoPode, numeroAtual)
			fmt.Print("Aux ")
			fmt.Print(vetorAuxiliarNaoPode)
			fmt.Print(" atual")
			fmt.Println(VetorNumeroAtual)
			if verificaSePodeInserir(vetorAuxiliarNaoPode, i) {
				VetorNumeroAtual[numeroAtual] = i
				vetorNaoPode[numeroAtual] = i
				verificaSolucoes(vetorNaoPode, VetorNumeroAtual, numeroAtual+1)
			}
		}
	}

}

func verificaSePodeInserir(vetorNaoPode [8]int, numero int) bool {
	for i := 0; i < 8; i++ {
		if vetorNaoPode[i] == numero {
			return false
		}
	}
	return true
}

/*
func criaVetorAuxiliar(vetorNaoPode [8]int, numeroAtual int) [8]int {
	i := numeroAtual
	if i <= numeroAtual {
		if !numeroEstaVetor(vetorNaoPode, (numeroAtual-i)+vetorNaoPode[i]) {
			if (numeroAtual-i)+vetorNaoPode[i] < 9 {
				vetorNaoPode[i] = (numeroAtual - i) + vetorNaoPode[i]
			}
			if -(numeroAtual-i)+vetorNaoPode[i] > 0 {
				vetorNaoPode[i] = -(numeroAtual - i) + vetorNaoPode[i]
			}
		}

	}

	return vetorNaoPode
}*/

func criaVetorAuxiliar(vetorNaoPode [8]int, numeroAtual int) [8]int {
	for i := 0; i < 8; i++ {
		if vetorNaoPode[i] != 0 && i <= numeroAtual {
			/*
				if !numeroEstaVetor(vetorNaoPode, vetorNaoPode[i] +) {

				}*/
			vetorNaoPode[numeroAtual] = vetorNaoPode[i] - numeroAtual
			fmt.Println(vetorNaoPode[i] - numeroAtual)
			fmt.Println(vetorNaoPode[i] + numeroAtual)
		}
	}
	return vetorNaoPode
}

func numeroEstaVetor(vetor [8]int, numero int) bool {
	for i := 0; i < 8; i++ {
		if vetor[i] == numero {
			return true
		}
	}
	return false
}
