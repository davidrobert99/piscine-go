package piscine

func FindNextPrime(nb int) int {
	encontrou := false
	i := nb
	valor := nb
	for !encontrou {
		valor = i
		encontrou = IsPrime(i)
		i++
	}
	return valor
}
