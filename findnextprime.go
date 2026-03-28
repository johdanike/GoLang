package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}

	for index := nb; ; index++ {
		isPrime := true
		for count := 2; count*count <= index; count++ {
			if index%count == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			return index
		}
	}
}
