package piscine

func StrRev(s string) string {
	var reversedString string = ""

	for index := len(s) - 1; index >= 0; index-- {
		reversedString += string(s[index])
	}

	return reversedString
}
