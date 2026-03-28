package piscine

func CamelToSnakeCase(sentence string) string {
	letter := ""

	for index := 0; index < len(sentence); index++ {
		if sentence[index] >= 'A' && sentence[index] <= 'Z' {
			if index != 0 {
				letter += "_"
			}
			letter += string(sentence[index] + 32)
		} else {
			letter += string(sentence[index])
		}
	}

	return letter
}
