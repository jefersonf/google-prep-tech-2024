package module1_test

import (
	module1 "preptech24/module-1/linked-list"
	"testing"
)

func TestInsertBegin(t *testing.T) {
	ll := module1.NewSinglyLinkedList()
	ll.InsertBegin("A")
	ll.InsertBegin("B")
	head, _ := ll.Head()
	if head != "B" {
		t.Errorf("want %v, got %v", "B", head)
	}
}
