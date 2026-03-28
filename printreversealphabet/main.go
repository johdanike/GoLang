package main

import "github.com/01-edu/z01"

func main() {
	alphabets := "abcdefghijklmnopqrstuvwxyz"

	for index := len(alphabets) - 1; index >= 0; index-- {
		z01.PrintRune(rune(alphabets[index]))
	}

	z01.PrintRune('\n')
}
