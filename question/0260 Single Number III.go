package question

// 给定一个整数数组nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。
//进阶：你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？
func singleNumber0260(nums []int) []int {
	return singleNumber0260Core3(nums)
}

// 主要思想：用一个map进行记录，最后遍历一次map，得到结果
func singleNumber0260Core1(nums []int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	res := make([]int, 0, len(nums))
	for i, cnt := range m {
		if cnt == 1 {
			res = append(res, i)
		}
	}
	return res
}

// 优化方向：如果利用异或运算进行处理，那么异或的最后结果肯定是a || b
// 但是如何从其中提取出a和b？
// 下面的思路很有意思：
// 利用 res & -res 取出最低位，则a或b中在这个位上一个是0一个是1
// 根据 这个情况，对nums中每个数按照相应位是否为1进行分类，一共两类
// 此时，每个数都会分入其中一组，并且a和b也分进了不同的组
// 那么，对这两组进行异或，最后就可以得到结果
func singleNumber0260Core2(nums []int) []int {
	xorsum := 0
	for i := 0; i < len(nums); i++ {
		xorsum ^= nums[i]
	}
	lsb := xorsum & -xorsum
	type1,type2:=0,0
	for i := 0; i < len(nums); i++ {
		if nums[i]&lsb>0{
			type1^=nums[i]
		}else{
			type2^=nums[i]
		}
	}
	return []int{type1,type2}
}

// 可以省一个变量
func singleNumber0260Core3(nums []int) []int {
	xorsum := 0
	for i := 0; i < len(nums); i++ {
		xorsum ^= nums[i]
	}
	lsb := xorsum & -xorsum
	type1:=0
	for i := 0; i < len(nums); i++ {
		if nums[i]&lsb>0{
			type1^=nums[i]
		}
	}
	return []int{type1,xorsum^type1}
}
