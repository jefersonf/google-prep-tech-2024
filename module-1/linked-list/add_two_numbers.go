package module1

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	currentNode := &ListNode{}
	headNode := currentNode

	var carryOut int

	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carryOut
		currentNode.Val = sum % 10
		carryOut = sum / 10
		if l1.Next != nil || l2.Next != nil {
			currentNode.Next = &ListNode{}
			currentNode = currentNode.Next
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum := l1.Val + carryOut
		currentNode.Val = sum % 10
		carryOut = sum / 10
		if l1.Next != nil {
			currentNode.Next = &ListNode{}
			currentNode = currentNode.Next
		}
		l1 = l1.Next
	}

	for l2 != nil {
		sum := l2.Val + carryOut
		currentNode.Val = sum % 10
		carryOut = sum / 10
		if l2.Next != nil {
			currentNode.Next = &ListNode{}
			currentNode = currentNode.Next
		}
		l2 = l2.Next
	}

	if carryOut > 0 {
		currentNode.Next = &ListNode{Val: carryOut}
	}

	return headNode
}
