package piscine

import "fmt"

func main() {
	slice := []string{
		"word1",
		"word2",
		"word3",
	}

	for index, word := range slice {
		fmt.Printf("The index is: %v; the element is: %v", index, word)
	}
}
