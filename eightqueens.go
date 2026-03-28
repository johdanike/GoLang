package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var positions [8]int
	solveQueens(0, &positions)
}

func solveQueens(col int, positions *[8]int) {
	if col == 8 {
		for i := 0; i < 8; i++ {
			z01.PrintRune(rune(positions[i] + 1 + 48))
		}
		z01.PrintRune('\n')
		return
	}

	for row := 0; row < 8; row++ {
		if isSafe(row, col, positions) {
			positions[col] = row
			solveQueens(col+1, positions)
		}
	}
}

func isSafe(row, col int, positions *[8]int) bool {
	for prevCol := 0; prevCol < col; prevCol++ {
		prevRow := positions[prevCol]

		if prevRow == row {
			return false
		}

		rowDiff := row - prevRow
		if rowDiff < 0 {
			rowDiff = -rowDiff
		}
		colDiff := col - prevCol

		if rowDiff == colDiff {
			return false
		}
	}
	return true
}
