package question

func MaxChunksToSorted(arr []int) int {
	//由于数组是0-n的排列, 故0-i间的最大值不会小于i, 记max(arr[0~i]) >= i
	//如果当前最大值大于当前位置i, 那么i的右边一定有小于当前最大值的元素, 则无法在i切分数组, 记max(arr[0~i] <= i)
	if len(arr) < 2 {
		return len(arr)
	}
	//目的是找出分割的数组中，最大值等于index的数，有的话就加1
	//count用来计算个数，cur用来表示最大的数
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	count := 0
	cur := arr[0]
	for i, v := range arr {
		cur = max(cur, v)
		//如果最大的数与对应的下标相等，那么说明排序后是从小到大，能够链接起来的
		if cur == i {
			count++
		}
	}
	return count

	// 基本思路是从第一个开始，往后找比他小的，直到没有更小的
	// 没有利用上本身是一个排列的变换
	// 最大值肯定在最后一个chunk中
	// 最小值肯定在第一个chunk中
	// 借助单调栈的结构：
	// 从后往前记录出现过的最小值的索引
	// 从前往后记录出现过的最大值的索引
	//minS := make([]int, len(arr))
	//minS[len(minS)-1] = len(arr) - 1
	//for i := len(arr) - 2; i >= 0; i-- {
	//	if arr[minS[i+1]] <= arr[i] {
	//		minS[i] = minS[i+1]
	//	} else {
	//		minS[i] = i
	//	}
	//}
	//maxS := make([]int, len(arr))
	//maxS[0] = 0
	//for i := 1; i < len(arr); i++ {
	//	if arr[maxS[i-1]] >= arr[i] {
	//		maxS[i] = maxS[i-1]
	//	} else {
	//		maxS[i] = i
	//	}
	//}
	//// 经过上面的操作，每个位置都是之后可能的最小值的索引
	//cnt := 0
	//// 记录前面一组的最大值，如果当前组的最大值比它小，则并称一组
	//// 找到后面比他小的，然后记录这一组的最大值，供后面比较
	//left, right := 0, 1
	//for right < len(minS) {
	//	// 如果minS(left)可以的话，尽可能不去扩展右边界
	//	right = minS[right-1] + 1
	//	// 判断这一段的最大值是否比后面一段的值还要大
	//	// 下一段
	//	tmpL:=right
	//	tmpR:=minS[right-1]+1
	//	if maxS[right-1] > maxS[tmpR-1]{
	//		right=tmpR
	//	}else{
	//		left=tmpL
	//		right=tmpR
	//		cnt++
	//	}
	//}
	//for i := len(minS)-2; i >=0; i-- {
	//	for i>=0 && minS[i]==minS[i+1]{
	//		i--
	//	}
	//	if i>=0{
	//		cnt++
	//	}
	//}
	//return cnt
}
