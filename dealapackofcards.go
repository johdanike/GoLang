package piscine

import "github.com/01-edu/z01"

func printNumberNoCast(n int) {
	if n >= 10 {
		printNumberNoCast(n / 10)
	}

	digit := n % 10

	for index, value := range "0123456789" {
		if index == digit {
			z01.PrintRune(value)
		}
	}
}

func DealAPackOfCards(deck []int) {
	players := 4
	cardsPerPlayer := 3

	for index := 0; index < players; index++ {
		for _, value := range "Player " {
			z01.PrintRune(value)
		}

		printNumberNoCast(index + 1)

		for _, value := range ": " {
			z01.PrintRune(value)
		}

		for count := 0; count < cardsPerPlayer; count++ {
			card := deck[(index*cardsPerPlayer)+count]

			printNumberNoCast(card)

			if count < cardsPerPlayer-1 {
				for _, value := range ", " {
					z01.PrintRune(value)
				}
			}
		}

		z01.PrintRune('\n')
	}
}
