package question

func exist(board [][]byte, word string) bool {
	if word == "" {
		return true
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	var dfs func(i, j int, cur int) bool
	use := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		use[i] = make([]bool, len(board[0]))
	}
	dfs = func(i, j int, cur int) bool {
		if board[i][j] == word[cur] { //使用当前位置
			use[i][j] = true
			if cur == len(word)-1 {
				return true
			}
			//上
			if i >= 1 && !use[i-1][j] && dfs(i-1, j, cur+1) {
				return true
			}
			//下
			if i < len(board)-1 && !use[i+1][j] && dfs(i+1, j, cur+1) {
				return true
			}
			//左
			if j >= 1 && !use[i][j-1] && dfs(i, j-1, cur+1) {
				return true
			}
			//右
			if j < len(board[0])-1 && !use[i][j+1] && dfs(i, j+1, cur+1) {
				return true
			}
			use[i][j] = false
		}
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}
