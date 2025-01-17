package hashtable

import "errors"

type node struct {
	Key, Value int
	Next       *node
}

type hashTable struct {
	array []*node
}

func NewHashTable(size int) *hashTable {
	return &hashTable{
		array: make([]*node, size),
	}
}

func (ht *hashTable) hash(key int) int {
	return key % len(ht.array)
}

func (ht *hashTable) Insert(key, value int) {
	index := ht.hash(key)
	curr := ht.array[index]
	newNode := &node{Key: key, Value: value}
	if curr == nil {
		ht.array[index] = newNode
		return
	}
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

func (ht *hashTable) Find(key int) ([]int, error) {
	index := ht.hash(key)
	curr := ht.array[index]
	values := make([]int, 0)
	for curr != nil {
		values = append(values, curr.Value)
		curr = curr.Next
	}

	if len(values) == 0 {
		return nil, errors.New("key not found")
	}
	return values, nil
}

func (ht *hashTable) Remove(key, value int) {
	index := ht.hash(key)
	curr := ht.array[index]
	if curr == nil {
		return
	}
	for curr != nil {
		if curr.Value == value {
			if curr.Next == nil {
				curr = nil
				return
			}
			curr.Value = curr.Next.Value
			curr.Next = curr.Next.Next
			return
		}
		curr = curr.Next
	}

}

func hash2(r rune) int {
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}
