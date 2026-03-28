package piscine

func SplitWhiteSpaces(s string) []string {
	var slices []string
	word := ""
	for index := 0; index < len(s); index++ {
		if ' ' == s[index] || '\n' == s[index] {
			if word != "" {
				slices = append(slices, word)
				word = ""
			}
		} else {
			word += string(s[index])
		}
	}

	if word != "" {
		slices = append(slices, word)
	}

	return slices
}
