package queue

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) Enqueue(value int) {
	newNode := &Node{value: value}
	if q.size == 0 {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.size += 1
}

func (q *Queue) Dequeue() (int, bool) {
	if q.size == 0 {
		return 0, false
	}
	value := q.head.value
	q.head = q.head.next
	q.size -= 1
	return value, true
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
