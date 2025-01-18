package heap

import "errors"

type HeapTree struct {
	Nodes []int
}

func NewHeapTree(arr []int) *HeapTree {
	h := &HeapTree{}
	for _, v := range arr {
		h.Add(v)
	}
	return h
}

func (t *HeapTree) swap(i, j int) {
	t.Nodes[i], t.Nodes[j] = t.Nodes[j], t.Nodes[i]
}

func (t *HeapTree) Add(x int) {
	t.Nodes = append(t.Nodes, x)
	index := len(t.Nodes) - 1
	parentIndex := (index - 1) / 2
	for parentIndex >= 0 && parentIndex != index {
		if t.Nodes[parentIndex] < t.Nodes[index] {
			t.swap(parentIndex, index)
		}
		index = parentIndex
		parentIndex = (index - 1) / 2
	}
}

func (t *HeapTree) Pop() (int, error) {
	size := len(t.Nodes)
	if size == 0 {
		return 0, errors.New("empty heap")
	}

	top := t.Nodes[0]

	t.swap(0, size-1)
	t.Nodes = t.Nodes[:size-1]

	t.heapify(0)

	return top, nil
}

func (t *HeapTree) heapify(currentIndex int) {
	size := len(t.Nodes)
	if size <= 1 {
		return
	}

	leftChildIndex := 2*currentIndex + 1
	if leftChildIndex >= size {
		return
	}

	rightChildIndex := 2*currentIndex + 2
	if rightChildIndex >= size {
		if t.Nodes[leftChildIndex] > t.Nodes[currentIndex] {
			t.swap(leftChildIndex, currentIndex)
		}
		return
	}

	if t.Nodes[leftChildIndex] > t.Nodes[rightChildIndex] {
		t.swap(leftChildIndex, currentIndex)
		t.heapify(leftChildIndex)
	} else {
		t.swap(rightChildIndex, currentIndex)
		t.heapify(rightChildIndex)
	}
}
