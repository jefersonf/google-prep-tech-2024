package module1

import "errors"

func (list *singlyLinkedList) InsertBegin(data any) {
	if list.head == nil {
		list.head = newNode(data)
		return
	}
	newHead := newNode(data)
	newHead.Next = list.head
	list.head = newHead
}

func (list *singlyLinkedList) InsertEnd(data any) {
	if list.head == nil {
		list.head = newNode(data)
		return
	}
	currentTail := list.head
	for currentTail.Next != nil {
		currentTail = currentTail.Next
	}
	currentTail.Next = newNode(data)
}

func (list *singlyLinkedList) Insert(index int, data any) error {
	if index == 0 {
		list.InsertBegin(data)
		return nil
	}
	currentIndex := 0
	currentNode := list.head
	for currentNode != nil && currentIndex < index-1 {
		currentIndex += 1
		currentNode = currentNode.Next
	}

	if currentNode == nil {
		return errors.New("index out of range")
	}

	node := newNode(data)
	nextNode := currentNode.Next
	currentNode.Next = node
	node.Next = nextNode

	return nil
}
