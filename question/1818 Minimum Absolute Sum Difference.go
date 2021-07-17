package question

import "sort"

func Answer1818(nums1 []int, nums2 []int) int {
	return minAbsoluteSumDiff(nums1, nums2)
}

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	// 最多使用一次：用nums1中的元素替换其中一个，使得sum of |nums1[i]-nums2[i]|最小
	// 首先是最暴力的方法，每一个元素都替换一次，记录最小值（每种可能都算一次）
	// 会超时
	//originalSum := 0
	//abs := func(a int) int {
	//	if a < 0 {
	//		return -a
	//	}
	//	return a
	//}
	//for i := 0; i < len(nums1); i++ {
	//	t:=abs(nums2[i]-nums1[i])
	//	originalSum = (originalSum + t) % 1000_000_007
	//}
	//res := originalSum
	//for i := 0; i < len(nums1); i++ {
	//	// nums1[i]替换nums1[j]
	//	for j := 0; j < len(nums1); j++ {
	//		tmp := (originalSum - abs(nums2[j]-nums1[j]) + abs(nums2[j]-nums1[i])) % 1000_000_007
	//		if tmp < res {
	//			res = tmp
	//		}
	//	}
	//}
	//return res

	// 优化策略
	// 用来替换的值越接近nums2[k]越好
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	const mod = 1000_000_007
	help := sort.IntSlice{}
	for i := 0; i < len(nums1); i++ {
		help = append(help, nums1[i])
	}
	// 那个值最接近？排序+二分
	help.Sort()
	originalSum := 0
	diff := 0
	for i := 0; i < len(nums1); i++ {
		t := abs(nums2[i] - nums1[i])
		originalSum = (originalSum + t) % mod
		// 找到的值是第一个大于nums2[i]的值，可能的替换是左右邻居
		j := help.Search(nums2[i])
		if j < len(nums1) {
			diff = max(diff, t-(help[j]-nums2[i]))
		}
		if j > 0 {
			diff = max(diff, t-(nums2[i]-help[j-1]))
		}
	}
	return (originalSum - diff + mod) % mod

}
