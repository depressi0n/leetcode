package question

func Answer0274(citations []int) int {
	return hIndex(citations)
}
func hIndex(citations []int) int {
	// 目标：找出citations中的一个值，满足大于等于它的值的数目刚好等于它
	// A scientist has an index h if h of their n papers have at least h citations each,
	// and the other n − h papers have no more than h citations each.
	// 思想：计数排序 or 桶排序
	// 先排序
	//res:=0
	//sort.Ints(citations)
	//// 可能取值是0->len(citations)-1,直到找到citations[i]<=h-index为止
	//for i := len(citations) - 1; i >= 0 && citations[i] > res; i-- {
	//	res++
	//}
	//return res

	//优化：计数排序，大于n的可以不考虑，结果不可能比n还大，所以直接全部放到n里面就可以
	//可能的取值[0,n+1)
	n := len(citations)
	counter := make([]int, n+1)
	for _, citation := range citations {
		if citation >= n { //比n大
			counter[n]++
		} else { // 比n小
			counter[citation]++
		}
	}
	// 累加比x要大于等于的值，如果数目比x还大，说明找到了
	for i, tot := n, 0; i >= 0; i-- {
		tot += counter[i]
		if tot >= i {
			return i
		}
	}
	return 0
}
