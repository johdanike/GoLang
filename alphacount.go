package piscine

func AlphaCount(s string) int {
	word := ""
	for _, index := range s {
		if index >= 'a' && index <= 'z' || index >= 'A' && index <= 'Z' {
			word += string(index)
		}
	}

	return len(word)
}
