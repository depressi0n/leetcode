package question

import "sort"

func StoneGameVI(aliceValues []int, bobValues []int) int {
	// Bob和Alice计算分数的方式不同，分别用bobValue和AliceValue表征
	// 每次都从数组中拿走一个
	// Alice先手
	// 思路：全部策略计算出来去比较肯定行不通，复杂度很高
	// 贪心算法：只有两个石头的情况下，Alice可能的选择有[0] or [1]
	// 更优的方案应该是a[0]-b[1]与a[1]-b[0]中更大的方案
	// 即a[0]-b[1]-(a[1]-b[0])=a[0]+b[0]-a[1]+b[1]
	// 贪心策略就是合并格各自的价值，每次取价值更大的那一组
	if len(aliceValues) == 0 {
		return 0
	}
	if len(aliceValues) == 1 {
		return 1
	}
	indexs := make([]int, len(aliceValues))
	for i := 0; i < len(indexs); i++ {
		indexs[i] = i
	}
	sort.Slice(indexs, func(i, j int) bool {
		return aliceValues[indexs[i]]+bobValues[indexs[i]] > aliceValues[indexs[j]]+bobValues[indexs[j]]
	})
	diff := 0
	for i := 0; i < len(indexs); i++ {
		diff += aliceValues[indexs[i]]
		i++
		if i >= len(indexs) {
			break
		}
		diff -= bobValues[indexs[i]]
	}
	if diff > 0 {
		return 1
	} else if diff == 0 {
		return 0
	} else {
		return -1
	}
}
