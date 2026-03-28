package piscine

func Split(s, sep string) []string {
	var slices []string

	for index := 0; index+len(sep) <= len(s); index++ {
		if s[index:index+len(sep)] == sep {
			slices = append(slices, s[:index])
			s = s[index+len(sep):]

			index = -1
		}
	}

	slices = append(slices, s)
	return slices
}
