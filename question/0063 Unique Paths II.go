package question
// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
//
//现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	return uniquePathsWithObstaclesCore2(obstacleGrid)
}

// 思路同上一题，从后往前计算，如果遇到障碍物则将该位置的值设置为0
func uniquePathsWithObstaclesCore1(obstacleGrid [][]int) int {
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

func uniquePathsWithObstaclesCore2(obstacleGrid [][]int) int {
	dp:=make([]int,len(obstacleGrid[0]))
	if obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[len(obstacleGrid)-1])-1]==1{
		return 0
	}
	dp[len(obstacleGrid[0])-1]=1
	for i:=len(obstacleGrid)-1;i>=0;i--{
		if obstacleGrid[i][len(obstacleGrid[i])-1]==1{
			dp[len(obstacleGrid[i])-1]=0
		}
		for j := len(obstacleGrid[i])-2; j >=0; j-- {
			if obstacleGrid[i][j] == 1 {
				dp[j]=0
			}else{
				dp[j]=dp[j]+dp[j+1]
			}
		}
	}
	return dp[0]
}
