package piscine

func BasicJoin(elems []string) string {
	joined := ""

	for elements := 0; elements < len(elems); elements++ {
		joined += elems[elements]
	}

	return joined
}
