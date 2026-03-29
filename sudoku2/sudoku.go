// package main

// import (
// 	"fmt"
// 	"os"
// )

// func solve(b *[9][9]int) bool {
// 	for r := 0; r < 9; r++ {
// 		for c := 0; c < 9; c++ {
// 			if b[r][c] == 0 {
// 				for v := 1; v <= 9; v++ {
// 					if valid(b, r, c, v) {
// 						b[r][c] = v
// 						if solve(b) {
// 							return true
// 						}
// 						b[r][c] = 0
// 					}
// 				}
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// func valid(b *[9][9]int, r, c, v int) bool {
// 	for i := 0; i < 9; i++ {
// 		if b[r][i] == v || b[i][c] == v || b[3*(r/3)+i/3][3*(c/3)+i%3] == v {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	var b [9][9]int
// 	args := os.Args[1:]
// 	if len(args) != 9 {
// 		fmt.Println("Error")
// 		return
// 	}

// 	for i, row := range args {
// 		if len(row) != 9 {
// 			fmt.Println("Error")
// 			return
// 		}
// 		for j, char := range row {
// 			if char >= '1' && char <= '9' {
// 				v := int(char - '0')
// 				if !valid(&b, i, j, v) {
// 					fmt.Println("Error")
// 					return
// 				}
// 				b[i][j] = v
// 			} else if char != '.' {
// 				fmt.Println("Error")
// 				return
// 			}
// 		}
// 	}

// 	if !solve(&b) {
// 		fmt.Println("Error")
// 		return
// 	}

// 	for _, row := range b {
// 		for j, v := range row {
// 			fmt.Printf("%d%s", v, map[bool]string{true: " ", false: ""}[j < 8])
// 		}
// 		fmt.Println()
// 	}
// }

package main

import (
	"fmt"
	"os"
)

// solve tries to fill the board. It returns true if it finishes the puzzle.
func solve(board *[9][9]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {                 // Found an empty spot
				for v := 1; v <= 9; v++ {     // Try numbers 1 through 9
					if valid(board, row, col, v) {
						board[row][col] = v           // Place the number
						if solve(board) {         // Try to solve the rest of the board
							return true
						}
						board[row][col] = 0           // It didn't work, erase it (backtrack)
					}
				}
				return false // No valid numbers worked here, go back
			}
		}
	}
	return true // No empty spots left, puzzle solved!
}

// valid checks if placing 'v' at row 'r', column 'c' follows Sudoku rules.
func valid(board *[9][9]int, row, col, v int) bool {
	for index := 0; index < 9; index++ {
		// 1. Check the entire row
		if board[row][index] == v {
			return false
		}
		// 2. Check the entire column
		if board[index][col] == v {
			return false
		}
		// 3. Check the 3x3 box
		boxRow := (row/3)*3 + (index/3)
		boxCol := (col/3)*3 + (index%3)
		if board[boxRow][boxCol] == v {
			return false
		}
	}
	return true
}

func main() {
	var board [9][9]int
	args := os.Args[1:]

	// Check if we have exactly 9 rows
	if len(args) != 9 {
		fmt.Println("Error")
		return
	}

	// Build the board from the arguments
	for index := 0; index < 9; index++ {
		row := args[index]
		if len(row) != 9 {
			fmt.Println("Error")
			return
		}

		for count := 0; count < 9; count++ {
			char := row[count]
			if char >= '1' && char <= '9' {
				v := int(char - '0')
				// Make sure the puzzle isn't broken from the start
				if !valid(&board, index, count, v) {
					fmt.Println("Error")
					return
				}
				board[index][count] = v
			} else if char != '.' {
				// Invalid character found (like a letter)
				fmt.Println("Error")
				return
			}
		}
	}

	// Solve the board
	if !solve(&board) {
		fmt.Println("Error")
		return
	}

	// Print the finished board clearly
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			fmt.Print(board[row][col])
			if col < 8 {
				fmt.Print(" ") // Add a space after the number, but not at the very end
			}
		}
		fmt.Println() // Go to the next line after finishing a row
	}
}

