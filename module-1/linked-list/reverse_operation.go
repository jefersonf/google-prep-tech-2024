package module1

func (list *singlyLinkedList) Reverse() {
	list.head = reverse(list.head)
}

func reverse(head *listNode) *listNode {
	var previousNode *listNode
	currentNode := head
	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
	return previousNode
}
