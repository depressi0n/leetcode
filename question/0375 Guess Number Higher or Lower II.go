package question

// 我们正在玩一个猜数游戏，游戏规则如下：
//
//我从1到 n 之间选择一个数字。
//你来猜我选了哪个数字。
//如果你猜到正确的数字，就会 赢得游戏 。
//如果你猜错了，那么我会告诉你，我选的数字比你的 更大或者更小 ，并且你需要继续猜数。
//每当你猜了数字 x 并且猜错了的时候，你需要支付金额为 x 的现金。如果你花光了钱，就会 输掉游戏 。
//给你一个特定的数字 n ，返回能够 确保你获胜 的最小现金数，不管我选择那个数字 。
func getMoneyAmount(n int) int {
	return getMoneyAmountCore(n)
}

// 像一个动态规划，从1开始
// 无论对方选择哪个数字，要确保能获胜的最小现金数
func getMoneyAmountCore(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	// 错误思路：错误原因在于[1,8] 和 [10,17]不能按照同一个长度去比较！！！
	// 建立的两棵树如下
	//      5                 14
	//   3     7          12      16
	// 1   4 6   8     10    13 15   17
	//   2   			  11
	// 但选择时候并不是同样的路径
	// 前者选择的是 5+7 > 5 + 3 +1
	// 后者选择的是 14 + 12 + 10 > 14 + 16
	// 下面的动态规划并不满足可以利用小规模的特性！！ => 这也说明必须增加维度，重新划分状态
	//// dp[i]表示在长度为i的序列中获胜的最小现金数目
	//dp := make([][]int, n+1)
	//dp[0] = []int{0, 0}
	//// 只有一个，必胜
	//dp[1] = []int{0, 0}
	//// 有两个时，选择小的
	//dp[2] = []int{1, 1}
	//for i := 3; i < n+1; i++ {
	//	// 最平凡的行为就是从小猜到大
	//	dp[i] = []int{i * (i + 1) / 2, i}
	//	for j := i; j >= 1; j-- {
	//		// 以j作为根
	//		// 右侧如果长度不为0或1，则还要加上j*
	//		left := j + dp[j-1][0]
	//		levelLeft := dp[j-1][1] + 1
	//		right := j + dp[i-j][0] + j*dp[i-j][1]
	//		levelRight := dp[i-j][1] + 1
	//		if (left < right || (left == right && levelLeft <= levelRight)) && (right < dp[i][0] || (right == dp[i][0] && levelRight < dp[i][1])) {
	//			dp[i][0] = right
	//			dp[i][1] = levelRight
	//		}
	//		if (right < left || (right == left && levelRight <= levelLeft)) && (left < dp[i][0] || (left == dp[i][0] && levelLeft < dp[i][1])) {
	//			dp[i][0] = left
	//			dp[i][1] = levelLeft
	//		}
	//	}
	//}
	//tmp := make([][]int, n+1)
	//for i := 1; i < n+1; i++ {
	//	tmp[i] = make([]int, n+1)
	//	tmp[i][i] = 0
	//	for j := i + 1; j < n+1; j++ {
	//		tmp[i][j] = dp[j-i+1][0] + (i-1)*dp[j-i+1][1]
	//	}
	//}
	//for i := 1; i < n+1; i++ {
	//	fmt.Printf("[%d]",i)
	//	for j := i; j < n+1; j++ {
	//		fmt.Printf("%d,", tmp[i][j])
	//	}
	//	fmt.Println()
	//}

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
		// 只有一个元素
		dp[i][i] = 0
	}
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j < n+1; j++ {
			minCost := (i + j)*(j-i+1) / 2
			for k := i; k < j; k++ {
				cost:=k+max(dp[i][k-1],dp[k+1][j])
				minCost=min(minCost,cost)
			}
			dp[i][j]=minCost
		}
	}
	//for i := 1; i < n+1; i++ {
	//	fmt.Printf("[%d]",i)
	//	for j := i; j < n+1; j++ {
	//		fmt.Printf("%d,", dp[i][j])
	//	}
	//	fmt.Println()
	//}
	return dp[1][n]
}
