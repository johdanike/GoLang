package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for index, value := range a {
		result[index] = f(value)
	}

	return result
}
