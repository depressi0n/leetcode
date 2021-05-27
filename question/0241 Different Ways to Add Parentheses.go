package question

func DiffWaysToCompute(expression string) []int {
	// 返回表达式所有可能的结果
	// 题意是通过对不同运算符优先级不做设定，可能的结果有哪些
	// 并且对结果不需要按照特定顺序
	// 运算符只包括+、-、*
	// 思路：对任意一个运算符进行划分，递归处理左右两边，得到的结果就是所求的所有结果
	// 递归的返回条件是没有运算符的时候
	operators := make([]int, 0, len(expression))
	for i := 0; i < len(expression); i++ {
		switch expression[i] {
		case '+', '-', '*':
			operators = append(operators, i)
		default:
		}
	}
	if len(operators) == 0 {
		// 将expression转行为数值
		t := 0
		for i := 0; i < len(expression); i++ {
			t = t*10 + int(expression[i]-'0')
		}
		return []int{t}
	}
	//m := make(map[int]struct{})
	res := make([]int, 0, len(expression))
	for i := 0; i < len(operators); i++ {
		lefts := DiffWaysToCompute(expression[:operators[i]])
		rights := DiffWaysToCompute(expression[operators[i]+1:])
		for j := 0; j < len(lefts); j++ {
			for k := 0; k < len(rights); k++ {
				var tmp int
				switch expression[operators[i]] {
				case '+':
					tmp = lefts[j] + rights[k]
				case '-':
					tmp = lefts[j] - rights[k]
				case '*':
					tmp = lefts[j] * rights[k]
				}
				res = append(res, tmp)
				//if _, ok := m[tmp]; !ok {
				//	res = append(res, tmp)
				//	m[tmp] = struct{}{}
				//}
			}
		}
	}
	return res
}
