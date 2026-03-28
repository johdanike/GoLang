package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	comb := make([]int, n)
	printCombRecursive(comb, 0, 0, n)
	z01.PrintRune('\n')
}

func printCombRecursive(comb []int, index, start, n int) {
	if index == n {
		for i := 0; i < n; i++ {
			z01.PrintRune(rune('0' + comb[i]))
		}
		if comb[0] != 10-n {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}
	for i := start; i <= 9; i++ {
		if 10-i >= n-index {
			comb[index] = i
			printCombRecursive(comb, index+1, i+1, n)
		}
	}
}
