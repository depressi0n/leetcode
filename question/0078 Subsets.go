package question

// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
//解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

// 二进制表示所有子集，选中为1不选中为0
func subsets(nums []int) [][]int {
	return subsetsCore(nums)
}
func subsetsCore(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return [][]int{{}}
	}
	if len(nums) == 1 {
		return [][]int{{}, nums}
	}
	res := [][]int{{}} // length=0
	var dfsIn func(int, int, []int)
	dfsIn = func(length, pos int, cur []int) {
		more := length - len(cur)
		if more == 0 {
			curCopy := make([]int, length)
			copy(curCopy, cur)
			res = append(res, curCopy)
			return
		}
		if more > 0 && pos < len(nums) {
			//考虑当前这个
			cur = append(cur, nums[pos])
			dfsIn(length, pos+1, cur)
			cur = cur[:len(cur)-1]
			//不考虑当前这个
			if pos+1 < len(nums) {
				dfsIn(length, pos+1, cur)
			}
		}
	}
	for i := 1; i <= len(nums); i++ { // length
		temp := []int{}
		dfsIn(i, 0, temp)
	}
	return res
}
