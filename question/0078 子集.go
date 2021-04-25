package question

// 二进制表示所有子集，选中为1不选中为0
func subsets(nums []int) [][]int {
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
