package main

import (
	"os"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		// No file provided
		os.Stdout.Write([]byte("File name missing\n"))
		return
	}

	if len(args) > 2 {
		// More than one argument
		os.Stdout.Write([]byte("Too many arguments\n"))
		return
	}

	// Exactly one argument
	fileName := args[1]
	file, err := os.Open(fileName)
	if err != nil {
		os.Stdout.Write([]byte("ERROR: " + err.Error() + "\n"))
		return
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if n > 0 {
			os.Stdout.Write(buf[:n])
		}
		if err != nil {
			break
		}
	}
}
