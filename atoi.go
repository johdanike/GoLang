package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}

	result := 0
	sign := 1
	i := 0

	if s[0] == '+' {
		i++
	} else if s[0] == '-' {
		sign = -1
		i++
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		result = result*10 + int(s[i]-'0')
	}

	return result * sign
}
