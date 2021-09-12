package question

// 题目描述：请你判断一个 9x9 的数独是否有效。
// 规则：
//	数字1-9在每一行只能出现一次。
//	数字1-9在每一列只能出现一次。
//	数字1-9在每一个以粗实线分隔的3x3宫内只能出现一次。
//	数独部分空格内已填入了数字，空白格用'.'表示。

func isValidSudoku(board [][]byte) bool {
	return isValidSudokuCore2(board)
}

// 检查是否符合规则即可
func isValidSudokuCore1(board [][]byte) bool {
	row := make(map[byte]struct{})
	col := make(map[byte]struct{})
	cross := make(map[byte]struct{})
	// 行
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if _, ok := row[board[i][j]]; board[i][j] == '.' || !ok {
				row[board[i][j]] = struct{}{}
			} else {
				return false
			}
		}
		row = make(map[byte]struct{})
	}
	// 列
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if _, ok := col[board[j][i]]; board[j][i] == '.' || !ok {
				col[board[j][i]] = struct{}{}
			} else {
				return false
			}
		}
		col = make(map[byte]struct{})
	}
	// 宫格
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if _, ok := cross[board[i*3+k][j*3+l]]; board[i*3+k][j*3+l] == '.' || !ok {
						cross[board[i*3+k][j*3+l]] = struct{}{}
					} else {
						return false
					}
				}
			}
			cross = make(map[byte]struct{})
		}
	}
	return true
}

// 用位图节省空间，并缩减为一次遍历
func isValidSudokuCore2(board [][]byte) bool {
	var row, col, block [9]uint16
	var cur uint16
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			cur = 1 << (board[i][j] - '1')                      // 当前数字的 二进制数位 位置
			bi := i/3 + j/3*3                                   // 3x3的块索引号
			if (row[i]&cur)|(col[j]&cur)|(block[bi]&cur) != 0 { // 使用与运算查重
				return false
			}
			// 在对应的位图上，加上当前数字
			row[i] |= cur
			col[j] |= cur
			block[bi] |= cur
		}
	}
	return true
}
