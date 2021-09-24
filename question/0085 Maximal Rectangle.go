package question

// 给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
func maximalRectangle(matrix [][]byte) int {
	return maximalRectangleCore2(matrix)
}

// 暴力思路：枚举所有的左上角和右下角坐标，检查矩形是否合法 =>复杂度爆炸
// 优化思想：计算出每一元素左边连续1的数量
// left[i][j]表示第i行第j列的左边连续1的数量
// 此时，枚举右下角元素matrix[i][j]，向上枚举所有的同列上元素，left[i][j],left[i-1][j],left[i-2][j]...,left[k][j]中的最小值即为矩形
func maximalRectangleCore1(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	left := make([][]int, len(matrix))
	for i, row := range matrix {
		left[i] = make([]int, len(matrix[i]))
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}
	res := 0
	// 枚举每一个可能的右下角元素
	for i, row := range matrix {
		for j, v := range row {
			// 不适合作为右下元素
			if v == '0' {
				continue
			}
			// 寻找高度
			width := left[i][j]
			area := width
			for k := i - 1; k >= 0; k-- {
				width = min(width, left[k][j])
				area = max(area, (i-k+1)*width)
			}
			res = max(res, area)
		}
	}
	return res
}

// 复用上一题的思路，依次增加一行
func maximalRectangleCore2(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	res := 0
	cnt := 0
	cur := make([]int, len(matrix[cnt]))
	for cnt < len(matrix) {
		for j := 0; j < len(matrix[cnt]); j++ {
			if matrix[cnt][j] == '1' {
				cur[j] += 1
			} else {
				cur[j] = 0
			}
		}
		left, right := make([]int, len(cur)), make([]int, len(cur))
		for i := 0; i < len(cur); i++ {
			right[i] = len(cur)
		}
		s := make([]int, 0, len(cur))
		for k := 0; k < len(cur); k++ {
			for len(s) > 0 && cur[s[len(s)-1]] >= cur[k] {
				right[s[len(s)-1]] = k
				s = s[:len(s)-1]
			}
			if len(s) == 0 {
				left[k] = -1
			} else {
				left[k] = s[len(s)-1]
			}
			s = append(s, k)
		}
		for i := 0; i < len(cur); i++ {
			res = max(res, (right[i]-left[i]-1)*cur[i])
		}
		cnt++
	}
	return res
}
