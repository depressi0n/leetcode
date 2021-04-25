package question

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	var dfs func(i int, j int)
	dfs = func(i, j int) {
		grid[i][j] = '2'
		if i+1 < len(grid) && grid[i+1][j] == '1' {
			dfs(i+1, j)
		}
		if j+1 < len(grid[0]) && grid[i][j+1] == '1' {
			dfs(i, j+1)
		}
		if i-1 >= 0 && grid[i-1][j] == '1' {
			dfs(i-1, j)
		}
		if j-1 >= 0 && grid[i][j-1] == '1' {
			dfs(i, j-1)
		}
	}
	res := 0
	// 遍历到一个1，则开始dfs把所有的1改成-1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				res++
			}
		}
	}
	return res
}
