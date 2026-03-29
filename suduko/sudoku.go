package main

import (
	"fmt"
	"os"
)

func printErrorAndExit() {
	fmt.Println("Error")
	os.Exit(0)
}

func parseArgs(args []string) ([9][9]int, bool) {
	var board [9][9]int
	if len(args) != 9 {
		return board, false
	}
	for i := 0; i < 9; i++ {
		row := args[i]
		if len(row) != 9 {
			return board, false
		}
		for j := 0; j < 9; j++ {
			c := row[j]
			if c == '.' {
				board[i][j] = 0
			} else if c >= '1' && c <= '9' {
				board[i][j] = int(c - '0')
			} else {
				return board, false
			}
		}
	}
	return board, true
}

func valid(board *[9][9]int, r, c, val int) bool {
	for i := 0; i < 9; i++ {
		if board[r][i] == val {
			return false
		}
		if board[i][c] == val {
			return false
		}
	}
	br := (r / 3) * 3
	bc := (c / 3) * 3
	for i := br; i < br+3; i++ {
		for j := bc; j < bc+3; j++ {
			if board[i][j] == val {
				return false
			}
		}
	}
	return true
}

func findEmpty(board *[9][9]int) (int, int, bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func boardIsInitiallyValid(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			v := board[i][j]
			if v != 0 {
				board[i][j] = 0
				if !valid(board, i, j, v) {
					board[i][j] = v
					return false
				}
				board[i][j] = v
			}
		}
	}
	return true
}

func solve(board *[9][9]int) bool {
	r, c, found := findEmpty(board)
	if !found {
		return true
	}
	for val := 1; val <= 9; val++ {
		if valid(board, r, c, val) {
			board[r][c] = val
			if solve(board) {
				return true
			}
			board[r][c] = 0
		}
	}
	return false
}

func printBoard(board *[9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(board[i][j])
		}
		fmt.Println()
	}
}

func main() {
	args := os.Args[1:]
	board, ok := parseArgs(args)
	if !ok {
		printErrorAndExit()
	}
	if !boardIsInitiallyValid(&board) {
		printErrorAndExit()
	}
	if !solve(&board) {
		printErrorAndExit()
	}
	printBoard(&board)
}
