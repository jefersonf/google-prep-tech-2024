package module1

// SinglyLinkedList defines all possible operations.
type SinglyLinkedList interface {
	InsertBegin(data any)
	InsertEnd(data any)
	Insert(zeroIndexedPosition int, data any) error

	RemoveHead() error
	RemoveTail() error
	Remove(zeroIndexedPosition int) error

	Search(index int) (any, error)
	FindAll(data any) []any
	Head() (any, error)
	Tail() (any, error)

	Reverse()
}

type listNode struct {
	Data any
	Next *listNode
}

type singlyLinkedList struct {
	head *listNode
}

func newNode(data any) *listNode {
	return &listNode{data, nil}
}

func NewSinglyLinkedList() SinglyLinkedList {
	return &singlyLinkedList{}
}
