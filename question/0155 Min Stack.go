package question
// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
//push(x) —— 将元素 x 推入栈中。
//pop()—— 删除栈顶的元素。
//top()—— 获取栈顶元素。
//getMin() —— 检索栈中的最小元素。

// 主要思想：使用双栈来实现最小栈，一个栈存放原始元素，一个栈存放栈中最小元素
// 满足minS[i]<=S[i]

type MinStack struct {
	s    []int
	minS []int
}

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
