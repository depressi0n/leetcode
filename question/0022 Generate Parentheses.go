package question

import "strings"

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
