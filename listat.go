package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}

	currentNode := l
	currentPosition := 0

	for currentNode != nil {

		if currentPosition == pos {
			return currentNode
		}

		currentNode = currentNode.Next
		currentPosition++

	}

	return nil
}
