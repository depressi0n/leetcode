package question

func StoneGameVII(stones []int) int {
	// 因为只能最左或者最右，所以可能存在动态规划的解
	// dp[i][j]表示先手的人在[i:j)中能获得的最大分差
	//prefixSum:=make([]int,len(stones)+1)
	//prefixSum[0]=0
	//dp:=make([][]int,len(stones))
	//for i := 0; i < len(stones); i++ {
	//	prefixSum[i+1]=prefixSum[i]+stones[i]
	//	dp[i]=make([]int,len(stones)+1)
	//	dp[i][i+1]=0
	//}
	//for length := 2; length <=len(stones); length++ {
	//	for i := 0; i < len(stones); i++ {
	//		if i+length>len(stones){
	//			break
	//		}
	//		dp[i][i+length]= func(a,b int)int{
	//			if a<b{
	//				return b
	//			}
	//			return a
	//		}(prefixSum[i+length]-prefixSum[i+1]-dp[i+1][i+length],prefixSum[i+length-1]-prefixSum[i]-dp[i][i+length-1])
	//	}
	//}
	//return dp[0][len(stones)]
	// TODO:可以省空间
	// length = 2 start
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	dp := make([]int, len(stones)-1)
	prefixSum := make([]int, len(stones)+1)
	prefixSum[0] = 0
	// 初始化
	for i := 0; i < len(stones)-1; i++ {
		dp[i] = max(stones[i], stones[i+1])
		prefixSum[i+1] = prefixSum[i] + stones[i]
	}
	prefixSum[len(stones)] = prefixSum[len(stones)-1] + stones[len(stones)-1]
	// 状态转移方程
	// dp[i] =max( Sum1-dp[i] ,Sum2 -dp[i+1])
	for length := 3; length <= len(stones); length++ {
		for i := 0; i < len(stones)-length+1; i++ {
			dp[i] = max(prefixSum[i+length]-prefixSum[i+1]-dp[i+1], prefixSum[i+length-1]-prefixSum[i]-dp[i])
		}
	}
	return dp[0]
}
