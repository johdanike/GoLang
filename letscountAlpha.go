package piscine

func CountAllAlpha(str string) int {

	count := 0

	for index := 0; index < len(str); index++ {
		if str[index] >= 'a' && str[index] <= 'z' || str[index] >= 'A' && str[index] <= 'Z' {
			count++
		}
	}

	return count
}
