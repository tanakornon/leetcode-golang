package adt

type LinkedListStack struct {
	head *Node
	size int
}

type Node struct {
	val  int
	next *Node
}

func NewLinkedListStack() LinkedListStack {
	return LinkedListStack{
		head: nil,
		size: 0,
	}
}

func (this *LinkedListStack) Push(val int) {
	this.head = &Node{val: val, next: this.head}
	this.size++
}

func (this *LinkedListStack) Pop() int {
	val := this.Peek()
	this.head = this.head.next
	this.size--
	return val
}

func (this *LinkedListStack) Peek() int {
	return this.head.val
}
