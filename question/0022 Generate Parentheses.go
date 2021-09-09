package question

import "strings"

// 题目描述：数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//有效括号组合需满足：左括号必须以正确的顺序闭合。

func generateParenthesis(n int) []string {
	return generateParenthesisCore(n)
}

func generateParenthesisCore(n int) []string{
	res:=make([]string,0,100)
	t := strings.Builder{}
	var dfs func(cur int,remain int,cache []byte)
	dfs = func(cur int,remain int,cache []byte) {
		if cur == 0 && remain == 0{
			t.Reset()
			t.Write(cache)
			res=append(res,t.String())
			return
		}
		if cur!=0{
			cache=append(cache,')')
			dfs(cur-1,remain,cache)
			cache=cache[:len(cache)-1]
		}
		if remain!=0{
			cache=append(cache,'(')
			dfs(cur+1,remain-1,cache)
			cache=cache[:len(cache)-1]
		}

	}
	dfs(0,n,[]byte{})
	return res
}
