package adt

type LinkedList struct {
	Head *ListNode
	Len  int
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkedList(nums []int) *ListNode {
	list := new(LinkedList)
	list.CreateList(nums)
	return list.Head
}

func (list *LinkedList) CreateNode(value int) *ListNode {
	return &ListNode{
		Val:  value,
		Next: nil,
	}
}

func (list *LinkedList) CreateList(nums []int) {
	head := new(ListNode)
	tail := head

	for _, num := range nums {
		tail.Next = list.CreateNode(num)
		tail = tail.Next
	}

	list.Head = head.Next
	list.Len = len(nums)
}

func (list *LinkedList) ToArray(head *ListNode) []int {
	var nums []int

	for ; head != nil; head = head.Next {
		nums = append(nums, head.Val)
	}

	return nums
}
