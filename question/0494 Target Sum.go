package question

func FindTargetSumWays(nums []int, target int) int {
	// 目标：在nums[i]前添加正负号，使得整体的和为target的不同方式
	// 基本思路：
	// 本质上是将nums[i]中的每个值取正负号，然后得出target，看看一共有多少种方式
	// 使用动态规划：
	// dp[i][j] 表示nums[:i]可以组合成j的方式有几种
	// 结果是dp[len(nums)][target]的结果
	// 这里考虑一下范围
	// nums[i] 在[0,1000]，一共不超过20个
	// target 在[-1000,1000]
	// 0 <= sum(nums[i]) <= 1000 这个条件可以用于缩小前面的和的范围在[-1000,1000]
	// 所以dp的第二维可以使用[-2000,2000]
	dp := make([][]int, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		dp[i] = make([]int, 4001)
	}
	// 初始化
	// dp[0][j]=0
	dp[1][-nums[0]+2000] += 1
	dp[1][nums[0]+2000] += 1
	// dp[i+1][j]=dp[i][j-nums[i]]+dp[i][j+nums[i]]
	for i := 2; i < len(nums)+1; i++ {
		for j := 0; j < 4001; j++ {
			if j-nums[i-1] >= 0 {
				dp[i][j] += dp[i-1][j-nums[i-1]]
			}
			if j+nums[i-1] < 4001 {
				dp[i][j] += dp[i-1][j+nums[i-1]]
			}
		}
	}
	return dp[len(nums)][target+2000]
}

//TODO 利用前缀和(sum)和变成负数的和(neg)
// target = sum - 2*neg
// neg = (sum - target)/2
// 此时neg必定大于0，则sum也必定大于target，如果不满足，则可以直接返回0
// 此时求target变成了挑选指定的数字使得这些数字的和是neg，
// 并且可以将二维优化为一维
//func findTargetSumWays(nums []int, target int) int {
//    sum := 0
//    for _, v := range nums {
//        sum += v
//    }
//    diff := sum - target
//    if diff < 0 || diff%2 == 1 {
//        return 0
//    }
//    n, neg := len(nums), diff/2
//    dp := make([][]int, n+1)
//    for i := range dp {
//        dp[i] = make([]int, neg+1)
//    }
//    dp[0][0] = 1
//    for i, num := range nums {
//        for j := 0; j <= neg; j++ {
//            dp[i+1][j] = dp[i][j]
//            if j >= num {
//                dp[i+1][j] += dp[i][j-num]
//            }
//        }
//    }
//    return dp[n][neg]
//}
//
