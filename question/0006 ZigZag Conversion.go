package question

// 题目描述：将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

func convert(s string, numRows int) string {
	return convertCore(s, numRows)
}


// 根据行数确定下一个位置，如果触及到了顶部，则向下，如果触及底部，则向左上
func convertCore(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	matrix := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		matrix[i] = make([]byte, 0, 1000)
	}
	cur := 0
	x, y := 0, 0
	dir := false // false 表示往下，true表示左上
	for cur < len(s) {
		if len(matrix[x]) < y-1 {
			matrix[x] = append(matrix[x], ' ')
		}
		matrix[x] = append(matrix[x], s[cur])
		if x == numRows-1 {
			dir = true
		} else if x == 0 {
			dir = false
		}
		if !dir {
			x++
		} else {
			x--
			y++
		}
		cur++
	}
	res := make([]byte, 0, len(s))
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != ' ' {
				res = append(res, matrix[i][j])
			}
		}
	}
	return string(res)
}
