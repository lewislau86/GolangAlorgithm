package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	head1 := l1
	head2 := l2
	n1, n2, carry, current := 0, 0, 0, head
	for head1 != nil || head2 != nil || carry != 0 {
		if head1 == nil {
			n1 = 0
		} else {
			n1 = head1.Val
			head1 = head1.Next
		}

		if head2 == nil {
			n2 = 0
		} else {
			n2 = head2.Val
			head2 = head2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10

	}

	return head.Next
}

func main() {
	// node1
	node10 := new(ListNode)
	node10.Val = 2
	node11 := new(ListNode)
	node11.Val = 4
	node12 := new(ListNode)
	node12.Val = 3
	node10.Next = node11
	node11.Next = node12
	node12.Next = nil
	// node2
	node20 := new(ListNode)
	node20.Val = 5
	node21 := new(ListNode)
	node21.Val = 6
	node22 := new(ListNode)
	node22.Val = 4
	node20.Next = node21
	node21.Next = node22
	node22.Next = nil
	addTwoNumbers(node10, node20)
}
