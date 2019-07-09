package piscine

func IsPrime(nb int) bool {
	//numerosDivisores := 0
	if nb <= 1 {
		return false
	} else {
		if nb > 3 {
			if nb%2 == 0 || nb%3 == 0 {
				return false
			}
			for i := 5; i*i < nb; i = i + 6 {
				if nb%(i-1) == 0 || nb%(i+2) == 0 {
					return false
				}
			}
			return true
		} else {
			return true
		}
	}
	/*for i := 1; i <= nb; i++ {
		if nb%i == 0 {
			numerosDivisores++
		}
	}
	if numerosDivisores == 2 {
		return true
	} else {
		return false
	}*/
}
