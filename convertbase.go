package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	lengthOfBaseFrom := len(baseFrom)
	decimalValue := 0

	for index := 0; index < len(nbr); index++ {
		runeIndex := 0

		for count := 0; count < lengthOfBaseFrom; count++ {
			if baseFrom[count] == nbr[index] {
				runeIndex = count
				break
			}
		}

		decimalValue = decimalValue*lengthOfBaseFrom + runeIndex
	}

	lengthOfBaseTo := len(baseTo)
	var result string

	if decimalValue == 0 {
		return string(baseTo[0])
	}

	for decimalValue > 0 {
		remainder := decimalValue % lengthOfBaseTo
		result = string(baseTo[remainder]) + result
		decimalValue /= lengthOfBaseTo
	}

	return result
}
