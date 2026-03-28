package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)

	start := 0

	for i := 0; i < len(str); i++ {
		// Every time we hit a space, extract the word (even if it's empty)
		if str[i] == ' ' {
			word := str[start:i]
			summary[word]++

			// Move the start index to the character immediately after the space
			start = i + 1
		}
	}

	// Catch the final word after the last space
	word := str[start:]
	summary[word]++

	return summary
}
