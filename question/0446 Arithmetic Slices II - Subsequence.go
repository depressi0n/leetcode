package question

func Answer0446(nums []int) int {
	return numberOfArithmeticSlices0446(nums)
}
func numberOfArithmeticSlices0446Error(nums []int) int {
	// 子序列是arithmetic subsequences：定义为中间的差值相等，长度不小于3
	// 思路：依旧是左右指针，每次都是去找最长的，这样复杂度会上去
	// 找到最长的之后就可以计算一个总值，然后左指针尽可能往右移动，移动到与左指针连续的结束位置，继续找
	if len(nums) < 3 {
		return 0
	}
	res := 0
	tmp := []int{0, 1}
	diff := nums[tmp[1]] - nums[tmp[0]]
	for tmp[len(tmp)-1] < len(nums) {
		for cur := tmp[len(tmp)-1] + 1; cur < len(nums); cur++ {
			if  nums[cur] == nums[tmp[0]]+len(tmp)*diff{
				tmp = append(tmp, cur)
			}
		}
		if len(tmp)<3{
			tmp=[]int{tmp[1],tmp[1]+1}
			diff=nums[tmp[1]] - nums[tmp[0]]
			continue
		}
		// 中间的累加是有问题的
		length := len(tmp)
		for ; length >= 3; length = (length + 1) / 2 {
			res += (1 + length - 2) * (length - 2) / 2
		}

		i := 1
		for ; i < len(tmp) && tmp[i]-tmp[i-1] == 1; i++ {
			// pass
		}
		if tmp[i-1]+1>=len(nums){
			break
		}
		tmp = []int{tmp[i-1], tmp[i-1] + 1}
		diff=nums[tmp[1]] - nums[tmp[0]]

	}
	return res
}

func numberOfArithmeticSlices0446(nums []int)int  {
	ans:=0
	//定义状态f[i][d]表示尾项为nums[i],公差为d的弱等差子序列的个数。
	f := make([]map[int]int, len(nums))
	for i, x := range nums {
		f[i] = map[int]int{}
		for j, y := range nums[:i] {
			d := x - y
			cnt := f[j][d]
			ans += cnt
			f[i][d] += cnt + 1
		}
	}
	return ans
}