package piscine

func LoafOfBread(str string) string {
	mystring := ""

	if str == "" {
		return "\n"
	}

	if len(str) < 5 {
		return "invalid output\n"
	}

	f := 0

	for _, count := range str {
		if count != ' ' && f != 5 {
			mystring += string(count)
			f++
		} else if f == 5 {
			mystring += " "
			f = 0
		}
	}

	if len(mystring) > 0 && mystring[len(mystring)-1] == ' ' {
		mystring = mystring[:len(mystring)-1]
	}
	return mystring + "\n"
}
