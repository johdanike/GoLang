package piscine

func Concat(str1 string, str2 string) string {
	newword1 := ""
	newword2 := ""

	for string1 := 0; string1 < len(str1); string1++ {
		newword1 += string(str1[string1])
	}

	for string2 := 0; string2 < len(str2); string2++ {
		newword2 += string(str2[string2])
	}

	word := newword1 + newword2

	return word
}
