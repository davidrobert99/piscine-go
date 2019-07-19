package piscine

func ActiveBits(n int) uint {
	number := PrintNbrBaseSimples(n, "01")
	valor := 0
	for i := range number {
		if number[i] == '1' {
			valor++
		}
	}
	return uint(valor)
}
