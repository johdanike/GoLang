package main

import "github.com/01-edu/z01"

func main() {
	var index int = 0
	for index < 10 {
		z01.PrintRune(rune(index) + 48)

		index++
	}
	z01.PrintRune('\n')
}
