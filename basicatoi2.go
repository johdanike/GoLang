package piscine

func BasicAtoi2(s string) int {
	result := 0
	index := 0

	for index < len(s) {
		if s[index] < '0' || s[index] > '9' {
			return 0
		}
		result = result*10 + int(s[index]-'0')
		index++
	}

	return result
}
