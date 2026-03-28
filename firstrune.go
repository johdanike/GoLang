package piscine

func FirstRune(s string) rune {
	// var firstrune rune
	// for index := 0; index < len(s); index++ {
	// 	if index == 0 {
	// 		firstrune += rune(s[index])
	// 	}
	// }
	// return firstrune

	str := []rune(s)

	return str[0]
}
