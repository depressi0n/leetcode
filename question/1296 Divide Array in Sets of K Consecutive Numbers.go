package question

import "math"

func IsPossibleDivide(nums []int, k int) bool {
	// TODO：避免性能问题，可以直接排序处理...
	// 将数组进行划分，划分成连续数字的集合
	// 所以第一步判断数组长度是否是k的整数倍，否则直接返回false
	if len(nums)%k != 0 || len(nums) < k {
		return false
	}
	// 然后对数组进行处理，处理方式如下
	// 记录数组中最大值和最小值，并保存在map中
	// 遍历最大值和最小值，依次取出一个集合，直到出现结束或者出现无法完成的问题
	m := make(map[int]int)
	minValue := math.MaxInt32
	maxValue := -1
	for i := 0; i < len(nums); i++ {
		minValue = min(minValue, nums[i])
		maxValue = max(maxValue, nums[i])
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = 1
		} else {
			m[nums[i]] = m[nums[i]] + 1
		}
	}
	for i := minValue; i <= maxValue; i++ {
		if _, ok := m[i]; ok {
			if m[i] > 0 {
				m[i] = m[i] - 1
				// 往后每个数都要减小一个
				for j := 1; j < k; j++ {
					if cnt, ok1 := m[i+j]; ok1 && cnt > 0 {
						m[i+j] = cnt - 1
					} else {
						return false
					}
				}
				i--
			}
		}
	}
	return true
}
