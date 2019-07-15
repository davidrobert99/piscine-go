package piscine

import (
	"fmt"
	"os"
)

func Sudoku() {
	arguments := os.Args
	if len(arguments) > 10 {

	} else if len(arguments) < 10 {

	} else {
		carateresCorretos := true
		board := [9][9]int{}
		for i := 1; i < 10; i++ {
			for j := 0; j < 9; j++ {
				if arguments[i][j] >= '1' && arguments[i][j] <= '9' {
					board[i-1][j] = int(arguments[i][j] - 48)
				} else if arguments[i][j] != '.' {
					carateresCorretos = false
				}
			}
		}
		if carateresCorretos {
			Solucao(board, 0, 0)

		}
	}
}

func print(board [9][9]int) {
	for j := 0; j < 9; j++ {
		fmt.Println(board[j])
	}
}

func podeColocar(coluna, linha, numero int, board [9][9]int) bool {
	if board[coluna][linha] != 0 {
		return false
	}
	for i := 0; i < 9; i++ {
		if board[linha][i] == numero {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if board[i][coluna] == numero {
			return false
		}
	}
	for i := linha - linha%3; i <= linha-linha%3+2; i++ {
		for j := coluna - coluna%3; j <= coluna-coluna%3+2; j++ {
			if board[i][j] == numero {
				return false
			}
		}
	}
	return true
}

func Solucao(board [9][9]int, coluna, linha int) bool {
	if matrizCheia(board) {
		print(board)
		return true
	} else {
		print(board)
		fmt.Println("\n")
		for i := 1; i <= 9; i++ {
			if podeColocar(coluna, linha, i, board) {
				board[linha][coluna] = i
				coluna++
				fmt.Print("Linha: ")
				fmt.Println(linha)
				fmt.Print("Coluna: ")
				fmt.Println(coluna)
				if coluna == 8 {
					coluna = 0
					linha++
				}
				if Solucao(board, coluna, linha) {
					return true
				}

				board[linha][coluna] = 0
			} else {
				coluna++
				fmt.Print("Linha: ")
				fmt.Println(linha)
				fmt.Print("Coluna: ")
				fmt.Println(coluna)
				if coluna == 8 {
					coluna = 0
					linha++
				}

				if Solucao(board, coluna, linha) {
					return true
				}
			}
		}
		return false
	}
}

func matrizCheia(board [9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
