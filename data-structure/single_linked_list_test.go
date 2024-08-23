package datastructure_test

import (
	datastructure "preptech24/data-structure"
	"testing"
)

func TestInsertBegin(t *testing.T) {
	ll := datastructure.NewSingleLinkedList()
	ll.InsertBegin("A")
	ll.InsertBegin("B")

	if ll.Head.Data() != "B" {
		t.Errorf("want %v, got %v", "B", ll.Head.Data())
	}
}
