package question

import "fmt"

// 三色国旗问题
func sortColors(nums []int) {
	i, k, j := 0, 0, len(nums)-1
	for k <= j {
		switch {
		case nums[j] == 2:
			j--
		case nums[j] == 1:
			nums[j], nums[k] = nums[k], nums[j]
			k++
		case nums[j] == 0:
			nums[i], nums[j] = nums[j], nums[i]
			if i != k {
				nums[j], nums[k] = nums[k], nums[j]
			}
			i++
			k++
		}
	}
	fmt.Printf("%v\n", nums)
}
