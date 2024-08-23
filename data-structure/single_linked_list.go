package datastructure

import "fmt"

type List struct {
	Head *node
}

type node struct {
	data any
	Next *node
}

func (n *node) Data() any {
	return n.data
}

func NewSingleLinkedList() *List {
	return &List{}
}

func newNode(data any) *node {
	return &node{data, nil}
}

func (list *List) InsertBegin(data any) {
	if list.Head == nil {
		list.Head = newNode(data)
		return
	}
	newHead := newNode(data)
	newHead.Next = list.Head
	list.Head = newHead
}

func (list *List) InsertEnd(data any) {
	if list.Head == nil {
		list.Head = newNode(data)
		return
	}
	currentTail := list.Head
	for currentTail.Next != nil {
		currentTail = currentTail.Next
	}
	currentTail.Next = newNode(data)
}

func (list *List) InsertMiddle(index int, data any) {
	if index == 0 {
		list.InsertBegin(data)
		return
	}
	currentIndex := 0
	currentNode := list.Head
	for currentNode != nil && currentIndex < index-1 {
		currentIndex += 1
		currentNode = currentNode.Next
	}
	middleNode := newNode(data)
	nextNode := currentNode.Next
	currentNode.Next = middleNode
	middleNode.Next = nextNode
}

func (list *List) RemoveHead() {
	if list == nil {
		return
	}
	if list.Head != nil {
		list.Head = list.Head.Next
	}
}

func (list *List) RemoveTail() {
	if list == nil {
		return
	}

	if list.Head != nil && list.Head.Next == nil {
		list.RemoveHead()
		return
	}
	tail := list.Head
	for tail.Next.Next != nil {
		tail = tail.Next
	}
	tail.Next = nil
}

func (list *List) RemoveMiddle(index int) {
	if list == nil || list.Head == nil {
		return
	}
	if index == 0 {
		list.RemoveHead()
		return
	}

	currentIndex := 0
	currentNode := list.Head
	for currentNode != nil && currentIndex < index-1 {
		currentIndex += 1
		currentNode = currentNode.Next
	}

	if currentNode.Next == nil {
		list.RemoveTail()
		return
	}

	if currentNode.Next.Next == nil {
		currentNode.Next = nil
		return
	}

	currentNode.Next = currentNode.Next.Next
}

func Print(list *List) {
	if list.Head == nil {
		return
	}
	currentNode := list.Head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.Next
	}
}
