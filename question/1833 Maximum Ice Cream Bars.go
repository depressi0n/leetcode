package question

func Answer1833(costs []int, coins int) int {
	return maxIceCream(costs, coins)
}
func maxIceCream(costs []int, coins int) int {
	// 目标：选择尽可能多的数字
	// 思路：
	// （1）背包问题+剪枝，复杂度还是很高
	// （2）排序
	//sort.Ints(costs)
	//used:=0
	//cur:=0
	//for cur<len(costs) &&used+costs[cur]<=coins{
	//	used+=costs[cur]
	//	cur++
	//}
	//if cur == 0{
	//	return 0
	//}
	//if used<=coins{
	//	return cur
	//}
	//return cur-1
	// TODO(没写完):优化策略：快排思想，先排一部分，同时在快排的时候计算前半段的和
	// 比较，
	// 如果恰好等于，则直接返回
	// 如果大于，则往前
	// 如果小于，则往后
	//quickSort:= func(start int,end int,prefixSum int)(int,int) {
	//	qivot:=costs[start]
	//	left:=start
	//	right:=end-1
	//	for left<right{
	//		for costs[right]>qivot{
	//			right--
	//		}
	//		costs[left]=costs[right]
	//		prefixSum+=costs[left]
	//		left++
	//		for costs[left]<qivot{
	//			left++
	//		}
	//		costs[right]=costs[left]
	//		prefixSum+=costs[right]
	//		right--
	//	}
	//	costs[left]=qivot
	//	prefixSum+=qivot
	//	return left,prefixSum
	//}
	//start:=0
	//end:=len(costs)
	//curSum:=0
	//for start<end{
	//	index,sum:=quickSort(start,end,curSum)
	//	if sum == coins{
	//		return index
	//	}else if sum>coins{
	//	}else{
	//		curSum=sum
	//		start=index
	//	}
	//}
	//if start==0{
	//	return 0
	//}else{
	//	return start
	//}

	// 优化策略2：计数排序，和上面思想其实是等价的
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	ans := 0
	const mx int = 1e5    //借助最大值不超过10^5
	freq := [mx + 1]int{} //使用map计数
	for _, c := range costs {
		freq[c]++
	}
	//从大到小累加，直到不能累加为止
	for i := 1; i <= mx && coins >= i; i++ {
		c := min(freq[i], coins/i)
		ans += c
		coins -= i * c
	}
	return ans

}
