package question

func findWords(board [][]byte, words []string) []string {
	res := make([]string, 0)
	// 定义遍历函数
	trace := func(i, j, k int) bool {
		flag := make([][]bool, len(board))
		for t := 0; t < len(board); t++ {
			flag[t] = make([]bool, len(board[0]))
		}
		var dfs func(i, j, k, s int) bool
		dfs = func(i, j, k, s int) bool {
			if s+1 >= len(words[k]) {
				return true
			}
			flag[i][j] = true
			if i+1 < len(board) && !flag[i+1][j] && board[i+1][j] == words[k][s+1] {
				if dfs(i+1, j, k, s+1) {
					return true
				}
			}
			if i-1 >= 0 && !flag[i-1][j] && board[i-1][j] == words[k][s+1] {
				if dfs(i-1, j, k, s+1) {
					return true
				}
			}
			if j+1 < len(board[0]) && !flag[i][j+1] && board[i][j+1] == words[k][s+1] {
				if dfs(i, j+1, k, s+1) {
					return true
				}
			}
			if j-1 >= 0 && !flag[i][j-1] && board[i][j-1] == words[k][s+1] {
				if dfs(i, j-1, k, s+1) {
					return true
				}
			}
			flag[i][j] = false
			return false
		}
		return dfs(i, j, k, 0)
	}
	for i := 0; i < len(words); i++ {
		// 首先找到首字母
		for c := 0; c < len(board); c++ {
			tf := false
			for l := 0; l < len(board[0]); l++ {
				if board[c][l] == words[i][0] {
					// 然后从首字母开始遍历查询
					if trace(c, l, i) {
						// 如果找到则放入结果中
						res = append(res, words[i])
						tf = true
						break
					}
				}
			}
			if tf {
				break
			}
		}
		// 否则跳到第一步，如果已经全部试过了则说明不可能找到
	}
	return res
}
