package question

// 给定一个 n×n 的二维矩阵matrix 表示一个图像。请你将图像顺时针旋转 90 度。
//
//你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

func rotate0048(matrix [][]int) {
	rotate0048Core(matrix)
}

func rotate0048Core(matrix [][]int) {
	// 顺时针旋转90度
	for i := 0; i < len(matrix)/2; i++ {
		for j := i; j < len(matrix[0])-i-1; j++ {
			matrix[i][j], matrix[j][len(matrix[0])-i-1],matrix[len(matrix[0])-i-1][len(matrix[0])-j-1],matrix[len(matrix[0])-j-1][i] = matrix[len(matrix[0])-j-1][i],matrix[i][j], matrix[j][len(matrix[0])-i-1],matrix[len(matrix[0])-i-1][len(matrix[0])-j-1]
		}
	}
	//log.Println(matrix)
	return
}
