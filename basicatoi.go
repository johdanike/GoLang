package piscine

func BasicAtoi(s string) int {
	var result int = 0

	index := 0
	for index < len(s) {
		result = result*10 + int(s[index]-48)
		index++
	}

	return result
}
