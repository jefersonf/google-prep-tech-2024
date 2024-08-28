package queue

type Person struct {
	position         int
	remainingTickets int
}

type Node struct {
	value Person
	next  *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) Enqueue(value Person) {
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

func (q *Queue) Dequeue() (Person, bool) {
	if q.size == 0 {
		return Person{}, false
	}
	value := q.head.value
	q.head = q.head.next
	q.size -= 1
	return value, true
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
