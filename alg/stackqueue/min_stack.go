package stackqueue

import "math"

type MinStack struct {
	elem []int
	minS []int
}

func Constructor() MinStack {
	return MinStack{
		elem: make([]int, 0),
		minS: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.elem = append(this.elem, val)
	curMin := math.MaxInt32
	if len(this.minS) > 0 {
		curMin = min(curMin, this.minS[len(this.minS)-1])
	}
	this.minS = append(this.minS, min(val, curMin))
}

func (this *MinStack) Pop() {
	if len(this.elem) > 0 {
		this.elem = this.elem[:len(this.elem)-1]
	}
	if len(this.minS) > 0 {
		this.minS = this.minS[:len(this.minS)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.elem) > 0 {
		return this.elem[len(this.elem)-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if len(this.minS) > 0 {
		return this.minS[len(this.minS)-1]
	}
	return -1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
