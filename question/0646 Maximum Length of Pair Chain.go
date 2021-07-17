package question

import "sort"

func Answer0646(pairs [][]int) int {
	return findLongestChain2(pairs)
}
func findLongestChain1(pairs [][]int) int {
	if len(pairs) == 0 {
		return 0
	}
	sort.Slice(pairs, func(i int, j int) bool {
		if pairs[i][0] != pairs[j][0] {
			return pairs[i][0] < pairs[j][0]
		} else {
			return pairs[i][1] < pairs[j][1]
		}
	})
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	// 动态规划
	dp := make([]int, len(pairs))
	// dp[i] 表示 以pairs[i]开始，最长的chain
	dp[len(pairs)-1] = 1
	res := 1
	for i := len(pairs) - 2; i >= 0; i-- {
		for j := i + 1; j < len(pairs); j++ {
			if pairs[i][1] < pairs[j][0] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
func findLongestChain2(pairs [][]int) int {
	if len(pairs) == 0 {
		return 0
	}
	// 按照右端点升序
	// 贪心正确性
	// （1）左边比之前的右边端点还小，不会被选择
	// （2）右边已经是尽可能小
	// （3）下一次选择也要使右边尽可能小，这样才能使链最长
	sort.Slice(pairs, func(i int, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	curRight := pairs[0][1]
	res := 1
	for i := 0; i < len(pairs); i++ {
		if curRight < pairs[i][0] {
			res++
			curRight = pairs[i][1]
		}
	}
	return res
}
