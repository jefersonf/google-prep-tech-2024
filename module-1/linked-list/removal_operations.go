package module1

import (
	"errors"
	"fmt"
)

func (list *singlyLinkedList) RemoveHead() error {
	if list == nil || list.head == nil {
		return errors.New("list is empty")
	}
	list.head = list.head.Next
	return nil
}

func (list *singlyLinkedList) RemoveTail() error {
	if list == nil {
		return errors.New("list is empty")
	}

	if list.head != nil && list.head.Next == nil {
		return list.RemoveHead()
	}
	tail := list.head
	for tail.Next.Next != nil {
		tail = tail.Next
	}
	tail.Next = nil
	return nil
}

func (list *singlyLinkedList) Remove(index int) error {
	if list == nil || list.head == nil {
		return nil
	}
	if index == 0 {
		return list.RemoveHead()
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

	if currentNode.Next == nil {
		return list.RemoveTail()
	}

	if currentNode.Next.Next == nil {
		currentNode.Next = nil
		return nil
	}

	currentNode.Next = currentNode.Next.Next
	return nil
}

func Print(list *singlyLinkedList) {
	if list.head == nil {
		return
	}
	currentNode := list.head
	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}
}
