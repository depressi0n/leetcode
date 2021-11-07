package question

// 给定一个初始元素全部为0，大小为 m*n 的矩阵M以及在M上的一系列更新操作。
// 操作用二维数组表示，其中的每个操作用一个含有两个正整数a 和 b 的数组表示，含义是将所有符合0 <= i < a 以及 0 <= j < b 的元素M[i][j]的值都增加 1。
// 在执行给定的一系列操作后，你需要返回矩阵中含有最大整数的元素个数。

func maxCount(m int, n int, ops [][]int) int {
	return maxCountCore(m, n, ops)
}

// 朴素思想：对每一次操作都进行作用到矩阵上，并记录最大值，最后遍历矩阵求得元素个数
func maxCountCore(m int, n int, ops [][]int) int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	res, cnt := 0, 0
	for i := 0; i < len(ops); i++ {
		for j := 0; j < ops[i][0]; j++ {
			for k := 0; k < ops[i][1]; k++ {
				matrix[j][k]++
				if matrix[j][k] > res {
					res = matrix[j][k]
					cnt = 1
				} else if matrix[j][k] == res {
					cnt++
				}
			}
		}
	}
	return cnt
}

// 优化方向：每次都从左上角开始，所有左上角必然是一个最大值，那么记录每次操作的交集，交集内元素个数即答案
func maxCountCore2(m int, n int, ops [][]int) int {
	row, col := m, n
	for i := 0; i < len(ops); i++ {
		row = min(row, ops[i][0])
		col = min(col, ops[i][1])
	}
	return row * col
}
