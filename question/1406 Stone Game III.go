package question

func StoneGameIII(stoneValue []int) string {
	// 每次可以从最前面拿走1，2，3堆石子，其中可能有负数
	// dp[i][j] 表示stoneValue[i:]中取走j+1堆石子能获得最高分数
	// 0<=i<=len(stoneValue),0<=j<=2,i+j<len(stoneValue)
	//if len(stoneValue) <= 3 {
	//	return "Alice"
	//}
	//// 用一个后缀数组进行管理
	//dp := make([][]int, len(stoneValue))
	//for i := 0; i < len(stoneValue); i++ {
	//	dp[i] = make([]int, 3)
	//}
	//// 初始化
	//suffixSum := stoneValue[len(stoneValue)-1]
	//dp[len(stoneValue)-1][0] = stoneValue[len(stoneValue)-1]
	//dp[len(stoneValue)-1][1] = stoneValue[len(stoneValue)-1]
	//dp[len(stoneValue)-1][2] = stoneValue[len(stoneValue)-1]
	//for i := len(stoneValue) - 2; i >= 0; i-- {
	//	suffixSum += stoneValue[i]
	//	for j := 0; j < 3 ; j++ {
	//		if i+j+1 < len(stoneValue) {
	//			dp[i][j] = suffixSum - max(max(dp[i+j+1][0], dp[i+j+1][1]), dp[i+j+1][2])
	//		} else if i+j+1==len(stoneValue) {
	//			dp[i][j] = suffixSum
	//		}
	//	}
	//}
	//flag := 0x0
	//for i := 0; i < 3; i++ {
	//	if dp[0][i]*2 > suffixSum {
	//		flag |= 1 << 2
	//	} else if dp[0][i]*2 == suffixSum {
	//		flag |= 1 << 1
	//	} else {
	//		flag |= 1 << 0
	//	}
	//}
	//if flag >= 4 {
	//	return "Alice"
	//} else if flag >= 2 {
	//	return "Tie"
	//} else {
	//	return "Bob"
	//}

	// dp[i] 表示 先手玩家 在还剩下[i:]时应该选择拿走的石子堆数
	// 转移方程为 dp[i]= max(sum(i,n-1)-f[j])=sum(i,n-1)-min(f[j]), 0<=j<=2
	dp := make([]int, len(stoneValue)+1)
	dp[len(stoneValue)] = 0
	suffixSum := 0
	for i := len(stoneValue) - 1; i >= 0; i-- {
		suffixSum += stoneValue[i]
		bestj := dp[i+1]
		for j := i + 2; j <= i+3 && j <= len(stoneValue); j++ {
			bestj = min(bestj, dp[j])
		}
		dp[i] = suffixSum - bestj
	}
	if dp[0]*2 > suffixSum {
		return "Alice"
	} else if dp[0]*2 == suffixSum {
		return "Tie"
	} else {
		return "Bob"
	}

}
