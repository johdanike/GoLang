package piscine

func IsLower(s string) bool {
	value := true
	for _, chars := range s {
		if chars < 'a' || chars > 'z' {
			value = false
			break
		} else {
			value = true
		}
	}

	return value
}
