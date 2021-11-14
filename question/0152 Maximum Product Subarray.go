package question

// 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
func maxProduct(nums []int) int {
	return maxProductCore1(nums)
}

// 主要思想：动态规划的思想
// maxF[i]表示以nums[i]结尾的最大乘积，minF[i]表示以nums[i]结尾的最小乘积
// maxF[i]=max(maxF[i-1]*nums[i],minF[i-1]*nums[i],nums[i])
// minF[i]=max(maxF[i-1]*nums[i],minF[i-1]*nums[i],nums[i])
// 最后的结果就是max(maxF[0],...,maxF[len(nums)-1])
// 根据转移方程，考虑空间优化：通过滚动数组的优化思想，可以使用两个值来代替两个数组
// 还有一种就是出现0的时候需要重置最大值和最小值
func maxProductCore1(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}

// TODO：另一条有意思的思路是
// 对出现的0进行分段处理
// 每段如果出现了偶数个负数，那么肯定不会有影响，全部乘
// 每段如果出现了奇数个负数，那么最大值肯定出现在偶数个负数之间的结果
// 操作步骤
// 首先找到所有的0，完成分段
// 对于每一段，查找负数有多少个，如果负数为偶数个，整段的乘积就是最大的，如果为奇数个，则要去掉第一个奇数以及之前的数或去掉最后一个奇数及之后的数
func maxProductCore2(nums []int) int {
	// 假定左右两端均有0的哨兵
	zeros := make([]int, 0, len(nums))
	zeros=append(zeros,-1)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeros = append(zeros, i)
		}
	}
	zeros=append(zeros,len(nums))
	res := 0
	for i := 0; i < len(zeros)-1; i++ {
		left, right := zeros[i], zeros[i+1]
		// 统计奇数的个数，并记录第一个奇数和最后一个奇数
		cnt := 0
		first, last := -1, -1
		for j := left + 1; j < right; j++ {
			if nums[j] < 0 {
				if first == -1 {
					first = j
				}
				if last == -1 || last < j {
					last = j
				}
				cnt++
			}
		}
		if cnt&1 == 0 {
			t := 1
			for j := left + 1; j < right; j++ {
				t *= nums[j]
			}
			res = max(res, t)
		} else {
			if left+1<last{
				t := 1
				for j := left + 1; j < last; j++ {
					t *= nums[j]
				}
				res = max(res, t)
			}
			if first+1<right{
				t := 1
				for j := first + 1; j < right; j++ {
					t *= nums[j]
				}
				res = max(res, t)
			}
		}
	}
	return res
}
// TODO:滑动窗口的思想是否能用到这里？
