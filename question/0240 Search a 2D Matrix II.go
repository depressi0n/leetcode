package question

// 编写一个高效的算法来搜索mxn矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
//每行的元素从左到右升序排列。
//每列的元素从上到下升序排列。

func searchMatrix0240(matrix [][]int, target int) bool {
	return searchMatrix0240Core2(matrix, target)
}

// 对行和列做二分查找
func searchMatrix0240Core(matrix [][]int, target int) bool {
	rowStart, rowEnd := 0, len(matrix)
	// 使用二分查找
	// 首先从行入手，找到可能的所在行
	for rowStart < rowEnd {
		mid := (rowEnd-rowStart)>>1 + rowStart
		if matrix[mid][0] < target {
			rowStart = mid + 1
		} else if matrix[mid][0] == target {
			return true
		} else {
			rowEnd = mid
		}
	}
	// 所有的行都有可能
	for i := rowStart - 1; i >= 0; i-- {
		colStart, colEnd := 0, len(matrix[i])
		// 然后从列入手，找到可能的所在列
		for colStart < colEnd {
			mid := (colEnd-colStart)>>1 + colStart
			if matrix[i][mid] < target {
				colStart = mid + 1
			} else if matrix[i][mid] == target {
				return true
			} else {
				colEnd = mid
			}
		}
	}
	return false
}

// 根据题设可以直接查找一个矩形右下角元素，其他元素必然比其小

func searchMatrix0240Core2(matrix [][]int, target int) bool {
	row, col := 0, len(matrix[0])-1
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}
