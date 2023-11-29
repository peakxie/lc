package main

import (
	"container/list"
)

func main() {
	// fmt.Println(minD(grid, moveCost))
}

type FrontMiddleBackQueue struct {
	part1 *list.List
	part2 *list.List
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		part1: list.New(),
		part2: list.New(),
	}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.part1.PushFront(val)
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	// 后面比前面大2
	if this.part2.Len() > this.part1.Len()+1 {
		l := this.part2.Len() - this.part1.Len()
		for i := 0; i < l/2; i++ {
			e := this.part2.Front()
			if e != nil {
				this.part2.Remove(e)
				this.part1.PushBack(e.Value)
			}
		}
	}

	// 前面比后面大1
	if this.part1.Len() > this.part2.Len() && this.part2.Len() > 0 {
		l := this.part1.Len() - this.part2.Len()
		for i := 0; i < (l+1)/2; i++ {
			e := this.part1.Back()
			if e != nil {
				this.part1.Remove(e)
				this.part2.PushFront(e.Value)
			}
		}
	}

	if this.part1.Len() == this.part2.Len() {
		this.part2.PushFront(val)
		return
	}
	if this.part2.Len() == this.part1.Len()+1 {
		this.part1.PushBack(val)
		return
	}
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.part2.PushBack(val)
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.part1.Len() > 0 {
		e := this.part1.Front()
		this.part1.Remove(e)
		return e.Value.(int)
	}

	if this.part2.Len() > 0 {
		e := this.part2.Front()
		this.part2.Remove(e)
		return e.Value.(int)
	}

	return -1
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	// 后面比前面大2
	if this.part2.Len() > this.part1.Len()+1 {
		l := this.part2.Len() - this.part1.Len()
		for i := 0; i < l/2; i++ {
			e := this.part2.Front()
			if e != nil {
				this.part2.Remove(e)
				this.part1.PushBack(e.Value)
			}
		}
	}

	// 前面比后面大1
	if this.part1.Len() > this.part2.Len() && this.part1.Len() > 0 {
		l := this.part1.Len() - this.part2.Len()
		for i := 0; i < (l+1)/2; i++ {
			e := this.part1.Back()
			if e != nil {
				this.part1.Remove(e)
				this.part2.PushFront(e.Value)
			}
		}
	}

	if this.part1.Len() == this.part2.Len() && this.part1.Len() > 0 {
		e := this.part1.Back()
		this.part1.Remove(e)
		return e.Value.(int)
	}
	if this.part2.Len() == this.part1.Len()+1 {
		e := this.part2.Front()
		this.part2.Remove(e)
		return e.Value.(int)
	}

	return -1
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.part2.Len() > 0 {
		e := this.part2.Back()
		this.part2.Remove(e)
		return e.Value.(int)
	}

	if this.part1.Len() > 0 {
		e := this.part1.Back()
		this.part1.Remove(e)
		return e.Value.(int)
	}

	return -1
}
