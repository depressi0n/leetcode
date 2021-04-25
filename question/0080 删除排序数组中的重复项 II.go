package question

func removeDuplicatesII(nums []int) int {
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

// 另一条思路是，维护两个指针，间隔为2，left=0,right=2，同时记录整体的长度
//（1）如果左边等于右边，则删除右边指向元素，
//（2）如果左边不等于右边，则两个一起往后移动一个
