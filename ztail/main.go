package main

import "os"

func atoi(s string) int {
	n := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return -1
		}
		n = n*10 + int(ch-'0')
	}
	return n
}

func main() {
	args := os.Args

	if len(args) < 4 || args[1] != "-c" {
		os.Stderr.Write([]byte("Usage: -c <count> <file1> [file2 ...]\n"))
		return
	}

	n := atoi(args[2])
	if n <= 0 {
		os.Stderr.Write([]byte("Invalid byte count\n"))
		return
	}

	files := args[3:]
	errorOccurred := false

	for i, fileName := range files {
		data, err := os.ReadFile(fileName)
		if err != nil {
			os.Stdout.Write([]byte(err.Error() + "\n"))
			errorOccurred = true
			continue
		}

		if len(files) > 1 {
			if i > 0 {
				os.Stdout.Write([]byte("\n"))
			}
			os.Stdout.Write([]byte("==> " + fileName + " <==\n"))
		}

		if len(data) > n {
			data = data[len(data)-n:]
		}
		os.Stdout.Write(data)
	}

	if errorOccurred {
		os.Exit(1)
	}
}
