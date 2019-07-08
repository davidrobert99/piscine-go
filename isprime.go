package piscine

func IsPrime(nb int) bool {
	numerosDivisores := 0
	for i := 1; i <= nb; i++ {
		if nb%i == 0 {
			numerosDivisores++
		}
	}
	if numerosDivisores == 2 {
		return true
	} else {
		return false
	}
}
