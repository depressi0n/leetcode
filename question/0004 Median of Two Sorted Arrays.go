package question

import "math"

// 题目描述：给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数 。
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArraysCore1(nums1, nums2)
}

// 通过合并的方式找到中位数，合并到一半就可以找到了
// TODO：合并的时候可以不合并，可以通过移动指针的方式定位的中位数，需要记录一个pre，这是为了处理偶数的情况
func findMedianSortedArraysCore1(nums1 []int, nums2 []int) float64 {
	merged := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
		if len(merged) > (len(nums1)+len(nums2))/2 {
			break
		}
	}
	if len(merged) > (len(nums1)+len(nums2))/2 {
		// 长度为偶数的情况
		if (len(nums1)+len(nums2))&1 == 0 {
			return float64(merged[len(merged)-1]+merged[len(merged)-2]) / 2
		} else {
			return float64(merged[len(merged)-1])
		}
	}
	for i < len(nums1) {
		merged = append(merged, nums1[i])
		if len(merged) > (len(nums1)+len(nums2))/2 {
			break
		}
		i++
	}
	for j < len(nums2) {
		merged = append(merged, nums2[j])
		if len(merged) > (len(nums1)+len(nums2))/2 {
			break
		}
		j++
	}
	// 长度为偶数的情况
	if (len(nums1)+len(nums2))&1 == 0 {
		return float64(merged[len(merged)-1]+merged[len(merged)-2]) / 2
	} else {
		return float64(merged[len(merged)-1])
	}
}

// 你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？
// 看到对数级别，想一想怎么去掉一半 => 回归到求第k小数
// 这里的思路是取两个数组的中间值，比较，去掉其中一个的左半部分
// 甚至于可以是递归的思路
func findMedianSortedArraysCore2(nums1 []int, nums2 []int) float64 {
	left, right := (len(nums1)+len(nums2)+1)/2, (len(nums1)+len(nums2)+2)/2
	// 合并奇偶，如果是奇数，会求两次相同的k
	return float64(findKth0004(nums1, nums2, left)+findKth0004(nums1, nums2, right)) * 0.5
}
func findKth0004(nums1 []int, nums2 []int, k int) int {
	min := func(x int, y int) int {
		if x < y {
			return x
		}
		return y
	}
	// 因为每次都是去掉左边部分，所有没有end，只有start
	index1, index2 := 0, 0
	for {
		// 如果有一个数组变空了，则从另一个数组返回相应的位置上的数即可
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		// 如果已经是最小的数了，则从两个数组开头找到第一个数即可
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		// 比较当前一半的位置，取min(index1+k/2,len(nums1))是为了避免数组越界
		newIndex1 := min(index1+k/2, len(nums1)) - 1
		newIndex2 := min(index2+k/2, len(nums2)) - 1
		if nums1[newIndex1] < nums2[newIndex2] {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}

// 中位数的本质：将一个集合划分为两个长度相等的子集，其中一个子集中的元素总是大于另一个子集中的元素
// 划分数组的方法：一个数组中有m个元素，则对应着m+1种划分方法
// 当总长度是偶数时，划分后的两个部分容量相同，并且左边的最大值小于右边的最小值，则取二者的平均值
// 当总长度是奇数时，划分后的两个部分容量--左边部分比右边部分多1个，并且左边的最大值小于右边的最小值，则取左边的最大值
// i+j=m-i+n-j （偶数） i+j=m-i+n-j+1（奇数） => i+j=（m+n+1）/2 => 规定第一个数组的长度比第二个数组的小，那么对任意的i，都有j=（m+n+1）/2-i
// 此时枚举i并得到j
// TODO：思想理解了，但代码需要进一步理解
func findMedianSortedArraysCore3(nums1 []int, nums2 []int) float64 {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)
	left, right := 0, m
	median1, median2 := 0, 0
	// 借助二分查找找到划分方式
	for left <= right {
		i := (left + right) / 2
		j := (m+n+1)/2 - i

		// 左边部分的最大值
		nums_im1 := math.MinInt32
		if i != 0 {
			nums_im1 = nums1[i-1]
		}
		nums_i := math.MaxInt32
		if i != m {
			nums_i = nums1[i]
		}
		// 右边部分的最小值
		nums_jm1 := math.MinInt32
		if j != 0 {
			nums_jm1 = nums2[j-1]
		}
		nums_j := math.MaxInt32
		if j != n {
			nums_j = nums2[j]
		}
		if nums_im1 <= nums_j {
			median1 = max(nums_im1, nums_jm1)
			median2 = min(nums_i, nums_j)
			left = i + 1
		} else {
			right = i - 1
		}
	}
	if (m+n)%2 == 0 {
		return float64(median1+median2) / 2.0
	}
	return float64(median1)
}
