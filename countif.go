package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0

	// for _, s := range tab {
	for index := 0; index < len(tab); index++ {
		if f(tab[index]) == true {
			count++
		}
	}

	return count
}
