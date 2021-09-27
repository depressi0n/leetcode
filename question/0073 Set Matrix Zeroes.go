package question

// 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

func setZeroes(matrix [][]int) {
	//fmt.Println(matrix)
	setZeroesCore2(matrix)
	//fmt.Println(matrix)
}

//第一种方案：利用一个标记矩阵，标记原始矩阵中0的位置，然后遍历修改
func setZeroesCore1(matrix [][]int) {
	flag := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		flag[i] = make([]bool, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				flag[i][j] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if flag[i][j] {
				for k := 0; k < len(matrix[i]); k++ {
					matrix[i][k] = 0
				}
				for k := 0; k < len(matrix); k++ {
					matrix[k][j] = 0
				}
			}
		}
	}

}

//第二种方案：用两个数组记录要被置零的行和列
func setZeroesCore2(matrix [][]int) {
	row:=make([]bool,len(matrix))
	col:=make([]bool,len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				row[i]=true
				col[j]=true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		if row[i] {
			for j := 0; j < len(matrix[i]); j++ {
				matrix[i][j]=0
			}
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		if col[i] {
			for j := 0; j < len(matrix); j++ {
				matrix[j][i]=0
			}
		}
	}
}

//第三种方案：用所在行列最前面的位置记录当前将要被置零的行和列，即将辅助数组借助原本的空间表示
// 需要处理的情况：首行或首列中有为0的情况，尤其是（0，0）为0的情况，借助两个变量row和col分别表示首行和首列是否需要为空
// 对于(0,0)的情况，单独处理
// 可以优化为只使用一个变量，用matrix[0][0]记录首行（列）是否需要清零，而变量用于记录首列（行）是否需要清零
func setZeroesCore3(matrix [][]int) {
	//第三种方案
	column := false
	row := false
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					row = true
					continue
				}
				if j == 0 {
					column = true
					continue
				}
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	//首行和首列需要最后处理
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			for j := 1; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}
	if matrix[0][0] == 0 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	} else {
		if row {
			for i := 0; i < len(matrix[0]); i++ {
				matrix[0][i] = 0
			}
		}
		if column {
			for i := 0; i < len(matrix); i++ {
				matrix[i][0] = 0
			}
		}
	}
	return
}
