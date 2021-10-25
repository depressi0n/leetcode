package question

// 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
//在「杨辉三角」中，每个数是它左上方和右上方的数的和。

func generate(numRows int) [][]int {
	return generateCore(numRows)
}

// 主要思想：dp[i][j]=dp[i-1][j]+dp[i-1][j-1] if j>0 else dp[i][j]=1
func generateCore(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}
	res := make([][]int, numRows)
	res[0] = []int{1}
	for i := 1; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		for j := 1; j < i; j++ {
			//fmt.Println(i, j)
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i][i] = 1
	}
	return res
}
