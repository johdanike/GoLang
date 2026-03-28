package piscine

func ToUpper(s string) string {
	capitalizedWord := ""
	for index := 0; index < len(s); index++ {
		if s[index] >= 'a' && s[index] <= 'z' {
			capitalizedWord += string(s[index] - 32)
		} else {
			capitalizedWord += string(s[index])
		}
	}

	return capitalizedWord
}
