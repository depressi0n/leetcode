package question

type MyQueue struct {
	mainStack []int
	helpStack []int
}

/** Initialize your data structure here. */
func Constructor0232() MyQueue {
	return MyQueue{
		mainStack: make([]int, 0, 100),
		helpStack: make([]int, 0, 100),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	// 辅助栈中没有值，则直接插入到主栈中
	if len(this.helpStack) == 0 {
		this.mainStack = append(this.mainStack, x)
		return
	}
	// 否则将辅助栈的值放回主栈，后再插入x到主栈
	for i := len(this.helpStack) - 1; i >= 0; i-- {
		this.mainStack = append(this.mainStack, this.helpStack[i])
	}
	this.helpStack = this.helpStack[:0] // 清空
	this.mainStack = append(this.mainStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.mainStack) == 0 {
		res := this.helpStack[len(this.helpStack)-1]
		this.helpStack = this.helpStack[:len(this.helpStack)-1]
		return res
	}
	for i := len(this.mainStack) - 1; i >= 1; i-- {
		this.helpStack = append(this.helpStack, this.mainStack[i])
	}
	res := this.mainStack[0]
	this.mainStack = this.mainStack[:0]
	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.mainStack) == 0 {
		return this.helpStack[len(this.helpStack)-1]
	}
	for i := len(this.mainStack) - 1; i >= 1; i-- {
		this.helpStack = append(this.helpStack, this.mainStack[i])
	}
	this.mainStack = this.mainStack[:1]
	return this.mainStack[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.mainStack) == 0 && len(this.helpStack) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
