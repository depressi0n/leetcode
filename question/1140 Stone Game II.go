package question

func StoneGameII(piles []int) int {
	// 每次不再是取前后一个，而是取前后连续X个，1<=X<=2M,M=max(M,X)，初始值M=1
	// 返回Alice能获取的最大值
	// Alice优先选择，只能选M个，Bob可以选X=[1,2M]个,M=max(M,X)
	// dp[i][j] 表示首先从piles[i:]中选择前j+1个能获得的最大值 0<=j<2 0<=i<len(piles)
	// dp[i][j] = dp[i+1][0] ... dp[i+1][2*2MAX(j,M)] 这里的设置和题目的策略不一致，这里要求下一个能选上一个的两倍宽，但实际策略不是这样的

	//prefixSum:=make([]int,len(piles)+1)
	//prefixSum[0]=0
	//for i := 1; i <len(piles)+1 ; i++ {
	//	prefixSum[i]=prefixSum[i-1]+piles[i-1]
	//}
	//dp := make([][][]int, len(piles))
	//for i := 0; i < len(piles); i++ {
	//	dp[i] = make([][]int, len(piles))
	//}
	//// 初始化
	//dp[len(piles)-1][0] = []int{piles[len(piles)-1], 0}
	//// 状态转移
	//for i := len(piles) - 2; i >= 0; i-- {
	//	for j := 0; j < len(piles)-i; j++ {
	//		// 取前j个
	//		if i+j+1<len(piles){
	//			dp[i][j] = []int{prefixSum[i+j+1]-prefixSum[i-1+1] + dp[i+j+1][0][1], dp[i+j+1][0][0]}
	//			for k := 1; k <2*(j+1) && i+j+1+k<len(piles); k++ {
	//				// 对方的策略是选择前k+1个
	//				if dp[i][j][1]<dp[i+j+1][k][0]{
	//					dp[i][j]=[]int{prefixSum[i+j+1]-prefixSum[i-1+1]+dp[i+j+1][k][1],dp[i+j+1][k][0]}
	//				}
	//			}
	//		}else{
	//			dp[i][j] = []int{prefixSum[i+j+1]-prefixSum[i-1+1], 0}
	//		}
	//	}
	//}
	//// 准备答案
	//return dp[0][0][0]

	// 上面的策略是 X -> 1<=Y<=2X -> 2Y 的变化，需要对转移方程进行修正
	// 同时将prefixSum转为suffixSum
	// 题目的策略是M=max(M,X) 这个保证 当 1<=X<=M 时，还是M，而 M<X<=2*M 时 M=X
	// dp[i][m] 表示在[i:]中，M=m时，先手玩家可以获得的最多石子数
	// f[i][M]=sum[i:]
	// 转移方程
	// f[i][m] = sum[i:i+m] + sum[i+m:] -  f[i+m][max{M,x}] 1<=x<=2M =sum[i:] - f[i+m][max{M,x}]
	suffixSum := 0
	dp := make([][]int, len(piles))
	for i := len(piles) - 1; i >= 0; i-- {
		dp[i] = make([]int, len(piles)+1)
		suffixSum += piles[i]
		for M := 1; M <= len(piles); M++ {
			if i+2*M >= len(piles) {
				dp[i][M] = suffixSum
			} else {
				for x := 1; x <= 2*M; x++ {
					dp[i][M] = max(dp[i][M], suffixSum-dp[i+x][max(M, x)])
				}
			}
		}
	}
	return dp[0][1]
}
