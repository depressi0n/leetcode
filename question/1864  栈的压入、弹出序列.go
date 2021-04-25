package question

func validateStackSequences(pushed []int, popped []int) bool {
	// 模拟一个栈，
	if len(pushed) != len(popped) {
		return false
	}
	if pushed == nil || len(pushed) == 0 {
		return true
	}
	stack := make([]int, 0, len(pushed))
	stack = append(stack, pushed[0])
	nextPushed := 1
	nextPoped := 0
	for {
		//若栈为空或当前栈顶不是弹出序列下一个要弹出的值，则压下一个，否则弹出
		if len(stack) == 0 || stack[len(stack)-1] != popped[nextPoped] {
			if nextPushed < len(pushed) {
				stack = append(stack, pushed[nextPushed])
				nextPushed++
			} else {
				// 没有能继续压入的值，而此时肯定还有需要弹出的值
				return false
			}
		} else {
			// 弹出，同时移动到下一个要弹出的值
			stack = stack[:len(stack)-1]
			nextPoped++
			if nextPoped == len(popped) {
				return true
			}
		}
	}
}
