package piscine

func MakeRange(min, max int) []int {
	if min < max {
		vetor := make([]int, max-min)
		for i := 0; i < len(vetor); i++ {
			vetor[i] = min + i
		}
		return vetor
	} else {
		vetor := make([]int, 0)
		return vetor
	}
}
