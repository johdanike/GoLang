package piscine

func IsPrime(nb int) bool {
	if nb < 0 {
		return false
	}

	if nb <= 1 {
		return false
	}

	for index := 2; index*index <= nb; index++ {
		if nb%index == 0 {
			return false
		}
	}

	return true
}
