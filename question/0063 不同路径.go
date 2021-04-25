package question

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid == nil || len(obstacleGrid) == 0 {
		return 0
	}
	if len(obstacleGrid) == 1 && len(obstacleGrid[0]) == 1 && obstacleGrid[0][0] == 1 {
		return 0
	}
	if obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == 1 {
		return 0
	}
	if len(obstacleGrid) == 1 {
		for i := 0; i < len(obstacleGrid[0])-1; i++ {
			if obstacleGrid[0][i] == 1 {
				return 0
			}
		}
		return 1
	}
	if len(obstacleGrid[0]) == 1 {
		for i := 0; i < len(obstacleGrid)-1; i++ {
			if obstacleGrid[i][0] == 1 {
				return 0
			}
		}
		return 1
	}
	if obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == 1 {
		return 0
	}
	res := make([]int, len(obstacleGrid[0]))
	res[len(res)-1] = 0
	for i := len(res) - 2; i >= 0; i-- {
		if obstacleGrid[len(obstacleGrid)-1][i] != 1 {
			res[i] = 1
		} else {
			res[i] = 0
			i--
			for i >= 0 {
				res[i] = 0
				i--
			}
		}
	}
	for i := len(obstacleGrid) - 2; i >= 0; i-- {
		if obstacleGrid[i][len(res)-1] == 1 {
			res[len(res)-1] = 0
		} else if i == len(obstacleGrid)-2 {
			res[len(res)-1] = 1
		} else if res[len(res)-1] == 0 {
			res[len(res)-1] = 0
		} else {
			res[len(res)-1] = 1
		}
		for j := len(res) - 2; j >= 0; j-- {
			if obstacleGrid[i][j] != 1 {
				res[j] = res[j] + res[j+1]
			} else {
				res[j] = 0
			}
		}
	}
	return res[0]
}
