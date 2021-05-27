package question

func ReverseParentheses(s string) string {
	// TODO 没理解这个思想
	//	每个括号内的字母都要翻转
	// 最简单的方法是遇到括号翻转，预期构造递归调用形式，用栈替换递归
	//	从前往后和从后往前
	//stack := [][]byte{}
	//str := []byte{}
	//for i := range s {
	//	if s[i] == '(' {
	//		stack = append(stack, str) //遇到左括号就添加到那一层
	//		str = []byte{}
	//	} else if s[i] == ')' {
	//		for j, n := 0, len(str); j < n/2; j++ {
	//			str[j], str[n-1-j] = str[n-1-j], str[j]
	//		} // 遇到右括号则将str的内容反转
	//		str = append(stack[len(stack)-1], str...) // 添加到后面，拼接起来
	//		stack = stack[:len(stack)-1]
	//	} else {
	//		str = append(str, s[i])
	//	}
	//}
	//return string(str)

	// 逆序遍历括号
	// 先向右移动匹配左括号，找到对应的右括号
	// 在括号内部向左移动
	// 向左移动到左括号，则跳到对应的右括号
	// 在括号外向右移动
	//n := len(s)
	//pair := make([]int, n)
	//stack := []int{}
	//for i, b := range s {
	//	if b == '(' {
	//		stack = append(stack, i)
	//	} else if b == ')' {
	//		j := stack[len(stack)-1]
	//		stack = stack[:len(stack)-1]
	//		pair[i], pair[j] = j, i // 括号配对
	//	}
	//}
	//
	//ans := []byte{}
	//for i, step := 0, 1; i < n; i += step {
	//	if s[i] == '(' || s[i] == ')' {
	//		i = pair[i] //跳到右括号
	//		step = -step
	//	} else {
	//		ans = append(ans, s[i])
	//	}
	//}
	//return string(ans)

	brackets := make([]int, len(s))
	stack := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			brackets[stack[len(stack)-1]] = i
			brackets[i] = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	res := []byte{}
	// 现在对lefts和rights进行处理
	i := 0
	dir := 1
	for i < len(s) {
		if s[i] == '(' {
			i = brackets[i] // 转向右括号并转向，再次回到左括号时，会跳到右括号继续往后
			dir = -dir
		} else if s[i] == ')' { // 转向左括号并转向，再次回到右括号时，会跳到左括号继续往前
			i = brackets[i]
			dir = -dir
		} else {
			res = append(res, s[i])
		}
		i += dir
	}
	return string(res)
}
