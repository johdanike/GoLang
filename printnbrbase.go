package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 {
		printStr("NV")
		return
	}
	for i, char := range base {
		if char == '+' || char == '-' {
			printStr("NV")
			return
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				printStr("NV")
				return
			}
		}
	}

	if nbr < 0 {
		z01.PrintRune('-')
	} else {
		nbr = -nbr
	}

	solve(nbr, base)
}

func solve(n int, base string) {
	l := len(base)
	if n <= -l {
		solve(n/l, base)
	}

	mod := n % l
	if mod < 0 {
		mod = -mod
	}
	z01.PrintRune(rune(base[mod]))
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
