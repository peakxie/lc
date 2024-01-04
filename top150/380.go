package top150

import "math/rand"

type RandomizedSet struct {
	nums   []int
	nummap map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nummap: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.nummap[val]; ok {
		return false
	}
	n := len(this.nums)
	this.nums = append(this.nums, val)
	this.nummap[val] = n
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	v, ok := this.nummap[val]
	if !ok {
		return false
	}
	n := len(this.nums)
	// 注意这个地方第一次做的时候，没有写
	this.nummap[this.nums[n-1]] = v
	this.nums[v] = this.nums[n-1]
	this.nums = this.nums[:n-1]
	delete(this.nummap, val)
	return true

}

func (this *RandomizedSet) GetRandom() int {
	if len(this.nums) == 0 {
		return -1
	}
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
