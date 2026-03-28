package piscine

func AppendRange(min, max int) []int {
	var result []int

	if min > max {
		return result
	}

	for index := min; index < max; index++ {
		result = append(result, index)
	}

	return result
}
