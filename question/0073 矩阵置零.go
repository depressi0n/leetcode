package question

//第一种方案：利用一个标记矩阵，标记原始矩阵中0的位置，然后遍历修改
//第二种方案：用两个数组记录要被置零的行和列
//第三种方案：用所在行列最前面的位置记录当前将要被置零的行和列
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	//第二种方案
	//var row,column []int
	//for i:=0;i<len(matrix);i++{
	//	for j:=0;j<len(matrix[0]);j++{
	//		if matrix[i][j]==0{
	//			row=append(row,i)
	//			column=append(column,j)
	//		}
	//	}
	//}
	//for i:=0;i<len(row);i++{
	//	for j:=0;j<len(matrix[0]);j++{
	//		matrix[row[i]][j]=0
	//	}
	//}
	//for i:=0;i<len(column);i++{
	//	for j:=0;j<len(matrix[0]);j++{
	//		matrix[j][column[i]]=0
	//	}
	//}
	//return

	//第三种方案
	//
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
	//[0][0]需要最后处理
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
