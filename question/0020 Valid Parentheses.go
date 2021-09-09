package question

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