package adt

type Stack struct {
	items []int
	n     int
}

func NewStack() Stack {
	return Stack{
		items: []int{},
		n:     0,
	}
}

func (this *Stack) Push(item int) {
	this.items = append(this.items, item)
	this.n++
}

func (this *Stack) Pop() int {
	top := this.Peek()
	this.n--
	this.items = this.items[:this.n]
	return top
}

func (this *Stack) Peek() int {
	return this.items[this.n-1]
}
