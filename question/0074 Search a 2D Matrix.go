package question

// 编写一个高效的算法来判断m x n矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//每行中的整数从左到右按升序排列。
//每行的第一个整数大于前一行的最后一个整数。
func searchMatrix0074(matrix [][]int, target int) bool {
	return searchMatrixCore(matrix, target)
}

//每行和每列都可以进行二分查找，首先用列二分查找定位行，再用行二分查找定位列
func searchMatrixCore(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	// 纵向二分查找
	start, end, mid := 0, len(matrix)-1, 0
	for start <= end { //因为上面的end是有效的，才能用等于
		mid = (start + end) / 2
		if matrix[mid][0] < target {
			start = mid + 1
		} else if matrix[mid][0] > target {
			end = mid - 1
		} else {
			return true
		}
	}
	//跳出来以后肯定有[end]<target<[start]
	if end < 0 {
		return false
	}
	row := end
	//横向二分查找
	start, end, mid = 0, len(matrix[0])-1, 0
	for start <= end {
		mid = (start + end) / 2
		if matrix[row][mid] < target {
			start = mid + 1
		} else if matrix[row][mid] > target {
			end = mid - 1
		} else {
			return true
		}
	}
	return false
}
