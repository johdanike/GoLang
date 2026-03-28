package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg0 := os.Args[1:]

	for index := 0; index < len(arg0)-1; index++ {
		for inner := index + 1; inner < len(arg0); inner++ {
			if arg0[index] > arg0[inner] {
				arg0[index], arg0[inner] = arg0[inner], arg0[index]
			}
		}
	}

	for _, value := range arg0 {
		for _, result := range value {
			z01.PrintRune(result)
		}
		z01.PrintRune('\n')
	}
}
