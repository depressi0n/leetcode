package question

import "reflect"

// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

func permuteUnique(nums []int) [][]int {
	return permuteUniqueCore(nums)
}

func permuteUniqueCore(nums []int) [][]int {
	nextPermutation := func(nums []int) []int {
		res := make([]int, len(nums))
		copy(res, nums)
		// 从右往左找到递减的第一个
		i := len(res) - 2
		for ; i >= 0 && res[i] >= res[i+1]; i-- {
			// pass
		}
		if i >= 0 {
			// 从右往左找到比i所在值大的第一个
			j := len(nums) - 1
			for ; j > i && res[j] <= res[i]; j-- {
				// pass
			}
			res[i], res[j] = res[j], res[i]
		}
		for j := i + 1; j < (len(res)+i+1)/2; j++ {
			res[j], res[len(res)-j+i] = res[len(res)-j+i], res[j]
		}
		return res
	}
	res := make([][]int, 0, len(nums))
	flag := false
	res=append(res,nums)
	for !reflect.DeepEqual(res[len(res)-1], nums) || !flag {
		flag=true
		res = append(res, nextPermutation(res[len(res)-1]))
	}
	return res[:len(res)-1]
}
