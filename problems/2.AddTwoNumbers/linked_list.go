package problems

type LinkedList struct {
	head *ListNode
	len  int
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *LinkedList) createNode(value int) *ListNode {
	return &ListNode{
		Val:  value,
		Next: nil,
	}
}

func (list *LinkedList) createList(nums []int) {
	head := new(ListNode)
	tail := head

	for _, num := range nums {
		tail.Next = list.createNode(num)
		tail = tail.Next
	}

	list.head = head.Next
	list.len = len(nums)
}

func (list *LinkedList) toArray(head *ListNode) []int {
	var nums []int

	for ; head != nil; head = head.Next {
		nums = append(nums, head.Val)
	}

	return nums
}
