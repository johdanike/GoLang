package piscine

func RepeatAlpha(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		word := GetAlpha(s, i)
		occuredNumberOfTimes := GetIndexFromAlphaSet(word)
		result += PrintAlphaSequence(occuredNumberOfTimes, word)
	}

	return result
}

func GetIndexFromAlphaSet(str string) int {

	indexValueInByte := 0
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for index := 0; index < 26; index++ {

		if !(str[index] >= 'a' && str[index] <= 'z' || str[index] >= 'A' && str[index] <= 'Z') {
			indexValueInByte += index + 1
		}

		if str == string(lower[index]) || str == string(upper[index]) {
			indexValueInByte += index + 1
		}

	}

	return indexValueInByte

}

func GetAlpha(str string, index int) string {
	return string(str[index])
}

func PrintAlphaSequence(num int, str string) string {

	newWord := ""
	for index := 0; index < num; index++ {
		newWord += string(str[index])
	}

	return newWord

}
