package top150

import "math"

type MinStack struct {
	s  []int
	ms []int
}

func Constructor() MinStack {
	return MinStack{
		s:  []int{},
		ms: []int{math.MaxInt},
	}
}

func (this *MinStack) Push(val int) {
	this.s = append(this.s, val)
	m := this.ms[len(this.ms)-1]
	if val < m {
		m = val
	}
	this.ms = append(this.ms, m)
}

func (this *MinStack) Pop() {
	this.s = this.s[:len(this.s)-1]
	this.ms = this.ms[:len(this.ms)-1]
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	return this.ms[len(this.ms)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
