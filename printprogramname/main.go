package main

import (
	"os"
)

func main() {
	arg0 := os.Args[0]

	lastSlash := -1
	for i, char := range arg0 {
		if char == '/' {
			lastSlash = i
		}
	}

	name := arg0[lastSlash+1:]

	for _, r := range name {
		os.Stdout.Write([]byte(string(r)))
	}

	os.Stdout.Write([]byte{'\n'})
}
