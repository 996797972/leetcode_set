// 删除单链表的倒数第N个结点
package linkops

import (
	. "leetcode_set/algorithm"
)

func RemoveNodeFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head} // 在原头结点前加入一个哑结点(dummy node)
	fast, slow := head, dummy

	for ; n > 0; n-- {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next // 由于slow是从哑结点开始的，循环结束后刚好停在待删结点的前继结点
	return dummy.Next
}