package piscine

func ListClear(l *List) {
	// if l.Head == nil {
	// 	re
	// }

	currentPosition := l.Head

	for currentPosition != nil {
		nextNode := currentPosition.Next
		currentPosition.Next = nil
		currentPosition = nextNode
	}

	l.Head = nil
	l.Tail = nil
}
