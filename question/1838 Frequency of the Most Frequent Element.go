package question

import "sort"

func Answer1838(nums []int, k int) int {
	return maxFrequency(nums, k)
}
func maxFrequency(nums []int, k int) int {
	// 最多k次操作，每次操作可以增加某个位置上的数+1
	// 目标：数组中元素可能的最大频度

	// 思想：首先应该排序，然后从后往前，尝试增加，求出此时最大频度
	//sort.Ints(nums)
	//res:=1
	//for i := len(nums)-1; i >=0; i-- {
	//	if i<res{
	//		break
	//	}
	//	t:=k
	//	j:=i-1
	//	for j>=0 && t>=nums[i]-nums[j] {
	//		t-= nums[i]-nums[j]
	//		j--
	//	}
	//	if i-j>res{
	//		res=i-j
	//	}
	//}
	//return res

	// 优化：
	//（1）如果可以将一组数字全部变成y，假设其中最大值为x，则一定可以全部变成x，且频数不变，操作次数更少
	//（2）变成较小者更优一点
	//（3）遍历每个元素能否优化？ => 滑动窗口
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	sort.Ints(nums)
	ans := 1
	// 探测左右边界
	for l, r, total := 0, 1, 0; r < len(nums); r++ {
		total += (nums[r] - nums[r-1]) * (r - l)
		for total > k { // 不够
			total -= nums[r] - nums[l]
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans

}
