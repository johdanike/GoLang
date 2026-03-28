package piscine

func ConcatParams(args []string) string {
	sentence := ""
	for index := 0; index < len(args); index++ {

		sentence += args[index]
		if index < len(args)-1 {
			sentence += "\n"
		}
	}

	return sentence
}
