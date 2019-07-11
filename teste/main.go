package main

import (
	piscine ".."
)

func main() {
	str := "Hello how are you?"
	table := piscine.SplitWhiteSpaces(str)
	piscine.PrintWordsTables(table)
}
