package piscine

import (
	"github.com/01-edu/z01"
	// "fmt"
)

func PrintWordsTables(a []string) {
	for index := 0; index < len(a); index++ {
		word := a[index]

		for count := 0; count < len(word); count++ {
			z01.PrintRune(rune(word[count]))
		}

		z01.PrintRune('\n')
	}
}
