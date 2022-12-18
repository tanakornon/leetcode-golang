package problems

import (
	. "leetcode-go/adt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var curr *ListNode

	length := 0

	for curr = head; curr != nil; curr = curr.Next {
		length++
	}

	if length == n {
		return head.Next
	}

	curr = head
	for i := 0; i < length-n-1; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next

	return head
}
