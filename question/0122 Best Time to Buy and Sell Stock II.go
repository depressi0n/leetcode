package question

// 给定一个数组 prices ，其中prices[i] 是一支给定股票第 i 天的价格。
//设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
func maxProfit0122(prices []int) int {
	return maxProfit0122Core(prices)
}
// 主要思想：赚取每一次差价，贪心的思路
// res+=prices[i]-prices[i-1] if prices[i]>prices[i-1]
// 等价于每天都在买卖，如果明天的价格比今天低，那么今天买入就卖出，否则保留
func maxProfit0122Core(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}
