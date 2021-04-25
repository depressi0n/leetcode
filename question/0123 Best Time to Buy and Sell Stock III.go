package question

import "fmt"

func maxProfit123(prices []int) int {
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

//TODO:有更优化的解法
