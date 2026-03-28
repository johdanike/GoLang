package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		if n/10 != 0 {
			PrintNbr(-(n / 10))
		}
		z01.PrintRune(rune('0' - n%10))
	} else {
		if n/10 != 0 {
			PrintNbr(n / 10)
		}
		z01.PrintRune(rune('0' + n%10))
	}
}
