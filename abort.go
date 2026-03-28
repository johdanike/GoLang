package piscine

func Abort(a, b, c, d, e int) int {
	val1 := a
	val2 := b
	val3 := c
	val4 := d
	val5 := e

	numbers := [5]int{val1, val2, val3, val4, val5}

	for value := 0; value < len(numbers); value++ {
		for index := value + 1; index < len(numbers); index++ {
			if numbers[value] > numbers[index] {
				temp := numbers[value]
				numbers[value] = numbers[index]
				numbers[index] = temp
			}
		}
	}

	return numbers[2]
}
