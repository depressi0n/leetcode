package question

func maximalRectangle(matrix [][]byte) int {
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	if len(matrix) == 0 {
		return 0
	}
	res := 0
	cur := make([]int, len(matrix[0]))
	for i := 0; i < len(cur); i++ {
		if matrix[0][i] == '1' {
			cur[i] = 1
		} else {
			cur[i] = 0
		}
	}
	res = max(res, largestRectangleArea(cur))
	for i := 1; i < len(matrix); i++ {
		//update the cur
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				cur[j] += 1
			} else {
				cur[j] = 0
			}
		}
		res = max(res, largestRectangleArea(cur))
	}
	return res
}
