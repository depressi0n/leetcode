package question

import "fmt"

// 给你两个按 非递减顺序 排列的整数数组nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
//请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
//注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
func mergeIn0088(nums1 []int, m int, nums2 []int, n int) {
	mergeIn0088Core1(nums1, m, nums2, n)
}

// 主要思想：保证nums1有足够的空间，从大到小，从后往前排序，这样不会覆盖掉nums1之前的数
func mergeIn0088Core1(nums1 []int, m int, nums2 []int, n int) {
	p, q := m-1, n-1
	cur := m + n - 1
	for p >= 0 && q >= 0 {
		if nums1[p] >= nums2[q] {
			nums1[cur] = nums1[p]
			p--
		} else {
			nums1[cur] = nums2[q]
			q--
		}
		cur--
	}
	// 只需考虑nums2是否排序完成即可
	for q >= 0 {
		nums1[cur] = nums2[q]
		cur--
		q--
	}
	fmt.Printf("%v", nums1)
	return
}
