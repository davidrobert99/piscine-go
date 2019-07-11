package piscine

func ConcatParams(args []string) string {
	answer := ""
	for i := 0; i < len(args); i++ {
		answer = Concat(answer, args[i])
		if i != len(args)-1 {
			answer = Concat(answer, "\n")
		}
	}
	return answer
}
