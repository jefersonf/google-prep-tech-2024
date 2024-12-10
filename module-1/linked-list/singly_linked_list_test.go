package module1_test

import (
	module1 "preptech24/module-1/linked-list"
	"testing"
)

func TestGetHead(t *testing.T) {
	t.Run("should return an error when linked list has no nodes", func(t *testing.T) {
		sll := module1.NewSinglyLinkedList()
		if _, err := sll.Head(); err == nil {
			t.Errorf("wants no error, got %v", err)
		}
	})
	t.Run("should return the head of a non-empty singly linked list", func(t *testing.T) {
		sll := module1.NewSinglyLinkedList()
		sll.InsertBegin("A")
		sll.InsertEnd("B")
		head, err := sll.Head()
		if err != nil {
			t.Errorf("wants no error, got %v", err)
		}
		if head != "A" {
			t.Errorf("want A, got %v", head)
		}
	})
}

func TestGetTail(t *testing.T) {
	t.Run("should return an error when linked list has no nodes", func(t *testing.T) {
		sll := module1.NewSinglyLinkedList()
		if _, err := sll.Tail(); err == nil {
			t.Errorf("wants no error, got %v", err)
		}
	})
	t.Run("should return the tail of a non-empty linked list", func(t *testing.T) {
		sll := module1.NewSinglyLinkedList()
		sll.InsertBegin("A")
		sll.InsertEnd("B")
		tail, err := sll.Tail()
		if err != nil {
			t.Errorf("wants no error, got %v", err)
		}
		if tail != "B" {
			t.Errorf("want B, got %v", tail)
		}
	})
}

func TestInsertBegin(t *testing.T) {
	ll := module1.NewSinglyLinkedList()
	ll.InsertBegin("A")
	ll.InsertBegin("B")
	head, _ := ll.Head()
	if head != "B" {
		t.Errorf("want %v, got %v", "B", head)
	}
}

func TestInsertEnd(t *testing.T) {
	ll := module1.NewSinglyLinkedList()
	ll.InsertEnd("A")
	ll.InsertEnd("B")
	tail, _ := ll.Tail()
	if tail != "B" {
		t.Errorf("want %v, got %v", "B", tail)
	}
}
