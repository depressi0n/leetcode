package question

// 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
func singleNumber0137(nums []int) int {
	return singleNumber0137Core4(nums)
}

// 朴素思想:使用一个map记录下来
func singleNumber0137Core1(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for i, cnt := range m {
		if cnt == 1 {
			return i
		}
	}
	return -1
}

// 优化思想：每一位上出现0，1的次数与3的余数正是出现一次的元素的相应位上的数
// 那么问题转变成：怎么计算每一个位上的数的和？ => (a  >> i )& 1 得到第i位上的数
func singleNumber0137Core2(nums []int) int {
	res := 0
	for i := 0; i < 32; i++ {
		totoal := 0
		for j := 0; j < len(nums); j++ {
			totoal += (nums[j] >> i) & 1
		}
		if totoal%3 == 1 {
			res |= 1 << i
		}
	}
	return res
}

// 优化思想：转变为电路设计
// 一共四种门电路：与门、或门、非门、异或门
// 借助门电路，【黑盒】的第i位为「0，1，2」其中一个，表示当前遍历过的所有整数的第i位之和除以3的余数
// 需要借助两个整数a，b来存储上面的数，00表示0，01 表示1，10表示2，当遍历到一个新的数时，若其第i位上的数为0，则a，b不变，若其第i位上位，则a、b的第i位按照00 -> 01 -> 10 -> 00循环变化
// 得到真值表
// ai,bi 	xi 	ai,bi  	ai		bi
// 00 		0 -> 00 	0		0
// 00 		1 -> 01 	0		1
// 01 		0 -> 01		0		1
// 01 		1 -> 10		1		0
// 10 		0 -> 10		1		0
// 10 		1 -> 00		0		0
// 转成逻辑表达式为
// ai = (-ai & bi & xi) | (ai & -bi & -xi)
// bi = -ai & (bi ^ xi)
// 结果返回时，aibi要么是00表示0，要么是01表示1，取bi作为最后结果即可
func singleNumber0137Core3(nums []int) int {
	a, b := 0, 0
	for _, num := range nums {
		a, b = b&^a&num|a&^b&^num, (b^num)&^a
	}
	return b
}

// 优化：简化计算a的规则，利用先计算好的bi
// 使用新的bi代替旧的bi
// 此时真值表为：
// ai,bi 	xi 	ai,bi  	ai
// 00 		0 -> 00 	0
// 01 		1 -> 01 	0
// 01 		0 -> 01		0
// 00 		1 -> 10		1
// 10 		0 -> 10		1
// 10 		1 -> 00		0
// 此时有逻辑表达式
// bi = -ai & (bi ^ xi)
// ai = -bi & (ai ^ xi)
func singleNumber0137Core4(nums []int) int {
	a, b := 0, 0
	for _, num := range nums {
		b = (b ^ num) &^ a
		a = (a ^ num) &^ b
	}
	return b
}
