package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(result rune) bool {
	vowels := "aeiouAEIOU"
	for _, val := range vowels {
		if result == val {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		z01.PrintRune('\n')
		return
	}

	var fullPath []rune
	for index, arg := range args {
		for _, result := range arg {
			fullPath = append(fullPath, result)
		}
		if index < len(args)-1 {
			fullPath = append(fullPath, ' ')
		}
	}

	var vowelIndices []int
	for index, result := range fullPath {
		if isVowel(result) {
			vowelIndices = append(vowelIndices, index)
		}
	}

	for index, j := 0, len(vowelIndices)-1; index < j; index, j = index+1, j-1 {
		index1 := vowelIndices[index]
		index2 := vowelIndices[j]
		fullPath[index1], fullPath[index2] = fullPath[index2], fullPath[index1]
	}

	for _, result := range fullPath {
		z01.PrintRune(result)
	}
	z01.PrintRune('\n')
}
