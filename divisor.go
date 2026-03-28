package piscine

func Divisor(n int) int {
	if n < 0 {
		return 0
	}

	slice := GetFactors(n)
	count := CountValues(slice)

	return count
}

func GetFactors(num int) []int {
	values := []int{}
	for index := 1; index <= num; index++ {
		if num%index == 0 {
			values = append(values, index)
		}
	}
	return values
}

func CountValues(nbr []int) int {
	count := 0
	for index := 0; index < len(nbr); index++ {
		count++
	}

	return count
}
