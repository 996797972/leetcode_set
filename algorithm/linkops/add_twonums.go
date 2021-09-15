package linkops

import . "leetcode_set/algorithm"

func AddTwoNums(l1, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		bitSum := n1 + n2 + carry
		if bitSum > 9 {
			carry = 1
			bitSum -= 10
		} else {
			carry = 0
		}
		if head == nil {
			head = &ListNode{bitSum, nil}
			tail = head
		} else {
			tail.Next = &ListNode{bitSum, nil}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{carry, nil}
	}
	return head
}