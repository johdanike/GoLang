package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}

	result := 1

	for index := 1; index <= power; index++ {
		result *= nb
	}

	return result
}
