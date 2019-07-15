package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	Sudoku(arguments)
}
func Sudoku(arguments []string) {
	carateresCorretos := true
	if len(arguments) > 10 {
		carateresCorretos = false
	} else if len(arguments) < 10 {
		carateresCorretos = false
	} else {
		board := [9][9]int{}
		for i := 1; i < 10; i++ {
			if len(arguments[i]) == 9 {
				for j := 0; j < 9; j++ {
					if arguments[i][j] >= '1' && arguments[i][j] <= '9' {
						board[i-1][j] = int(arguments[i][j] - 48)
					} else if arguments[i][j] != '.' {
						carateresCorretos = false
					}
				}
			} else {
				carateresCorretos = false
			}
		}
		if carateresCorretos && verificaBoardInic(board) {
			if !Solucao(board) {
				fmt.Println("Error")
			}
		} else {
			carateresCorretos = false
		}
	}
	if !carateresCorretos {
		fmt.Println("Error")
	}
}

func print(board [9][9]int) {
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			if i < 8 {
				fmt.Print(board[j][i], " ")
			} else {
				fmt.Println(board[j][i])
			}
		}
	}
}

func podeColocar(coluna, linha, numero int, board [9][9]int) bool {
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

func Solucao(board [9][9]int) bool {
	coluna, linha := 0, 0
	if !posicaoVazia(board, &linha, &coluna) {
		print(board)
		return true
	}
	for i := 1; i <= 9; i++ {
		if podeColocar(coluna, linha, i, board) {
			board[linha][coluna] = i
			if Solucao(board) {
				return true
			}
			board[linha][coluna] = 0
		}
	}
	return false
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

func posicaoVazia(board [9][9]int, row, column *int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				*row = i
				*column = j
				return true
			}
		}
	}
	return false
}

func verificaBoardInic(board [9][9]int) bool {
	for l := 0; l < 9; l++ {
		for u := 0; u < 9; u++ {
			if board[l][u] != 0 {
				for i := 0; i < 9; i++ {
					if board[l][i] == board[l][u] && i != u {
						return false
					}
				}
				for i := 0; i < 9; i++ {
					if board[i][u] == board[l][u] && i != l {
						return false
					}
				}
				for i := l - l%3; i <= l-l%3+2; i++ {
					for j := u - u%3; j <= u-u%3+2; j++ {
						if board[i][j] == board[l][u] && i != l && j != u {
							return false
						}
					}
				}
			}
		}
	}
	return true
}
