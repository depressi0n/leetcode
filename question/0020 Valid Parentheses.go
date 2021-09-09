package question

//题目描述：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//有效字符串需满足：
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。


func isValid(s string) bool {
	return isValidCore(s)
}

func isValidCore(s string)bool{
	stack:=make([]byte,0,len(s))
	for i:=0;i<len(s);i++{
		switch s[i] {
		case '(','[','{':
			stack=append(stack,s[i])
		case ')':
			if len(stack)==0 || stack[len(stack)-1]!='('{
				return false
			}else{
				stack=stack[:len(stack)-1]
			}
		case ']':
			if len(stack)==0 ||stack[len(stack)-1]!='['{
				return false
			}else{
				stack=stack[:len(stack)-1]
			}
		case '}':
			if len(stack)==0 ||stack[len(stack)-1]!='{'{
				return false
			}else{
				stack=stack[:len(stack)-1]
			}
		}
	}
	if len(stack)!=0{
		return false
	}
	return true
}