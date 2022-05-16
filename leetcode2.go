package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 这个使用一个虚拟的头指针来遍历链表，这样就不用考虑链表为空的情况了
 * head1 -> 2 -> 4 -> 3
 * head2 -> 5 -> 6 -> 4
 * head  -> 7 -> 0 -> 8
 * carry 表示是否进位
 */
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
		// 这里有点没看懂
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10

	}

	return head.Next
}

func main() {
	// node1
	node12 := &ListNode{Val: 3, Next: nil}
	node11 := &ListNode{Val: 4, Next: node12}
	node10 := &ListNode{Val: 2, Next: node11}
	// node2
	node22 := &ListNode{Val: 4, Next: nil}
	node21 := &ListNode{Val: 6, Next: node22}
	node20 := &ListNode{Val: 5, Next: node21}
	result := addTwoNumbers(node10, node20)
	for result != nil {
		fmt.Print(result.Val)
		result = result.Next
	}

}
