package question

func CanEat(candiesCount []int, queries [][]int) []bool {
	res := make([]bool, 0, len(queries))
	// 在queries[i][1]天吃queries[i][0]类型的，每天吃[1,queries[i][2]]个
	// 前缀数组
	prefixSum := make([]int, len(candiesCount)+1)
	prefixSum[0] = 0
	for i := 1; i < len(candiesCount)+1; i++ {
		prefixSum[i] = prefixSum[i-1] + candiesCount[i-1]
	}
	for i := 0; i < len(queries); i++ {
		// 首先判断在这之前的天数中，最多和最少吃掉多少颗
		minCnt, maxCnt := queries[i][1], queries[i][1]*queries[i][2]
		// 比较第1到i种类型第糖果是否够之前的天数所吃
		// 取有交集的情况为true
		if minCnt < prefixSum[queries[i][0]] {
			if minCnt+queries[i][2] > prefixSum[queries[i][0]] {
				res = append(res, true)
			} else if maxCnt+queries[i][2] > prefixSum[queries[i][0]] {
				res = append(res, true)
			} else {
				res = append(res, false)
			}
		} else if minCnt < prefixSum[queries[i][0]+1] {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}
