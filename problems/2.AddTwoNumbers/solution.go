package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	tail := head
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		}

		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		}

		sum := n1 + n2 + carry
		carry = sum / 10
		tail.Next = &ListNode{Val: sum % 10, Next: nil}
		tail = tail.Next
	}

	return head.Next
}
