package main

import "os"

func main() {
	args := os.Args

	if len(args) == 1 {
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf)
			if n > 0 {
				os.Stdout.Write(buf[:n])
			}
			if err != nil {
				break
			}
		}
		return
	}

	for _, fileName := range args[1:] {
		data, err := os.ReadFile(fileName)
		if err != nil {
			os.Stdout.Write([]byte("ERROR: " + err.Error() + "\n"))
			os.Exit(1)
		}
		os.Stdout.Write(data)
	}
}
