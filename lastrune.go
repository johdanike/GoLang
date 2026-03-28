package piscine

func LastRune(s string) rune {
	str := []rune(s)
	var result rune

	for index := len(str) - 1; index >= 0; index-- {

		result += str[index]
		break
	}
	return result
}
