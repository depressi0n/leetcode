package question

import "fmt"

// 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
//设计一个算法来计算你所能获取的最大利润。你最多可以完成两笔交易。
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
func maxProfit0123(prices []int) int {
	return maxProfit0123Core1(prices)
}

// dp[i][j][2]int 在第i天到第j天做1/2次交易的最大收益
// 初始化
// dp[i][i+2][1]=prices[i+1]-prices[i] if prices[i+1]>prices[i] else 0
// dp[i][i+2][2]=dp[i][i+2][1]
// 状态转移方程
// dp[i][j+1][1]=max(dp[i][k][1],dp[k][j+1][1]) for all i+2<=k<=j
// dp[i][j+1][2]=max(dp[i][k][2],dp[k][j+1][2],dp[i][k][1]+dp[k][j+1][2]) for all i+2<=k<=j
func maxProfit0123Core1(prices []int) int {
	dp := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][]int, len(prices))
		for j := i; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	// 初始化
	for i := 0; i+1 < len(prices); i++ {
		if prices[i+1] > prices[i] {
			dp[i][i+1][0] = prices[i+1] - prices[i]
			dp[i][i+1][1] = prices[i+1] - prices[i]
		}
	}
	for length := 2; length <= len(prices); length++ {
		for i := 0; i+length <= len(prices); i++ {
			for sublength := 2; sublength < length; sublength++ {
				dp[i][i+length-1][0] = max(dp[i][i+length-1][0], max(max(dp[i][i+sublength-1][0], dp[i+sublength-1][i+length-1][0]), prices[i+length-1]-prices[i]))
				dp[i][i+length-1][1] = max(dp[i][i+length-1][1], max(max(dp[i][i+sublength-1][1], dp[i+sublength-1][i+length-1][1]), dp[i][i+sublength-1][0]+dp[i+sublength-1][i+length-1][0]))
			}
		}
	}

	//for length := 4; length <= len(prices); length++ {
	//	for i := 0; i+length <= len(prices); i++ {
	//		for sublength := 2; sublength < length && length-sublength >= 2; sublength++ {
	//			dp[i][i+length][0] = max(dp[i][i+length][0], max(dp[i][i+sublength][0], dp[i+sublength][i+length][0]))
	//			dp[i][i+length][1] = max(max(dp[i][i+sublength][1], dp[i+sublength][i+length][1]), dp[i][i+sublength][0]+dp[i+sublength][i+length][0])
	//		}
	//	}
	//}
	return dp[0][len(prices)-1][1]
}
func maxProfit0123Core2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	//可以省空间第
	min, max := 0, 0
	start := make([]int, len(prices))
	end := make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		// 更新是有条件的，否则10，3，7，1这种就会出错
		end[i] = end[i-1]
		if prices[i] < prices[min] {
			min = i
			max = i
		}

	}
	max, min = len(prices)-1, len(prices)-1
	for i := len(prices) - 2; i >= 0; i-- {
	}
	//start[i] 表示第i天开始
	for i := 0; i < len(prices); i++ {
		if prices[max] > prices[min] {
			start[i] = prices[max] - prices[min]
		} else {
			start[i] = 0
		}
	}
	fmt.Printf("%v\n", start)
	res := 0
	for i := 0; i < len(prices); i++ {
		if end[i]+start[i] > res {
			res = end[i] + start[i]
		}
	}
	return res
}

// 别人提供一个完整的思路
// 1、从前往后扫，找到end[i]前i天获得的最大利润，需要维护一个买进值
// 2、从后往前扫，找出start[i]表示从i到最后一天的最大利润，需要维护一个卖出值
func maxProfit123_1(prices []int) int {
	//从前往后维护end[i]
	min := 0
	end := make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[min] {
			min = i
		}
		if prices[i]-prices[min] > end[i-1] {
			end[i] = prices[i] - prices[min]
		} else {
			end[i] = end[i-1]
		}
	}
	//max:=len(prices)-1
	//start:=make([]int,len(prices))
	//for i:=len(prices)-2;i>=0;i--{
	//	if prices[i]>prices[max]{
	//		max=i
	//	}
	//	start[i]=prices[max]-prices[i]
	//}
	//res:=0
	//for i:=0;i<len(prices);i++{
	//	if end[i]+start[i]>res{
	//		res=end[i]+start[i]
	//	}
	//}
	max := len(prices) - 1
	start, res := 0, 0
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > prices[max] {
			max = i
		}
		start = prices[max] - prices[i]
		if start+end[i] > res {
			res = start + end[i]
		}
	}
	return res
}

// 拆分的本质找到一个中间点，然后分别求两边的最大值，最后和最大的入选
//（1）可以维护[0,i]表示前i天的最大利润（可以维护）和[i,len-1]的最大利润（不好维护）
//（2）可以维护[0,i]表示第i天卖出的最大利润（可以维护）和[i,len-1]的最大利润（不好维护）
//（3）可以维护[0,i]表示前i天的最大利润（可以维护）和[i,len-1]表示第i天买入的最大利润（可以维护）
//（4）最后这种直观的组合是错的，因为没有规定必须在第i天买入后卖出

func maxProfit0123Core3(prices []int) int {
	// profits[i] 表示前i-1天的最大利润，从前往后维护一个最小值
	lowprice, highprice := prices[0], prices[0]
	profits := make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		profits[i] = max(profits[i-1], prices[i]-lowprice)
		if prices[i] < lowprice {
			lowprice = prices[i]
		}
	}
	// starts[i] 表示第i天买入第最大利润，需要从后往前维护最大值
	highprice = prices[len(prices)-1]
	starts := make([]int, len(prices))
	for i := len(prices) - 2; i >= 0; i-- {
		starts[i] = highprice - prices[i]
		if prices[i] > highprice {
			highprice = prices[i]
		}
	}
	res := 0
	for i := 0; i < len(prices)-1; i++ {
		res = max(res, profits[i]+starts[i])
	}
	return res
}

//TODO:有更优化的解法
// 到第i天为止有五个状态
// （1）什么都没做，利润恒定为0
// （2）只买过一次，利润此时为负数，用buy1
// （3）完成了一笔交易，此时可能有利润，也可能没有，用sell1
// （4）完成了一笔交易并第二次买过，此时可能有利润，可能没有，用buy2
// （5）完成两笔交易，最终结果，用sell2
// 第一天的状态可能是以上状态之一，buy1=-prices[0],sell1=0,buy2=-prices[0],sell2=0
// 往后更新即可
//
func maxProfit0123Core4(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i]) // 可以转移到此状态的是从（1）买入或者（2）什么也不做
		sell1 = max(sell1, buy1+prices[i]) // 转移到此状态 可以从（2）卖出 或 （3）什么也不做
		buy2 = max(buy2, sell1-prices[i]) // （3）买入 或 （4） 什么也不做
		sell2 = max(sell2, buy2+prices[i]) // （4）卖出 或 （5）什么也不做
	}
	return sell2
}
