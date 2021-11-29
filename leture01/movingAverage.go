package leture01

import "container/list"

type MovingAverage struct {
	nums     *list.List
	capacity int
	sum      int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		list.New(),
		size,
		0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.nums.Len() < this.capacity {
		this.sum += val
		this.nums.PushBack(val)
		return float64(this.sum) / float64(this.nums.Len())
	}
	this.sum = this.sum - this.nums.Front().Value.(int) + val
	this.nums.Remove(this.nums.Front())
	this.nums.PushBack(val)
	return float64(this.sum) / float64(this.capacity)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
