package main

import (
	"os"
)

func main() {
	arg0 := os.Args[0:]

	for index := len(arg0) - 1; index > 0; index-- {
		// for _, r := range arg0 {
		os.Stdout.Write([]byte(arg0[index]))
		os.Stdout.Write([]byte{'\n'})
	}
}
