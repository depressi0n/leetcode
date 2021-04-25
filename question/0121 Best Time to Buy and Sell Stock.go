package question

func maxProfit(prices []int) int {
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
