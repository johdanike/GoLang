package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false
	start := 0
	length := len(args)

	if length > 0 && args[0] == "--upper" {
		upper = true
		start = 1
	}

	if length > start {
		for i := start; i < length; i++ {
			number := stringToInt(args[i])

			if number >= 1 && number <= 26 {
				if upper {
					z01.PrintRune(rune('A' + number - 1))
				} else {
					z01.PrintRune(rune('a' + number - 1))
				}
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func stringToInt(str string) int {
	number := 0
	if len(str) == 0 {
		return -1
	}
	for _, r := range str {
		if r < '0' || r > '9' {
			return -1
		}
		number = number*10 + int(r-'0')
	}
	return number
}
