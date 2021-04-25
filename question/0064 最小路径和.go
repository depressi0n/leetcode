package question

func minPathSum(grid [][]int) int {
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
