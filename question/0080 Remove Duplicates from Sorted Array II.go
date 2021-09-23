package question

// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 最多出现两次 ，返回删除后数组的新长度。
//
//不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

func removeDuplicatesII(nums []int) int {
	return removeDuplicatesIICore2(nums)
}

// 使用一个k记录当前结尾的位置，因为k<=i，所以不会有丢失的情况发生
// 而最多出现几次这个过程可以通过一个往后查找添加的方式完成，如果需要的话可以函数化，直接返回下一个位置，在函数内完成最多k个重复的过程
func removeDuplicatesIICore1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prev := nums[0] - 1
	cnt := 0
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != prev { //如果和前面的值不相同，则插入到给定位置上，并将cnt置为1
			prev = nums[i]
			nums[k] = nums[i]
			cnt = 1
			k++
			continue
		}
		if nums[i] == prev { // 如果和前面的值相同
			if cnt < 2 { // 如果此时还没满，则添加
				nums[k] = nums[i]
				k++
				cnt++
			} else { // 满了，则跳过后面相同的，然后肯定会跳到上面的循环
				j := i + 1
				for j < len(nums) && nums[j] == nums[i] {
					j++
				}
				i = j - 1
			}
		}
	}
	return k
}
// TODO：思考一下这个思想的实现
// 慢指针是记录当前存下来的值的末尾，快指针是工作指针，这个实现很简洁！思想是一样的，但是实现确实比上面简洁很多
// 另一条思路是，维护两个指针，间隔为2，left=0,right=2，同时记录整体的长度
//（1）如果左边等于右边，则删除右边指向元素，
//（2）如果左边不等于右边，则两个一起往后移动一个
func removeDuplicatesIICore2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	left, right := 2, 2
	for right < len(nums) {
		if nums[left-2] != nums[right] {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}
