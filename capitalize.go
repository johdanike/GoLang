package piscine

func Capitalize(s string) string {
	res := []rune(s)
	newWord := true

	for index := 0; index < len(res); index++ {
		if isAlphanumeric(res[index]) {
			if newWord {
				if res[index] >= 'a' && res[index] <= 'z' {
					res[index] -= 32
				}
				newWord = false
			} else {
				if res[index] >= 'A' && res[index] <= 'Z' {
					res[index] += 32
				}
			}
		} else {
			newWord = true
		}
	}
	return string(res)
}

func isAlphanumeric(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
		return true
	}
	return false
}
