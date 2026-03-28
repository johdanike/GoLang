package piscine

func Rot14(s string) string {
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	newWord := ""

	for index := 0; index < len(s); index++ {
		isAlpha := false

		if s[index] >= 'a' && s[index] <= 'z' {
			for count := 0; count < len(lowercase); count++ {
				if s[index] == lowercase[count] {
					newWord += string(lowercase[(count+14)%len(lowercase)])
					isAlpha = true
					break
				}
			}
		}

		if s[index] >= 'A' && s[index] <= 'Z' {
			for count := 0; count < len(uppercase); count++ {
				if s[index] == uppercase[count] {
					newWord += string(uppercase[(count+14)%len(uppercase)])
					isAlpha = true
					break
				}
			}
		}

		if !isAlpha {
			newWord += string(s[index])
		}
	}

	return newWord
}
