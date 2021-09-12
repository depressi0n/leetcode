package question

// 题目描述：缺失的第一个正数
// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//
//请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
func firstMissingPositive(nums []int) int {
	return firstMissingPositiveCore3(nums)
}

// 缺失的第一个正数不可能超过数组的长度，利用这个信息，建立一个映射
// 如果值比长度大，则映射到n上即可
func firstMissingPositiveCore1(nums []int) int {
	cnt := make([]int, len(nums)+2)
	for i := 0; i < len(nums); i++ {
		if nums[i] > len(nums) {
			cnt[len(nums)+1]++
		} else if nums[i] <= 0 {
			cnt[0]++
		} else {
			cnt[nums[i]]++
		}
	}
	for i := 1; i < len(cnt); i++ {
		if cnt[i] == 0 {
			return i
		}
	}
	return len(nums) + 1
}

// 将map的信息集成到原始数组中：
// 第一次遍历所有的数，如果数在【1，N】中，变成其负数[可以和第二次一起，此时要考虑周全一点]，如果是负数，修改为比N大即可（取N+1）
// 第二次遍历，对负数的绝对值减去位置上的数变成相反数（回到正数）
// 第三次遍历，如果这个数不是正数，说明没有出现过，即为答案
func firstMissingPositiveCore2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			nums[i] = len(nums) + 1
		}
	}
	for i := 0; i < len(nums); i++ {
		num := abs(nums[i])
		if num <= len(nums) {
			nums[num-1] = -abs(nums[num-1])
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}

// 直接置换，如果当前位置上的数不在应该的位置上，多次置换
func firstMissingPositiveCore3(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for 0 < nums[i] && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
