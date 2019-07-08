package piscine

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
