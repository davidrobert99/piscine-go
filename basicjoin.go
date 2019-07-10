package piscine

func BasicJoin(strs []string) string {
	if len(strs) > 0 {
		s := strs[0]
		for i := 1; i < len(strs); i++ {
			s = Concat(s, strs[i])
		}
		return s
	}
	return "erro"
}
