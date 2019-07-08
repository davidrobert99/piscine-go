package piscine

func FindNextPrime(nb int) int {
	if nb > 0 {
		encontrou := false
		i := nb
		valor := nb
		for !encontrou {
			valor = i
			encontrou = IsPrime(i)
			i++
			if i < 0 {
				return 0
			}
		}
		return valor
	} else {
		return 2
	}
}
