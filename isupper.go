package piscine

func IsUpper(s string) bool {
	value := true
	for _, chars := range s {
		if chars < 'A' || chars > 'Z' {
			value = false
			break
		} else {
			value = true
		}
	}

	return value
}
