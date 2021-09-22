package question

import "fmt"

// 给定一个包含红色、白色和蓝色，一共n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
//此题中，我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。
// 三色国旗问题
func sortColors(nums []int) {
	fmt.Println(nums)
	sortColorsCore(nums)
	fmt.Println(nums)
}
func sortColorsCore(nums []int) {
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
	//fmt.Printf("%v\n", nums)
}

// TODO:理解思想，本质上无论如何都是三个指针，只是使用的思想有点点区别
func sortColorsCore2(nums []int) {
	p0, p1 := 0, 0
	for i, c := range nums {
		if c == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if c == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}
func sortColorsCore3(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for ; i <= p2 && nums[i] == 2; p2-- {
			nums[i], nums[p2] = nums[p2], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}
