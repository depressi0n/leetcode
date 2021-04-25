package question

func maxSubArray(nums []int) int {
	var res = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i+1] > 0 {
			nums[i] += nums[i+1]
		}
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}
