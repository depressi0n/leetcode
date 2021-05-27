package question

func StoneGameV(stoneValue []int) int {
	// 类比切西瓜
	// Alice负责切，Bob负责选
	// 如果不平均，则Bob选择丢弃多的，ALice得分是少的那一组
	// 如果两边相同，那么Alice可以选择留下哪一组
	// 下一轮使用剩下的一组
	// 动态规划的做法：
	// dp[i][j] //表示Alice在[i:j]的情况下可以获得的最大分数
	// 要求0<=i<len(stoneValue), i<j<=len(stoneValue)
	// 最后的结果是dp[0][len(stoneValue])
	prefixSum := make([]int, len(stoneValue)+1) // [0:j)的结果
	prefixSum[0] = 0
	dp := make([][]int, len(stoneValue))
	for i := 0; i < len(stoneValue); i++ {
		prefixSum[i+1] = prefixSum[i] + stoneValue[i]
		dp[i] = make([]int, len(stoneValue)+1)
	}
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	score := func(i, j, k int) int {
		if prefixSum[j]-prefixSum[k] < prefixSum[k]-prefixSum[i] {
			return prefixSum[j] - prefixSum[k] + dp[k][j]
		} else if prefixSum[j]-prefixSum[k] == prefixSum[k]-prefixSum[i] {
			return prefixSum[j] - prefixSum[k] + max(dp[i][k], dp[k][j])
		} else {
			return prefixSum[k] - prefixSum[i] + dp[i][k]
		}
	}
	// 转移方程
	// dp[i][j]=score(i,j,k)
	for length := 2; length <= len(stoneValue); length++ {
		for i := 0; i < len(stoneValue); i++ {
			if i+length > len(stoneValue) {
				break
			}
			for k := i + 1; k < i+length; k++ {
				dp[i][i+length] = max(dp[i][i+length], score(i, i+length, k))
			}
		}
	}
	return dp[0][len(stoneValue)]
}

// 优化策略，之前选择的一个划分，在后续划分时不再需要比较，而是按照之前的划分取出来
// TODO 将下面的细节看懂并实现
// if sum(l,r) < sum(l+1,r+1) => dp[l][r]=f[l][i] + sum(l,r)
// then for r+1 sum(l,i) < sum(i+1,r+1) => dp[l][r+1]=f[l][i]+sum(l,i)
// 可以将时间复杂度降低到二维
// 通过维护两个辅助数组 maxl 和 maxr
// maxl[l][r] = max_r^l(dp[l][i]+sum(l,i)
// maxr[l][r] = max_r^l(dp[i][r]+sum(i,r)
