package main

import (
	"math/big"
)

func main() {
}

const (
	MaxBit = 1001
)

type SmallestInfiniteSet struct {
	num *big.Int
}

func Constructor() SmallestInfiniteSet {
	num := new(big.Int)
	for i := 1; i < MaxBit; i++ {
		num.SetBit(num, i, 1)
	}
	return SmallestInfiniteSet{
		num: num,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	// 查找第一个值为 1 的位置
	for i := 0; i < MaxBit; i++ {
		if this.num.Bit(i) == 1 {
			this.num.SetBit(this.num, i, 0)
			return i
		}
	}
	return 0
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num > MaxBit {
		return
	}
	this.num.SetBit(this.num, num, 1)
}
