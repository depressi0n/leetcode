package question

// 给定一个三角形 triangle ，找出自顶向下的最小路径和。
//每一步只能移动到下一行中相邻的结点上。
//相邻的结点在这里指的是
//（1）下标 与 上一层结点下标 相同
//（2）或者等于 上一层结点下标 + 1 的两个结点
// 也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。

func minimumTotal(triangle [][]int) int {
	return minimumTotalCore(triangle)
}

// 典型的动态规划问题，从下往上
func minimumTotalCore(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	res := make([]int, len(triangle))
	// 初始化
	for i := 0; i < len(res); i++ {
		res[i] = triangle[len(triangle)-1][i]
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if res[j] < res[j+1] {
				res[j] = res[j] + triangle[i][j]
			} else {
				res[j] = res[j+1] + triangle[i][j]
			}
		}
	}
	return res[0]
}

// 直接是triangle作为辅助空间
func minimumTotalCore2(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if triangle[i+1][j] < triangle[i+1][j+1] {
				triangle[i][j] += triangle[i+1][j]
			} else {
				triangle[i][j] += triangle[i+1][j+1]
			}
		}
	}
	return triangle[0][0]
}
