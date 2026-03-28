package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	isfound := false

	for index := 0; index < len(s); index++ {
		if s[index] == '-' && !isfound {
			sign = -1
		}

		if s[index] >= '0' && s[index] <= '9' {
			isfound = true
			digit := int(s[index] - '0')
			result = result*10 + digit
		}

	}

	return result * sign
}
