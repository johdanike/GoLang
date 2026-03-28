package piscine

func ListReverse(l *List) {
	if l == nil || l.Head == nil {
		return
	}

	var prev *NodeL = nil
	currentNode := l.Head

	l.Tail = l.Head

	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = prev

		prev = currentNode
		currentNode = nextNode
	}

	l.Head = prev
}
