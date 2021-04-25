package question

type MinStack struct {
	s    []int
	minS []int
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	return MinStack{
		s:    make([]int, 0),
		minS: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.s = append(this.s, x)
	if len(this.minS) == 0 {
		this.minS = append(this.minS, x)
		return
	}
	if x < this.minS[len(this.minS)-1] {
		this.minS = append(this.minS, x)
	} else {
		this.minS = append(this.minS, this.minS[len(this.minS)-1])
	}
	return
}

func (this *MinStack) Pop() {
	this.s = this.s[:len(this.s)-1]
	this.minS = this.minS[:len(this.minS)-1]
	return
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	return this.minS[len(this.minS)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
