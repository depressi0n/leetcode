package question

// 给定一个m x n 二维字符网格board 和一个字符串单词word 。如果word 存在于网格中，返回 true ；否则，返回 false 。
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

func exist(board [][]byte, word string) bool {
	return existCore(board,word)
}

// 找到起始位置，然后使用深度优先遍历，如果完成则直接返回，失败则找下一个位置
// 剪枝策略是：如果四周都没有合适的则直接返回而不是往下走
func existCore(board [][]byte, word string) bool {
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
