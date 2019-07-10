package piscine

func Join(strs []string, sep string) string {
	if len(strs) > 0 {
		s := strs[0]
		for i := 1; i < len(strs); i++ {
			if i < len(strs) {
				s = Concat(s, sep)
				s = Concat(s, strs[i])
			} else {
				s = Concat(s, strs[i])
			}
		}
		return s
	}
	return "erro"
}
