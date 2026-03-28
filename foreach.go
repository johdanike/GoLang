package piscine

// import "github.com/01-edu/z01"

func ForEach(f func(int), a []int) {
	for _, value := range a {
		f(value)
	}
}

// func PrintNbr(n int) {
// 	if n == 0 {
// 		z01.PrintRune('0')
// 		return
// 	}
// 	if n < 0 {
// 		z01.PrintRune('-')
// 		n = -n
// 	}

// 	var digits []int
// 	for n > 0 {
// 		digits = append([]int{n % 10}, digits...)
// 		n /= 10
// 	}

// 	for _, d := range digits {
// 		z01.PrintRune(rune('0' + d))
// 	}
// }
