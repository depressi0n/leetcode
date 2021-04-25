package question

import "fmt"

func mergeIn88(nums1 []int, m int, nums2 []int, n int) {
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
	for q >= 0 {
		nums1[cur] = nums2[q]
		cur--
		q--
	}
	fmt.Printf("%v", nums1)
	return
}
