package question

func getRow(rowIndex int) []int {
	switch {
	case rowIndex < 0:
		return nil
	case rowIndex == 0:
		return []int{1}
	default:
		res := make([]int, rowIndex+1)
		//dp[i][j]=dp[i-1][j-1]+dp[i-1][j] 滚动数组
		res[0] = 1
		res[1] = 1
		res[rowIndex] = 1
		var tmp int
		for length := 3; length <= rowIndex+1; length++ {
			tmp = 1
			for i := 1; i < length-1; i++ {
				tmp, res[i] = res[i], tmp+res[i]
			}
			res[length-1] = 1
		}

		return res
	}
}
