package question

//一个机器人位于一个 m x n网格的左上角 （起始点在下图中标记为 “Start” ）。
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//问总共有多少条不同的路径？

func uniquePaths(m int, n int) int {
	return uniquePathsCore(m,n)
}

// 动态规划的思路，从最后一个开始，dp[m-1][n-1]=1
// dp[i][j]=dp[i+1][j]+dp[i][j+1]
// 从后往前更新，可以改成一维的，因为只与之前的值和当前的值相关
// dp[i]=dp[i+1]+dp[i]
// 初始化dp[n-1]=1
// 更新m次
func uniquePathsCore(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}
	//res:=make([][]int,m)
	//for i:=0;i<m;i++{
	//	res[i]=make([]int,n)
	//}
	//res[m-1][n-1]=0
	//for i:=0;i<m-1;i++{
	//	res[i][n-1]=1
	//}
	//for i:=0;i<n-1;i++{
	//	res[m-1][i]=1
	//}
	//for i:=m-2;i>=0;i--{
	//	for j:=n-2;j>=0;j--{
	//		res[i][j]=res[i][j+1]+res[i+1][j]
	//	}
	//}
	//return res[0][0]
	//TODO：这里是完全可以节省空间的
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}
	for i := 0; i < m-1; i++ {
		for j := n - 2; j >= 0; j-- {
			res[j] = res[j] + res[j+1]
		}
	}
	return res[0]
}
