package piscine

func NRune(s string, n int) rune {
	if len(s) == 0 {
		return 0
	}

	var result rune
	str := []rune(s)
	for index := 0; index < len(str); index++ {
		if n-1 == index {
			result += str[index]
			break
		}
	}

	return result
}
