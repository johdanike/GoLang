package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	lengthofarray := len(a)
	for firstchar := 0; firstchar < lengthofarray; firstchar++ {
		for secondchar := firstchar + 1; secondchar < lengthofarray; secondchar++ {
			if f(a[firstchar], a[secondchar]) > 0 {
				temp := a[firstchar]
				a[firstchar] = a[secondchar]
				a[secondchar] = temp
			}
		}
	}
}
