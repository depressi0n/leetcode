package question

func Answer1846(arr []int) int {
	return MaximumElementAfterDecrementingAndRearranging(arr)
}
func MaximumElementAfterDecrementingAndRearranging(arr []int) int {
	// 为使数组arr满足以下条件
	// （1）arr[0]=1
	// （2）abs(arr[i]-arr[i-1])<=1
	// 可以执行两类操作：一是将任意一个值减小，二是以任意顺序排列数组
	// 返回满足条件时可能的最大值
	// 思路 似乎结果只是跟长度以及数组的最大值有关系，有如下理由：
	// 首先，值只可能减小，最好的可能就是不需要减小；
	// 其次，如果长度上不够，那么最大值只能减小，才能满足条件
	//var max int
	//for i := 0; i < len(arr); i++ {
	//	if max < arr[i] {
	//		max = arr[i]
	//	}
	//}
	//// 检查这个最大值是否满足条件，与长度有关
	//// arr[len(arr)-1] <= len(arr)
	//// 还需要保证中间的值够做好桥接!!
	//if len(arr) > max {
	//	return max
	//}
	//return len(arr)
	// TODO 最后的优化应该在这里，能否借助排序去提高速度
	//sort.Ints(arr)
	//// 直接做一次排列
	//// 从前往后开始保证条件满足
	//// 首先如果arr[0] 不是1则置位1
	//arr[0] = 1
	//for i := 1; i < len(arr); i++ {
	//	if arr[i] > arr[i-1] {
	//		arr[i] = arr[i-1] + 1
	//	}
	//}
	//return arr[len(arr)-1]

	// 计数排序+贪心
	// 出于最后的结果不可能比len(arr)还要大，如果比其还大直接放在最后
	cnt := make([]int, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		if arr[i] >= len(arr) {
			cnt[len(arr)]++
		} else {
			cnt[arr[i]]++
		}
	}
	// 如果有缺失的数字，记录下来，后续遇到不缺失的数字则其减少到已有的数字那里
	miss := 0
	for i := 1; i < len(cnt); i++ {
		if cnt[i] == 0 {
			miss++
		} else {
			if cnt[i]-1 <= miss {
				miss = miss - cnt[i] + 1 // 留一个站住本身的位置
			} else { // 还有更多的，但是只能保持这样，不能增大只能减小的规则
				miss = 0
			}
		}
	}
	return len(arr) - miss
}
