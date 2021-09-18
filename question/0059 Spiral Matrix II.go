package question

// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

func generateMatrix(n int) [][]int {
	return generateMatrixCore(n)
}

func generateMatrixCore(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	i, j, start, end := 0, 0, 0, n
	cur := 1
	for start+1 != end || start+2 != end {
		for j < end {
			res[i][j] = cur
			cur++
			j++
		}
		j--
		i++
		for i < end {
			res[i][j] = cur
			cur++
			i++
		}
		i--
		j--
		for j >= start {
			res[i][j] = cur
			cur++
			j--
		}
		j++
		i--
		for i > start {
			res[i][j] = cur
			cur++
			i--
		}
		start++
		end--
		if start >= end {
			return res
		}
		i = start
		j = start
	}
	return res
}
