package main

import (
	"os"
)

func main() {
	arg0 := os.Args[1:]

	for _, r := range arg0 {
		os.Stdout.Write([]byte(r))
		os.Stdout.Write([]byte{'\n'})
	}
}
