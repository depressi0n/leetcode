package question

func LetterCasePermutation(S string) []string {
	if len(S) == 0 {
		return nil
	}
	var res []string
	var dfs func(cur int, ts string)
	dfs = func(cur int, ts string) {
		if cur == len(S) {
			res = append(res, ts)
			return
		}
		if 'a' <= S[cur] && S[cur] <= 'z' {
			ts += string(S[cur] + 'A' - 'a')
			dfs(cur+1, ts)
			dfs(cur+1, ts[:len(ts)-1]+string(S[cur]))
		} else if 'A' <= S[cur] && S[cur] <= 'Z' {
			ts += string(S[cur] - 'A' + 'a')
			dfs(cur+1, ts)
			dfs(cur+1, ts[:len(ts)-1]+string(S[cur]))
		} else {
			dfs(cur+1, ts+string(S[cur]))
		}
	}
	dfs(0, "")
	return res
}
