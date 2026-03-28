package piscine

func SortWordArr(a []string) {
	lengthofarray := len(a)
	for firstchar := 0; firstchar < lengthofarray; firstchar++ {
		for secondchar := firstchar + 1; secondchar < lengthofarray; secondchar++ {
			if a[firstchar] > a[secondchar] {
				temp := a[firstchar]
				a[firstchar] = a[secondchar]
				a[secondchar] = temp
			}
		}
	}
}
