package question

// 题目描述：给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

func longestValidParentheses(s string) int {
	return longestValidParenthesesCore3(s)
}

// 两遍扫描
func longestValidParenthesesCore1(s string) int {
	// 首先去掉左边的右括号和右边的左括号，因为他们不可能包括在答案里面
	start, end := 0, len(s)-1
	for ; start < len(s) && s[start] == ')'; start++ {
	}
	for ; end >= 0 && s[end] == '('; end-- {
	}
	s = s[start : end+1]
	if len(s) == 0 {
		return 0
	}
	// 此时保证s必然是以左括号开头，右括号结尾
	remain := 0 // 当前剩余的左括号
	left := -1
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			remain++
		} else if s[i] == ')' {
			remain--
			if remain == 0 {
				res = max(res, i-left)
			} else if remain < 0 {
				for i < len(s) && s[i] == ')' {
					i++
				}
				i--
				remain = 0
				left = i
			}
		}
	}
	remain = 0 // 当前剩余的右括号
	right := len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			remain++
		} else if s[i] == '(' {
			remain--
			if remain == 0 {
				res = max(res, right-i)
			} else if remain < 0 {
				for i >= 0 && s[i] == '(' {
					i--
				}
				i++
				remain = 0
				right = i

			}
		}
	}
	return res
}

// 动态规划
// dp[i]表示以下标i结尾的最长有效括号的长度
// if s[i] == ')' && s[i-1] == '(' dp[i]=dp[i-2]+2
// if s[i] == ')' && s[i-1] == ')' { if s[i-dp[i-1]-1] == '(' dp[i]=dp[i-1]+dp[i-dp[i-1]-2]+2}
// 第一种情况是()
// 第二种情况是yyyyy(xxxxxxx)这是因为排除掉中间部分的结果，右括号和左括号匹配起来了
func longestValidParenthesesCore2(s string) int {
	res:=0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			}else if i - dp[i-1]>0 && s[i-dp[i-1]-1]=='('{
				if i-dp[i-1]>=2{
					dp[i]=dp[i-1]+dp[i-dp[i-1]-2]+2
				}else{
					dp[i]=dp[i-1]+2
				}
			}
			res=max(res,dp[i])
		}
	}
	return res
}

// 栈: 始终保持栈顶元素为当前已经遍历过的元素中【最后一个没有被匹配的右括号下标】
func longestValidParenthesesCore3(s string) int {
	res:=0
	stk:=make([]int,0,len(s)+1)
	stk=append(stk,-1)
	for i:=0;i<len(s);i++{
		if s[i] == '(' {
			stk = append(stk, i)
		}else{
			stk=stk[:len(stk)-1]
			if len(stk) == 0{
				stk=append(stk,i)
			}else{
				res=max(res,i-stk[len(stk)-1])
			}
		}
	}
	return res
}