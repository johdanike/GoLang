package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var str1 []rune
	for n > 0 {
		digit := n % 10
		str1 = append(str1, rune(digit+'0'))
		n /= 10
	}

	for i := 0; i < len(str1)-1; i++ {
		for index := 0; index < len(str1)-1; index++ {
			if str1[index] > str1[index+1] {
				temp := str1[index+1]
				str1[index+1] = str1[index]
				str1[index] = temp
			}
		}
	}

	for index := 0; index < len(str1); index++ {
		z01.PrintRune(str1[index])
	}
}
