package question

// 考虑几种情况
//1. 整个数组总和的一半是多少，如果有超过的则不可能划分成两
//2. 总和为奇数的也不可能划分
//3. 剩下就是解决0-1背包问题：动态规划「可以压缩空间」或者dfs就能解决
func canPartition(nums []int) bool {
	if nums == nil || len(nums) < 2 {
		return false
	}
	max, sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		sum += nums[i]
	}
	if max<<1 > sum || sum&0x1 == 1 {
		return false
	}
	if max<<1 == sum {
		return true
	}
	target := sum >> 1 //还差这么多到一半
	// 利用动态规划解决受限制的背包问题
	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, target+1)
	}
	//dp[i][j]表示从nums[0]至nums[i]，能否选出和为j的序列
	// dp[i+1][j]=dp[i-1][j] 如果j<nums[i]的话
	// dp[i+1][j]=dp[i-1][j] || dp[i-1][j-nums[i]],前者是不选择nums[j],后者是选择nums[j]
	dp[0][0] = true
	if nums[0] < target+1 {
		dp[0][nums[0]] = true
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < target+1; j++ {
			dp[i][j] = dp[i-1][j]
			if nums[i] <= j {
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[len(nums)-1][target]
}
