package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}

	for index := 0; index < len(s)-len(toFind); index++ {
		isFound := true

		for count := 0; count < len(toFind); count++ {
			if s[index+count] != toFind[count] {
				isFound = false
				break
			}
		}
		if isFound {
			return index
		}
	}
	return -1
}
