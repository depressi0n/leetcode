package question

// 给定一个数组 prices ，它的第i个元素prices[i] 表示一支给定股票第 i 天的价格。
//你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
//返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
func maxProfit0121(prices []int) int {
	return maxProfitCore(prices)
}

// 主要思想：某一天买入然后以后卖出，只能买卖一次
// 本质上是要 max(prices[i]-prices[j]) \forall{i>j}
// 记录当前遇到的最小值，然后用当前值减去最小值，更新最大值
func maxProfitCore(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	res, min := 0, 0

	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[min] > res {
			res = prices[i] - prices[min]
		}
		if prices[i] < prices[min] {
			min = i
		}
	}
	return res
}
