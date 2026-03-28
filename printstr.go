package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	// fmt.Print(s)

	for _, word := range s {
		z01.PrintRune(word)
	}
}
