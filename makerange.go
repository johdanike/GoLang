package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	difference := max - min
	result := make([]int, difference)

	for index := 0; index < difference; index++ {
		result[index] = min + index
	}

	return result
}
