package question

// 题目描述：给你一个数组 nums和一个值 val，你需要 原地 移除所有数值等于val的元素，并返回移除后数组的新长度。
//不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

func removeElement(nums []int, val int) int {
	return removeElementCore(nums,val)
}

// 思路同上一题
func removeElementCore(nums []int, val int) int {
	k:=0
	for i:=0;i<len(nums);i++{
		if nums[i]==val{
			continue
		}
		nums[k]=nums[i]
		k++
	}
	return k

}