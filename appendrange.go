package piscine

func AppendRange(min, max int) []int {
	var resposta []int
	if min < max {
		for i := 0; i < max-min; i++ {
			resposta = append(resposta, min+i)
		}
	}
	return resposta

}
