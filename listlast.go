package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}

	currentPosition := l.Head

	for currentPosition.Next != nil {
		currentPosition = currentPosition.Next
	}

	return currentPosition.Data
}
