package piscine

func ToLower(s string) string {
	toLowerCasedWord := ""
	for index := 0; index < len(s); index++ {
		if s[index] >= 'A' && s[index] <= 'Z' {
			toLowerCasedWord += string(s[index] + 32)
		} else {
			toLowerCasedWord += string(s[index])
		}
	}

	return toLowerCasedWord
}
