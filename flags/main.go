package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || hasHelp(args) {
		printHelp()
		return
	}

	var insertStr string
	var order bool
	var targetStr string

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if len(arg) >= 9 && compare(arg[:9], "--insert=") {
			insertStr = arg[9:]
		} else if arg == "--insert" {
			if i+1 < len(args) {
				insertStr = args[i+1]
				i++
			}
		} else if len(arg) >= 3 && compare(arg[:3], "-i=") {
			insertStr = arg[3:]
		} else if arg == "-i" {
			if i+1 < len(args) {
				insertStr = args[i+1]
				i++
			}
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else {
			targetStr = arg
		}
	}

	result := targetStr + insertStr

	if order {
		result = sortString(result)
	}

	printStr(result)
	z01.PrintRune('\n')
}

func compare(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func hasHelp(args []string) bool {
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			return true
		}
	}
	return false
}

func printHelp() {
	help := "--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n" +
		"--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n"
	printStr(help)
}

func sortString(s string) string {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		for j := i + 1; j < len(r); j++ {
			if r[i] > r[j] {
				r[i], r[j] = r[j], r[i]
			}
		}
	}
	return string(r)
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
