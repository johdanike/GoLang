package piscine

import "fmt"

func QuadA(x, y int) {
	for row := 1; row <= y; row++ { // print row

		for col := 1; col <= x; col++ { // print col

			if (col == 1 || col == x) && (row == 1 || row == y) {
				fmt.Print("o")
			} else if row == 1 || row == y {
				fmt.Print("-")
			} else if col == 1 || col == x {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()

	}
}
