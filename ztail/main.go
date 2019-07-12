
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(char)
	} else {
		for i := 1; i < len(arguments); i++ {
			dat, err := ioutil.ReadFile(arguments[i])
			check(err)
			fmt.Println(string(dat))
			fmt.Println("")
		}
	}
}
