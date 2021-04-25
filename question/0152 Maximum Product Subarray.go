package question

// 求最大值和最小值
// 这里是通过滚动数组的思想做了优化
// 以当前值结尾的最大积需要判断当前值的符号和之前的符号
// 如果和之前同号，那么最大值和最小值都会保持
// 而如果出现异号，则最大值和最小值会反转
// 还有一种就是出现0的时候需要重置最大值和最小值
func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}

//TODO：另一条有意思的思路是
// 对出现的0进行分段处理
// 每段如果出现了偶数个负数，那么肯定不会有影响，全部乘
// 每段如果出现了奇数个负数，那么最大值肯定出现在偶数个负数之间的结果

// TODO:窗口扩展的思路应该是成立的,但是应该怎么写？下面的写法是错误的
func maxProduct1(nums []int) int {
	flag := false
	res := 0
	// 跳过前置0
	start := 0
	for start < len(nums) {
		if nums[start] == 0 { //跳过0
			flag = true
			start++
			continue
		}
		// 保证当前的区间积是正的
		end := start + 1
		if end == len(nums) { //说明只有一个值
			if nums[start] < 0 {
				if res == 0 && !flag {
					res = nums[start]
				}
			} else {
				if res < nums[start] {
					res = nums[start]
				}
			}
			break
		}
		for end < len(nums) && nums[end]*nums[start] < 0 {
			end++
		}
		if end == len(nums) { //nums[start:]就nums[start]一个负数
			tmp := 1
			for i := start; i < end; i++ {
				tmp *= nums[i]
			}
			if tmp > res {
				res = tmp
			}
			return res
		}
		// 扩展
		// 右侧扩展
		for end+1 < len(nums) && nums[end+1] > 0 {
			end++
		}
		// 右侧不能扩展的时候，尝试往两侧扩展
		if start >= 1 && end+1 < len(nums) && nums[start-1]*nums[end+1] > 0 {
			start--
			end++
			for start >= 1 && nums[start-1] > 0 {
				start--
			}
			continue //继续往右边扩展
		}
		// 说明无法扩展
		tmp := 1
		for i := start; i <= end; i++ {
			tmp *= nums[i]
		}
		if tmp > res {
			res = tmp
		}
		start = end + 1
	}
	if flag && res < 0 {
		return 0
	}
	return res
}
