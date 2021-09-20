package question

import "math"

// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//说明：每次只能向下或者向右移动一步。

func minPathSum(grid [][]int) int {
	return minPathSumCore2(grid)
}
func minPathSumCore1(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	if len(grid) == 1 {
		res := 0
		for i := 0; i < len(grid[0]); i++ {
			res += grid[0][i]
		}
		return res
	}
	if len(grid[0]) == 1 {
		res := 0
		for i := 0; i < len(grid); i++ {
			res += grid[i][0]
		}
		return res
	}
	for i := len(grid) - 2; i >= 0; i-- {
		grid[i][len(grid[0])-1] = grid[i][len(grid[0])-1] + grid[i+1][len(grid[0])-1]
	}
	for i := len(grid[0]) - 2; i >= 0; i-- {
		grid[len(grid)-1][i] = grid[len(grid)-1][i] + grid[len(grid)-1][i+1]
	}
	for i := len(grid) - 2; i >= 0; i-- {
		for j := len(grid[0]) - 2; j >= 0; j-- {
			grid[i][j] = grid[i][j] + min(grid[i][j+1], grid[i+1][j])
		}
	}
	return grid[0][0]
}
func minPathSumCore2(grid [][]int) int {
	for i := len(grid)-1; i >=0 ; i-- {
		for j := len(grid[i])-1; j >=0 ; j-- {
			pre:=grid[i][j]
			grid[i][j]=math.MaxInt32
			if i+1<len(grid){
				grid[i][j]=min(grid[i][j],grid[i+1][j])
			}
			if j+1<len(grid[i]){
				grid[i][j]=min(grid[i][j],grid[i][j+1])
			}
			if grid[i][j]!=math.MaxInt32{
				grid[i][j]+=pre
			}else{
				grid[i][j]=pre
			}
		}
	}
	return grid[0][0]
}