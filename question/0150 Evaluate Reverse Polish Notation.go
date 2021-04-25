package question

//逆波兰式求值
func evalRPN(tokens []string) int {
	// 遇到数字直接入栈
	// 遇到运算符则取出两个数字计算后再入栈即可
	s := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			op1 := s[len(s)-2]
			op2 := s[len(s)-1]
			s = s[:len(s)-2]
			s = append(s, op1+op2)
			continue
		case "-":
			op1 := s[len(s)-2]
			op2 := s[len(s)-1]
			s = s[:len(s)-2]
			s = append(s, op1-op2)
			continue
		case "*":
			op1 := s[len(s)-2]
			op2 := s[len(s)-1]
			s = s[:len(s)-2]
			s = append(s, op1*op2)
			continue
		case "/":
			op1 := s[len(s)-2]
			op2 := s[len(s)-1]
			s = s[:len(s)-2]
			s = append(s, op1/op2)
			continue
		default:
			//处理符号
			j := 0
			flag := false //默认正数
			if tokens[i][0] == '-' {
				flag = true
				j = 1
			}
			tmp := 0
			for ; j < len(tokens[i]); j++ {
				tmp = tmp*10 + int(tokens[i][j]-'0')
			}
			if flag {
				tmp = -tmp
			}
			s = append(s, tmp)
		}
	}
	return s[0]
}
