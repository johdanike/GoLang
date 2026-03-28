package piscine

func Join(strs []string, sep string) string {
	withcolon := sep
	fullword := ""

	for element := 0; element < len(strs); element++ {

		fullword += strs[element]

		if element < len(strs)-1 {
			fullword += withcolon
		}

	}

	return fullword
}
