package question

func majorityElement(nums []int) int {
	may := 0
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[may] {
			cnt++
		} else {
			cnt--
			if cnt == 0 {
				may = i + 1
			}
		}
	}
	if may >= len(nums) {
		return -1
	}
	cnt = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[may] {
			cnt++
		}
	}
	if cnt > len(nums)/2 {
		return nums[may]
	} else {
		return -1
	}
}
