package main

import "os"

// PrintString writes a string directly to standard output
func PrintString(text string) {
	os.Stdout.Write([]byte(text))
}

// PrintNumber converts an integer to its ASCII characters and prints it
func PrintNumber(number int) {
	if number == 0 {
		os.Stdout.Write([]byte{'0'})
		return
	}

	var reversedDigits []byte

	// Extract digits right-to-left
	for number > 0 {
		lastDigit := byte((number % 10) + '0')
		reversedDigits = append(reversedDigits, lastDigit)
		number /= 10
	}

	// Print the collected digits in the correct left-to-right order
	for i := len(reversedDigits) - 1; i >= 0; i-- {
		os.Stdout.Write([]byte{reversedDigits[i]})
	}
}

// GenerateQuadShape builds a perfect reference shape based on the provided dimensions and border characters
func GenerateQuadShape(width, height int, topLeft, topRight, bottomLeft, bottomRight, topEdge, bottomEdge, leftEdge, rightEdge byte) string {
	if width <= 0 || height <= 0 {
		return ""
	}

	var quadBytes []byte

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			// Descriptive booleans make the logic instantly readable
			isTopRow := (row == 0)
			isBottomRow := (row == height-1)
			isLeftCol := (col == 0)
			isRightCol := (col == width-1)

			// Determine which character to place based on the current position
			if isTopRow && isLeftCol {
				quadBytes = append(quadBytes, topLeft)
			} else if isTopRow && isRightCol {
				quadBytes = append(quadBytes, topRight)
			} else if isBottomRow && isLeftCol {
				quadBytes = append(quadBytes, bottomLeft)
			} else if isBottomRow && isRightCol {
				quadBytes = append(quadBytes, bottomRight)
			} else if isTopRow {
				quadBytes = append(quadBytes, topEdge)
			} else if isBottomRow {
				quadBytes = append(quadBytes, bottomEdge)
			} else if isLeftCol {
				quadBytes = append(quadBytes, leftEdge)
			} else if isRightCol {
				quadBytes = append(quadBytes, rightEdge)
			} else {
				quadBytes = append(quadBytes, ' ') // Inner empty space
			}
		}
		quadBytes = append(quadBytes, '\n') // End of the row
	}

	return string(quadBytes)
}

func main() {
	// 1. Read all piped data from standard input into a buffer
	var inputData []byte
	buffer := make([]byte, 1024)

	for {
		bytesRead, err := os.Stdin.Read(buffer)
		if bytesRead > 0 {
			inputData = append(inputData, buffer[:bytesRead]...)
		}
		if err != nil { // Stop reading when we reach the End Of File (EOF)
			break
		}
	}

	if len(inputData) == 0 {
		PrintString("Not a quad function\n")
		return
	}

	inputText := string(inputData)

	// 2. Calculate dimensions (Width and Height)
	width := 0
	height := 0

	for index := 0; index < len(inputText); index++ {
		if inputText[index] == '\n' {
			if height == 0 {
				width = index // The width is the number of characters before the first newline
			}
			height++
		}
	}

	// If no valid dimensions were found, exit early
	if width == 0 || height == 0 {
		PrintString("Not a quad function\n")
		return
	}

	// 3. Compare the input to our perfect generated reference shapes
	successfulMatches := 0

	// Helper function to format and print successful matches uniformly
	printMatch := func(quadName string) {
		if successfulMatches > 0 {
			PrintString(" || ") // Add separator if this isn't the first match
		}
		PrintString("[" + quadName + "] [")
		PrintNumber(width)
		PrintString("] [")
		PrintNumber(height)
		PrintString("]")

		successfulMatches++
	}

	// Generate the 5 shapes and check for perfect matches in alphabetical order
	if inputText == GenerateQuadShape(width, height, 'o', 'o', 'o', 'o', '-', '-', '|', '|') {
		printMatch("quadA")
	}
	if inputText == GenerateQuadShape(width, height, '/', '\\', '\\', '/', '*', '*', '*', '*') {
		printMatch("quadB")
	}
	if inputText == GenerateQuadShape(width, height, 'A', 'A', 'C', 'C', 'B', 'B', 'B', 'B') {
		printMatch("quadC")
	}
	if inputText == GenerateQuadShape(width, height, 'A', 'C', 'A', 'C', 'B', 'B', 'B', 'B') {
		printMatch("quadD")
	}
	if inputText == GenerateQuadShape(width, height, 'A', 'C', 'C', 'A', 'B', 'B', 'B', 'B') {
		printMatch("quadE")
	}

	// 4. Final output handling
	if successfulMatches == 0 {
		PrintString("Not a quad function\n")
	} else {
		PrintString("\n")
	}
}
