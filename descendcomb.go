package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for a1 := '9'; a1 >= '0'; a1-- {
		for a2 := '9'; a2 >= '0'; a2-- {
			for b1 := '9'; b1 >= '0'; b1-- {
				for b2 := '9'; b2 >= '0'; b2-- {
					if a1 > b1 || (a1 == b1 && a2 > b2) {
						continue
					}

					z01.PrintRune(a1)
					z01.PrintRune(a2)
					z01.PrintRune(' ')

					z01.PrintRune(b1)
					z01.PrintRune(b2)

					if a1 != '0' || a2 != '1' || b1 != '0' || b2 != '0' {

						z01.PrintRune(',')

						z01.PrintRune(' ')

					} else {
						z01.PrintRune('\n')
					}
				}
			}
		}
	}
}

func main() {
	DescendComb()
}
