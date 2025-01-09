package stack

import "errors"

type pair struct {
	Value, MinimumValueFromBottom int
}

type minStack []pair

func NewMinStack() minStack {
	return make(minStack, 0)
}

func (ms *minStack) Add(elem int) {
	if len(*ms) == 0 {
		*ms = append(*ms, pair{elem, elem})
	} else {
		*ms = append(*ms, pair{elem, min(elem, (*ms)[len(*ms)-1].MinimumValueFromBottom)})
	}
}

func (ms minStack) GetMinimum() (int, error) {
	if len(ms) == 0 {
		return 0, errors.New("stask is empty")
	}
	return ms[len(ms)-1].MinimumValueFromBottom, nil
}
