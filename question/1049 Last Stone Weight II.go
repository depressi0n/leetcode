package question

func LastStoneWeightII(stones []int) int {
	// 每次选择两个石头，如果重量相同则两个都销毁，如果重量不同，则小的销毁，大的重量减去这个重量
	// 返回最小的可能重量，如果没有石头剩下，则返回0

	// 类似于在不同的石头上加上正负号
	// 求和得到结果

	// 如果可以为0则返回0，否则返回和的绝对值
	// 思路是加上负号的值的和为neg，则最后的结果是sum-2*neg，保证这个绝对值最小
	// 考虑动态规划
	// dp[i][j]表示nums[:i]中neg为j
	// 从最后的neg中取出一个与sum进行作用，得出最后的结果
	// 速度很慢...
	//sum := func() int {
	//	res := 0
	//	for i := 0; i < len(stones); i++ {
	//		res += stones[i]
	//	}
	//	return res
	//}()
	//dp := make([]map[int]struct{}, len(stones)+1)
	////	初始化
	//dp[0] = map[int]struct{}{0: {}}
	//for i := 1; i < len(dp); i++ {
	//	dp[i] = make(map[int]struct{})
	//	for x, _ := range dp[i-1] {
	//		dp[i][x] = struct{}{}
	//		dp[i][x+stones[i-1]] = struct{}{}
	//	}
	//}
	//res := 101
	//for x, _ := range dp[len(stones)] {
	//	tmp := sum - 2*x
	//	if tmp==0{
	//		return 0
	//	}
	//	if 0 < tmp && tmp < res {
	//		res = tmp
	//	} else if tmp < 0 && -tmp < res {
	//		res = -tmp
	//	}
	//}
	//return res

	//	优化
	// 思路：
	// 如果没有石头剩下，则可以认为剩下了最后一块重量为0的石头
	// 什么情况下加上正负号是合法的呢？
	// 使用正负号将石头划分为两堆，两组石头重量之差的绝对值是所有划分当中最小的
	// 如果可以找到一组正负号，并且两组重量之差为diff，有一种粉碎方案，最后一块石头的重量也是diff，则这一组合法
	// 取出正号中最大的一块石头，负号中任意一个，如果没有完全粉碎，将剩下的石头放到正号中
	// 这个操作可以保证重量之差的绝对值在操作前后是不变的
	// 最后如果出现了一组为空，另一组不止一块石头的情况，那么与上面"最小"矛盾，因为一旦从非空的一组中移入一块石头到
	// 空的一组中，会导致绝对值减小，就将不再是"最小"的假设
	// 转换为0-1背包问题 dp[i][j] 表示stones[:i]能否凑出重量j
	// dp[i+1][j] = dp[i][j] or dp[i][j]|dp[i][j-stones[i]]
	// 由动态转移方程可以考虑将空间优化为一维
	sum := 0
	for _, v := range stones {
		sum += v
	}
	m := sum / 2
	dp := make([]bool, m+1)
	dp[0] = true
	for _, weight := range stones {
		for j := m; j >= weight; j-- {
			dp[j] = dp[j] || dp[j-weight]
		}
	}
	for j := m; ; j-- {
		if dp[j] {
			return sum - 2*j
		}
	}
}
